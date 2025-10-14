package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"strings"
)

type PlanRepoMemory struct {
	data map[int]*entity.Plan
}

func NewPlanRepoMemory(seed []*entity.Plan) repository.PlanRepository {
	m := &PlanRepoMemory{data: map[int]*entity.Plan{}}
	for _, p := range seed {
		cp := *p
		m.data[p.ID] = &cp
	}
	return m
}

func (m *PlanRepoMemory) FindByID(id int) (*entity.Plan, error) {
	if p, ok := m.data[id]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, nil
}

func (m *PlanRepoMemory) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	if keyword == "" {
		out := make([]*entity.Plan, 0, len(m.data))
		for _, p := range m.data {
			cp := *p
			out = append(out, &cp)
		}
		return out, nil
	}
	kw := strings.ToLower(keyword)
	var out []*entity.Plan
	for _, p := range m.data {
		if strings.Contains(strings.ToLower(p.Name), kw) || strings.Contains(strings.ToLower(p.Keyword), kw) {
			cp := *p
			out = append(out, &cp)
		}
	}
	return out, nil
}
