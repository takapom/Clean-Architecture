package db

import (
	"bookingapp/internal/infrastructure/db/models"
	"bookingapp/internal/infrastructure/db/models/user" // UserModel をインポート
)

func Migrate(db any) error {
	gdb := db.(interface{ AutoMigrate(...any) error })
	return gdb.AutoMigrate(
		&models.PlanModel{},
		&models.ReservationModel{},
		&user.UserModel{}, // UserModel を追加
	)
}
