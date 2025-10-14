package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"
	"strings"

	"gorm.io/gorm"
)

type PlanRepo struct{ db *gorm.DB }

func NewPlanRepo(db *gorm.DB) repository.PlanRepository { return &PlanRepo{db: db} }

func (r *PlanRepo) FindByID(id int) (*entity.Plan, error) {
	var m models.PlanModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Plan{ID: m.ID, Name: m.Name, Keyword: m.Keyword, Price: m.Price}, nil
}

func (r *PlanRepo) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	ctx := context.Background()
	var list []models.PlanModel

	q := r.db.WithContext(ctx).Model(&models.PlanModel{})
	if strings.TrimSpace(keyword) != "" {
		kw := "%" + strings.TrimSpace(keyword) + "%"
		q = q.Where("LOWER(name) LIKE LOWER(?) OR LOWER(keyword) LIKE LOWER(?)", kw, kw)
	}
	if err := q.Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Plan, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Plan{ID: copy.ID, Name: copy.Name, Keyword: copy.Keyword, Price: copy.Price})
	}
	return out, nil
}
