package user

import (
	"futureHealth/achievment"

	"github.com/google/uuid"
)

type User struct {
	Id          string                  `json:"id"`
	Achievments []achievment.Achievment `json:"achievments"`
}

type Persistence interface {
	read() map[string]User
	save(users map[string]User)
}

type UserHandler struct {
	Pers Persistence
}

type LoginToken struct {
	Token string `json:"token"`
}

func (h *UserHandler) Create() LoginToken {
	id := uuid.New().String()
	users := h.Pers.read()
	users[id] = User{
		Id:          id,
		Achievments: []achievment.Achievment{},
	}
	h.Pers.save(users)
	return LoginToken{id}
}

type Point struct {
	Count int `json:"count"`
}

func (h *UserHandler) Points(userId string) Point {
	users := h.Pers.read()
	user, found := users[userId]
	points := 0
	if !found {
		return Point{}
	}

	for _, achiev := range user.Achievments {
		points += achiev.Points
	}
	return Point{points}
}

func (h *UserHandler) UserAchieved() []achievment.Achievment {
	return []achievment.Achievment{}
}
