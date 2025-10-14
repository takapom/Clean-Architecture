package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"sync"
)

type ReservationRepoMemory struct {
	mu   sync.RWMutex
	data map[int]*entity.Reservation
	next int
}

func NewReservationRepoMemory() repository.ReservationRepository {
	return &ReservationRepoMemory{
		data: make(map[int]*entity.Reservation),
		next: 1,
	}
}

func (r *ReservationRepoMemory) NextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.next
	r.next++
	return id
}

func (r *ReservationRepoMemory) Save(res *entity.Reservation) (*entity.Reservation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *res
	r.data[cp.ID] = &cp
	out := cp
	return &out, nil
}

func (r *ReservationRepoMemory) FindByID(id int) (*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.data[id]; ok {
		cp := *v
		return &cp, nil
	}
	return nil, nil
}

func (r *ReservationRepoMemory) List() ([]*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*entity.Reservation, 0, len(r.data))
	for _, v := range r.data {
		cp := *v
		out = append(out, &cp)
	}
	return out, nil
}
