package user

import (
	"futureHealth/achievment"
	"futureHealth/api"

	runtasticAPI "github.com/Metalnem/runtastic/api"
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
type Runtastic interface {
	ApiLogin(username string, password string) (*runtasticAPI.Session, error)
	GetExercises(session *runtasticAPI.Session) ([]api.Exercise, error)
}

type UserHandler struct {
	Pers   Persistence
	RunApi Runtastic
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

type Progress struct {
	achievment.Achievment
	progress float64
}

func (h *UserHandler) UserAchieved(achievments []achievment.Achievment) []Progress {
	session, err := h.RunApi.ApiLogin("g3483706@nwytg.com", "123456789")
	if err != nil {
		panic(err)
	}
	exercise, err := h.RunApi.GetExercises(session)
	if err != nil {
		panic(err)
	}

	achieved := make([]Progress, len(achievments))
	for i, achievs := range achievments {
		prog := Progress{achievs, 0.0}
		for _, ex := range exercise {
			if achievs.Type == ex.Type {
				switch achievs.Unit {
				case "Kilometers":
					prog.progress += float64(ex.Distance)
				case "Calories":
					prog.progress += float64(ex.Calories)
				case "Minutes":
					prog.progress += ex.Duration
				}
			}
		}
		achieved[i] = prog
	}
	return achieved
}
