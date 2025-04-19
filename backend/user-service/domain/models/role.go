package models

import "time"

type Role struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Code      string `gorm:"varcchar(15);not null;"`
	Name      string `gorm:"varchar(30);not null;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
