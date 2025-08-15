package enums

type UserType string

const (
	User       UserType = "user"
	Admin      UserType = "admin"
	SuperAdmin UserType = "superadmin"
)

// func (s UserType) String() string {
// 	switch s {
// 	case User:
// 		return "user"
// 	case Admin:
// 		return "admin"
// 	case SuperAdmin:
// 		return "superAdmin"
// 	default:
// 		return "unknow"
// 	}
// }

// func UserTypeMap() map[int]string {
// 	return map[int]string{
// 		int(User):       User.String(),
// 		int(Admin):      Admin.String(),
// 		int(SuperAdmin): SuperAdmin.String(),
// 	}
// }

// func IsValidUserType(value int) bool {
// 	switch UserType(value) {
// 	case User, Admin, SuperAdmin:
// 		return true
// 	default:
// 		return false
// 	}
// }
