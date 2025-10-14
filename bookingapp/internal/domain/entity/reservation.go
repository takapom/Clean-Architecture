package entity

import "time"

type Reservation struct {
	ID       int
	UserID   string
	PlanID   int
	Number   int
	Checkin  time.Time
	Checkout time.Time
	Total    int // 計算済み合計金額
}

func (r *Reservation) Nights() int {
	d := r.Checkout.Sub(r.Checkin).Hours() / 24
	if d < 0 {
		return 0
	}
	return int(d)
}
