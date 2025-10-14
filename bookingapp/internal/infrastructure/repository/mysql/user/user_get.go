package user

import (
	"bookingapp/internal/domain/entity"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ---- ユーザー情報取得 ----
func (r *UserRepo) Get(id string) (*entity.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("id is empty")
	}
	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("id = ?", strings.TrimSpace(id)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}
