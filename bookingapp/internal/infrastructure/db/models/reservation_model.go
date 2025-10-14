package models

import "time"

type ReservationModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	PlanID    int       `gorm:"not null;index"`
	Number    int       `gorm:"not null"`
	Checkin   time.Time `gorm:"type:date;not null"`
	Checkout  time.Time `gorm:"type:date;not null"`
	Total     int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ReservationModel) TableName() string { return "reservations" }
