package httpi

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ReservationHandler struct {
	UC *usecase.ReservationUsecase
}

type createReq struct {
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`  // "2025-10-12"
	Checkout string `json:"checkout"` // "2025-10-13"
}

type createResp struct {
	ID int `json:"id"`
}

type reservationView struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Total    int    `json:"total"`
	Nights   int    `json:"nights"`
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in createReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	ci, err1 := time.Parse("2006-01-02", in.Checkin)
	co, err2 := time.Parse("2006-01-02", in.Checkout)
	if err1 != nil || err2 != nil {
		http.Error(w, "invalid date format (yyyy-mm-dd)", http.StatusBadRequest)
		return
	}
	res, err := h.UC.Create(in.UserID, in.PlanID, in.Number, ci, co)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, usecase.ErrInvalidDates):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrInvalidNumber):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrPlanNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	writeJSON(w, http.StatusOK, createResp{ID: res.ID})
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reservations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, _ := h.UC.Get(id)
	if res == nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, toView(res))
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	list, _ := h.UC.List()
	views := make([]reservationView, 0, len(list))
	for _, v := range list {
		views = append(views, toView(v))
	}
	writeJSON(w, http.StatusOK, views)
}

func (h *ReservationHandler) SearchPlans(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("keyword")
	plans, _ := h.UC.SearchPlans(q)
	type planView struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Keyword string `json:"keyword"`
		Price   int    `json:"price"`
	}
	out := make([]planView, 0, len(plans))
	for _, p := range plans {
		out = append(out, planView{ID: p.ID, Name: p.Name, Keyword: p.Keyword, Price: p.Price})
	}
	writeJSON(w, http.StatusOK, out)
}

// ここを *entity.Reservation にする（別型を作らない）
func toView(r *entity.Reservation) reservationView {
	return reservationView{
		ID:       r.ID,
		UserID:   r.UserID,
		PlanID:   r.PlanID,
		Number:   r.Number,
		Checkin:  r.Checkin.Format("2006-01-02"),
		Checkout: r.Checkout.Format("2006-01-02"),
		Total:    r.Total,
		Nights:   r.Nights(), // entity に既にあるメソッドを使う
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
