package repository

import "bookingapp/internal/domain/entity"

type PlanRepository interface {
	FindByID(id int) (*entity.Plan, error)
	SearchByKeyword(keyword string) ([]*entity.Plan, error)
}

type ReservationRepository interface {
	NextID() int
	Save(reservation *entity.Reservation) (*entity.Reservation, error)
	FindByID(id int) (*entity.Reservation, error)
	List() ([]*entity.Reservation, error)
}

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Get(id string) (*entity.User, error)
}
