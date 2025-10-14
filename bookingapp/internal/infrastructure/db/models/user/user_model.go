package user

import "time"

type UserModel struct {
	ID           string     `gorm:"primaryKey;type:char(36)"`
	Name         string     `gorm:"size:255;not null"`
	Email        string     `gorm:"size:255;uniqueIndex;not null"`
	PhoneNumber  string     `gorm:"size:50"`
	Address      string     `gorm:"size:255"`
	DateOfBirth  *time.Time `gorm:"type:date"`
	RegisteredAt time.Time  `gorm:"not null"`
	Status       string     `gorm:"size:50;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (UserModel) TableName() string { return "users" }
