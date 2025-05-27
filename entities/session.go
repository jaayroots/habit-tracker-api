package entities

import "time"

type Session struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UserID    uint    `gorm:"not null;"`
	Token     string    `gorm:"type:text;not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	ExpiresAt time.Time `gorm:"not null;"`
	IpAddress string    `gorm:"type:varchar(128);not null;"`
}
