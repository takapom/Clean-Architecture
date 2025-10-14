package httpi

import (
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	UC *usecase.UserUsecase
}

type registerUserReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}

type registerUserResp struct {
	ID string `json:"id"`
}

type userView struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	DateOfBirth  string `json:"date_of_birth"`
	RegisteredAt string `json:"registered_at"`
	Status       string `json:"status"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user lookup unavailable", http.StatusServiceUnavailable)
		return
	}

	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/users/"))
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.UC.GetUser(id)
	if err != nil {
		if errors.Is(err, usecase.ErrUserInvalidInput) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	view := userView{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  formatDate(user.DateOfBirth),
		RegisteredAt: user.RegisteredAt.Format(time.RFC3339),
		Status:       user.Status,
	}

	writeJSON(w, http.StatusOK, view)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user registration unavailable", http.StatusServiceUnavailable)
		return
	}

	var in registerUserReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.UC.Register(usecase.RegisterUserInput{
		Name:        in.Name,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		DateOfBirth: in.DateOfBirth,
	})
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserInvalidInput):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserEmailAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	writeJSON(w, http.StatusCreated, registerUserResp{ID: user.ID})
}

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}
