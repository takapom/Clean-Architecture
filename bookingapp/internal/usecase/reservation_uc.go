package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidDates  = errors.New("invalid dates: checkout must be after checkin")
	ErrPlanNotFound  = errors.New("plan not found")
	ErrInvalidNumber = errors.New("number must be >= 1")
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user id")
)

type ReservationUsecase struct {
	Plans repository.PlanRepository
	Resv  repository.ReservationRepository
	Users repository.UserRepository
}

// 　予約作成
func (u *ReservationUsecase) Create(userID string, planID, number int, checkin, checkout time.Time) (*entity.Reservation, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, ErrInvalidUserID
	}
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}
	user, err := u.Users.Get(strings.TrimSpace(userID))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if !checkout.After(checkin) {
		return nil, ErrInvalidDates
	}
	if number < 1 {
		return nil, ErrInvalidNumber
	}
	plan, err := u.Plans.FindByID(planID)
	if err != nil || plan == nil {
		return nil, ErrPlanNotFound
	}
	r := &entity.Reservation{
		ID:       u.Resv.NextID(),
		UserID:   strings.TrimSpace(userID),
		PlanID:   planID,
		Number:   number,
		Checkin:  checkin,
		Checkout: checkout,
	}
	//ドメイン層のメソッドを使って宿泊数を計算
	nights := r.Nights()
	//合計金額を計算してセット
	r.Total = plan.Price * number * nights
	//保存してID付きの予約情報を返す
	return u.Resv.Save(r)
}

// 予約取得
func (u *ReservationUsecase) Get(id int) (*entity.Reservation, error) {
	return u.Resv.FindByID(id)
}

// 予約一覧取得
func (u *ReservationUsecase) List() ([]*entity.Reservation, error) {
	return u.Resv.List()
}

// プラン検索
func (u *ReservationUsecase) SearchPlans(keyword string) ([]*entity.Plan, error) {
	return u.Plans.SearchByKeyword(keyword)
}
