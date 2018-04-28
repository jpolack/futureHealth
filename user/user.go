package user

import (
	"futureHealth/achievment"

	"github.com/google/uuid"
)

type LoginToken struct {
	Token string `json:"token"`
}

func Login() LoginToken {
	return LoginToken{uuid.New().String()}
}

type User struct {
	Id          string                  `json:"id"`
	Achievments []achievment.Achievment `json:"achievments"`
}

type Persistence interface {
	read() []User
	save([]User)
}

type UserHandler struct {
	Pers Persistence
}

func (h *UserHandler) Points(userId string) int {
	//TODO
	return 0
}

func (h *UserHandler) UserAchieved() []achievment.Achievment {
	return []achievment.Achievment{}
}
