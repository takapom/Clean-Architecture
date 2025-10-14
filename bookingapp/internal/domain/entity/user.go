package entity

import "time"

type User struct {
	ID           string    // ユーザーID
	Name         string    // 名前
	Email        string    // メールアドレス
	PhoneNumber  string    // 電話番号
	Address      string    // 住所（オプション）
	DateOfBirth  time.Time // 生年月日
	RegisteredAt time.Time // 登録日
	Status       string    // アカウントステータス（例: "active", "inactive"）
}
