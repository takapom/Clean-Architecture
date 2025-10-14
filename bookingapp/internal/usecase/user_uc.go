package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"
)

var (
	ErrUserInvalidInput       = errors.New("invalid user input")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)

type RegisterUserInput struct {
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	DateOfBirth string
}

// ユースケース層からrepository層のinterfaceを使えるようにする
type UserUsecase struct {
	Users repository.UserRepository
	Now   func() time.Time
}

func (u *UserUsecase) Register(in RegisterUserInput) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	name := strings.TrimSpace(in.Name)
	email := strings.TrimSpace(in.Email)
	if name == "" || email == "" {
		return nil, ErrUserInvalidInput
	}

	existing, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUserEmailAlreadyExists
	}

	var dob time.Time
	if v := strings.TrimSpace(in.DateOfBirth); v != "" {
		dob, err = time.Parse("2006-01-02", v)
		if err != nil {
			return nil, ErrUserInvalidInput
		}
	}

	now := time.Now
	if u.Now != nil {
		now = u.Now
	}

	user := &entity.User{
		Name:         name,
		Email:        email,
		PhoneNumber:  strings.TrimSpace(in.PhoneNumber),
		Address:      strings.TrimSpace(in.Address),
		DateOfBirth:  dob,
		RegisteredAt: now(),
		Status:       "active",
	}

	return u.Users.Create(user)
}

func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, ErrUserInvalidInput
	}

	return u.Users.Get(trimmed)
}
