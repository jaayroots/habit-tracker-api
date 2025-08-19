package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/habit-tracker-api/config"
	"github.com/habit-tracker-api/database"
	"github.com/habit-tracker-api/entities"

	"github.com/habit-tracker-api/command/migration/packImport"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {

	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	migrateInteractive(tx)
}

func migrateInteractive(tx *gorm.DB) {

	if !tx.Migrator().HasTable("migrations") {
		tx.Migrator().CreateTable(&entities.Migration{})
		tx.Commit()
		if tx.Error != nil {
			tx.Rollback()
			panic(tx.Error)
		}
		fmt.Println("✅ Created migrations table successfully.")
		fmt.Println("Please restart the application to continue with migrations.")
		return
	}

	// create or update new file fr migrationInit
	migrationPendingList := getMigrationList(tx, false)
	createMigrationInitFile(migrationPendingList)

	for {
		fmt.Println("\n[ 🤖 MIGRATION MENU ]")
		fmt.Println("1) 📜  Migration List")
		fmt.Println("2) ✅  Create Migration")
		fmt.Println("3) 🚀  Run Migration")
		fmt.Println("4) 🔙  Rollback Migration")
		fmt.Println("5) ❌  Exit")
		fmt.Println("\n")
		fmt.Print("➡️  Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)
		fmt.Println("\n")
		switch choice {
		case 1:
			fmt.Println("📜 Migration List:")
			getMigrationList(tx, true)
			return
		case 2:
			var name string
			fmt.Print("📝 Enter migration name (e.g. create_users_table): ")
			fmt.Scanln(&name)

			if strings.TrimSpace(name) == "" {
				fmt.Println("❌ Migration name cannot be empty.")
				continue
			}
			createMigrationFile(tx, name)
			return
		case 3:
			migrationPendingList := getMigrationList(tx, false)
			runMigration(tx, migrationPendingList)
			return
		case 4:
			migrationPendingList := getMigrationList(tx, false)
			rollbackMigration(tx, migrationPendingList)
			return
		case 5:
			return
		default:
			fmt.Println("❌ Invalid choice, please try again.")
			return
		}
	}
}

type MigrationPendingFile struct {
	Status bool
	Name   string
}

const migrationFolder = "./command/migration/list"

func getMigrationList(tx *gorm.DB, isRender bool) []MigrationPendingFile {

	var appliedMigrations []entities.Migration
	if err := tx.Find(&appliedMigrations).Error; err != nil {
		fmt.Println("❌ Error loading migrations from DB:", err)
		return nil
	}

	appliedMap := make(map[string]bool)
	for _, m := range appliedMigrations {
		appliedMap[m.Name] = true
	}

	files := []string{}
	err := filepath.WalkDir(migrationFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), ".sql")) {
			files = append(files, d.Name())
		}
		return nil
	})
	if err != nil {
		fmt.Println("❌ Error reading migration files:", err)
		return nil
	}

	if len(files) == 0 {
		fmt.Println("📭 No migration files found in folder:", migrationFolder)
		return nil
	}

	sort.Slice(files, func(i, j int) bool {
		ti := files[i]
		tj := files[j]

		prefixI := ti
		prefixJ := tj
		if len(ti) > 14 {
			prefixI = ti[:14]
		}
		if len(tj) > 14 {
			prefixJ = tj[:14]
		}

		return prefixI < prefixJ
	})

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Migration File", "Status"})

	pendingMigration := make([]MigrationPendingFile, 0)
	for _, file := range files {
		base := strings.TrimSuffix(file, ".go")
		status := ""
		if appliedMap[base] {
			status = "✅ Applied"
		} else {
			status = "⏳ Pending"
		}

		pendingMigration = append(pendingMigration, MigrationPendingFile{
			Name:   file,
			Status: appliedMap[base],
		})

		t.AppendRow(table.Row{file, status})
	}
	if isRender {
		t.Render()
	}
	return pendingMigration
}

func createMigrationFile(tx *gorm.DB, name string) {

	safeName := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	files, err := os.ReadDir("command/migration/list")
	if err != nil && !os.IsNotExist(err) {
		panic(fmt.Errorf("❌ Cannnot read migration: %w", err))
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()

		namePart := filename
		if len(filename) > 15 && filename[14] == '_' {
			namePart = filename[15:]
		}
		namePart = strings.TrimSuffix(namePart, ".go")
		namePart = strings.TrimSuffix(namePart, ".sql")

		if namePart == safeName {
			panic(fmt.Errorf("❌ migration name '%s' is already %s", safeName, filename))
		}
	}

	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("command/migration/list/%s_%s.go", timestamp, safeName)
	funcName := fmt.Sprintf("%s_%s", timestamp, safeName)

	if err := os.MkdirAll("command/migration/list", os.ModePerm); err != nil {
		panic(fmt.Errorf("❌ Cannot create migration file: %w", err))
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("❌ Cannote create file: %w", err))
	}
	defer file.Close()

	const migrationTemplate = `package migration

import (
	"gorm.io/gorm"
)

// Up_{{ .FuncName }} applies the migration.
func Up_{{ .FuncName }}(tx *gorm.DB) error {
	// tx.Migrator().CreateTable(&entities.User{})
	// if tx.Error != nil {
	// 	tx.Rollback()
	// 	return tx.Error
	// }
	return nil
}

// Down_{{ .FuncName }} rolls back the migration.
func Down_{{ .FuncName }}(tx *gorm.DB) error {
	// tx.Migrator().DropTable(&entities.User{})
	// if tx.Error != nil {
	// 	tx.Rollback()
	// 	return tx.Error
	// }
	return nil
}
`

	tmpl := template.Must(template.New("migration").Parse(migrationTemplate))

	type MigrationFile struct {
		Timestamp string
		Name      string
		FuncName  string
	}

	err = tmpl.Execute(file, MigrationFile{
		Timestamp: timestamp,
		Name:      safeName,
		FuncName:  funcName,
	})
	if err != nil {
		panic(fmt.Errorf("❌ Cannot write template: %w", err))
	}

	fmt.Println("✅ Migration file created:", filename)
}

func runMigration(tx *gorm.DB, pendingFiles []MigrationPendingFile) error {

	allValid := false
	for _, m := range pendingFiles {
		if !m.Status {
			allValid = true
			continue
		}
	}

	if len(pendingFiles) == 0 || !allValid {
		fmt.Println("📭 No pending migrations to run.")
		return nil
	}

	for _, m := range pendingFiles {

		if m.Status {
			continue
		}

		name := strings.TrimSuffix(m.Name, ".go")

		funcName := packImport.MappingUpFuncMigration()
		fn, ok := funcName[name]
		if !ok {
			return fmt.Errorf("migration %s not found", name)
		}

		if err := fn(tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %s fail: %w", name, err)
		}

		if err := createMigrationHistory(name, tx); err != nil {
			tx.Rollback()
			return err
		}

		fmt.Printf("✅ Migration %s applied\n", name)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %w", err)
	}

	return nil

}

func rollbackMigration(tx *gorm.DB, pendingFiles []MigrationPendingFile) error {

	allValid := false
	for _, m := range pendingFiles {
		if m.Status {
			allValid = true
			continue
		}
	}

	if len(pendingFiles) == 0 || !allValid {
		fmt.Println("📭 No pending migrations to rollback.")
		return nil
	}

	for i := len(pendingFiles) - 1; i >= 0; i-- {
		m := pendingFiles[i]
		if !m.Status {
			continue
		}

		name := strings.TrimSuffix(m.Name, ".go")

		funcName := packImport.MappingDownFuncMigration()
		fn, ok := funcName[name]
		if !ok {
			return fmt.Errorf("migration %s not found", name)
		}

		if err := fn(tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %s fail: %w", name, err)
		}

		if err := tx.Table("migrations").Where("name = ?", name).Delete(&entities.Migration{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("delete migration history %s fail: %w", name, err)
		}

		fmt.Printf("🔙 Migration %s rolled back\n", name)
		break
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %w", err)
	}

	return nil
}

func createMigrationHistory(name string, tx *gorm.DB) error {

	safeName := strings.ToLower(strings.ReplaceAll(name, " ", "_"))

	migration := map[string]interface{}{
		"id":        uuid.New(),
		"name":      fmt.Sprintf("%s", safeName),
		"timestamp": time.Now(),
	}

	if err := tx.Table("migrations").Create(migration).Error; err != nil {
		panic(tx.Error)
	}

	return nil
}

func createMigrationInitFile(migrationFiles []MigrationPendingFile) {

	safeName := "packImport"
	filename := fmt.Sprintf("command/migration/packImport/%s.go", safeName)

	if err := os.MkdirAll("command/migration/list", os.ModePerm); err != nil {
		panic(fmt.Errorf("❌ cannot create migration: %w", err))
	}

	var MigrationPathBuilder strings.Builder
	var mappingUpFuncStringBuilder strings.Builder
	var mappingDownFuncStringBuilder strings.Builder
	for _, m := range migrationFiles {

		name := strings.TrimSuffix(m.Name, ".go")
		slug := strings.ReplaceAll(name, " ", "_")
		slug = strings.ToLower(slug)

		line := fmt.Sprintf(`"%s": migration.Up_%s,`, name, slug)
		mappingUpFuncStringBuilder.WriteString("\t" + line + "\n")

		line = fmt.Sprintf(`"%s": migration.Down_%s,`, name, slug)
		mappingDownFuncStringBuilder.WriteString("\t" + line + "\n")
	}

	// Template Go file
	const migrationTemplate = `package packImport

import (
	{{ .MigrationPath }}
	"gorm.io/gorm"
)

var migrationUpFuncs = map[string]func(*gorm.DB) error{
{{ .MappingUpFuncString }}
}

var migrationDownFuncs = map[string]func(*gorm.DB) error{
{{ .MappingDownFuncString }}
}

func MappingUpFuncMigration() map[string]func(*gorm.DB) error {
	return migrationUpFuncs
}

func MappingDownFuncMigration() map[string]func(*gorm.DB) error {
	return migrationDownFuncs
}
`
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("❌ Cannot create file init: %w", err))
	}
	defer file.Close()

	migrationPath := ""
	if len(migrationFiles) != 0 {
		migrationPath = fmt.Sprintf(`"%s"`, "github.com/habit-tracker-api/command/migration/list")
	}
	MigrationPathBuilder.WriteString(migrationPath)

	// Render template
	tmpl := template.Must(template.New("migration").Parse(migrationTemplate))

	type MigrationTemplateData struct {
		MigrationPath         template.HTML
		MappingUpFuncString   template.HTML
		MappingDownFuncString template.HTML
	}

	err = tmpl.Execute(file, MigrationTemplateData{
		MigrationPath:         template.HTML(MigrationPathBuilder.String()),
		MappingUpFuncString:   template.HTML(mappingUpFuncStringBuilder.String()),
		MappingDownFuncString: template.HTML(mappingDownFuncStringBuilder.String()),
	})

	if err != nil {
		panic(fmt.Errorf("❌ Cannot create template: %w", err))
	}

	fmt.Println("✅ Migration init file created:", filename)
}
