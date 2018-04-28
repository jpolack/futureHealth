package lib

import (
	"time"
)

type StressLevel struct {
	Percent float64   `json:"percent"`
	Date    time.Time `json:"date"`
}

type StressPersistence interface {
	read() map[int]StressLevel
	save(map[int]StressLevel)
}

type StressLevelHandler struct {
	Pers StressPersistence
}

func (s *StressLevelHandler) Create(stress StressLevel, userId string) {

}
