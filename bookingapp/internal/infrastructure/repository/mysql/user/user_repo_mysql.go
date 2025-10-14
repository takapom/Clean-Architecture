package user

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// データベース操作を行うための実装
type UserRepo struct {
	db *gorm.DB
}

// 依存注入のためのコンストラクタ
func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	var dob *time.Time
	if !user.DateOfBirth.IsZero() {
		d := user.DateOfBirth
		dob = &d
	}

	model := usermodel.UserModel{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  dob,
		RegisteredAt: user.RegisteredAt,
		Status:       user.Status,
	}

	if err := r.db.WithContext(context.Background()).Create(&model).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, nil
	}

	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("email = ?", strings.TrimSpace(email)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}

var _ repository.UserRepository = (*UserRepo)(nil)

func modelToEntity(model *usermodel.UserModel) *entity.User {
	if model == nil {
		return nil
	}

	var dob time.Time
	if model.DateOfBirth != nil {
		dob = *model.DateOfBirth
	}

	return &entity.User{
		ID:           model.ID,
		Name:         model.Name,
		Email:        model.Email,
		PhoneNumber:  model.PhoneNumber,
		Address:      model.Address,
		DateOfBirth:  dob,
		RegisteredAt: model.RegisteredAt,
		Status:       model.Status,
	}
}
