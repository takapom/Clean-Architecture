package models

import "time"

type PlanModel struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255;not null"`
	Keyword   string `gorm:"size:255;index"`
	Price     int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PlanModel) TableName() string { return "plans" }
