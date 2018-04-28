package lib

import "github.com/google/uuid"

type Achievment struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Value       float64 `json:"value"`
	Points      int     `json:"points"`
	Unit        string  `json:"unit"`
}

type Persistence interface {
	read() []Achievment
	save([]Achievment)
}

type AchievmentHandler struct {
	Pers Persistence
}

func (h *AchievmentHandler) Create(achiev Achievment) {
	achiev.Id = uuid.New().String()
	achievs := h.Pers.read()
	achievs = append(achievs, achiev)
	h.Pers.save(achievs)
}

func (h *AchievmentHandler) Read() []Achievment {
	return h.Pers.read()
}
