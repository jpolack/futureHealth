package lib

//DEPRECATED

import (
	"time"
)

type StressLevel struct {
	Percent float64   `json:"percent"`
	Date    time.Time `json:"date"`
}

type StressPersistence interface {
	read() []StressLevel
	save([]StressLevel)
}

type StressLevelHandler struct {
	Pers StressPersistence
}

func (s *StressLevelHandler) Create(stress StressLevel, userId string) {
	levels := s.Pers.read()
	levels = append(levels, stress)
	s.Pers.save(levels)
}
func (s *StressLevelHandler) Read() []StressLevel {
	levels := s.Pers.read()
	return levels
}
