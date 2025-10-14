package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
)

type UserUsecase struct {
	Users repository.UserRepository
}

// ユーザー情報取得
func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	return u.Users.Get(id)
}
