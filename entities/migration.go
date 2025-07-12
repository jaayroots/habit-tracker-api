package entities

import (
	"time"

	"github.com/google/uuid"
)

type Migration struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}
