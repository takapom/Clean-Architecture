package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"

	"gorm.io/gorm"
)

type ReservationRepo struct{ db *gorm.DB }

func NewReservationRepo(db *gorm.DB) repository.ReservationRepository {
	return &ReservationRepo{db: db}
}

// DBのauto-incrementに委譲するのでNextIDは使わないが、interface満たすために実装
func (r *ReservationRepo) NextID() int { return 0 }

func (r *ReservationRepo) Save(res *entity.Reservation) (*entity.Reservation, error) {
	ctx := context.Background()
	m := models.ReservationModel{
		ID:       res.ID, // 0ならAUTO_INCREMENT
		UserID:   res.UserID,
		PlanID:   res.PlanID,
		Number:   res.Number,
		Checkin:  res.Checkin,
		Checkout: res.Checkout,
		Total:    res.Total,
	}
	if err := r.db.WithContext(ctx).Save(&m).Error; err != nil {
		return nil, err
	}
	// 生成されたIDを反映
	res.ID = m.ID
	return res, nil
}

func (r *ReservationRepo) FindByID(id int) (*entity.Reservation, error) {
	var m models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Reservation{
		ID:       m.ID,
		UserID:   m.UserID,
		PlanID:   m.PlanID,
		Number:   m.Number,
		Checkin:  m.Checkin,
		Checkout: m.Checkout,
		Total:    m.Total,
	}, nil
}

func (r *ReservationRepo) List() ([]*entity.Reservation, error) {
	var list []models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Reservation, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Reservation{
			ID:       copy.ID,
			UserID:   copy.UserID,
			PlanID:   copy.PlanID,
			Number:   copy.Number,
			Checkin:  copy.Checkin,
			Checkout: copy.Checkout,
			Total:    copy.Total,
		})
	}
	return out, nil
}

var _ repository.ReservationRepository = (*ReservationRepo)(nil)
