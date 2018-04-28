package lib

import (
	"futureHealth/api"

	runtasticAPI "github.com/Metalnem/runtastic/api"
	"github.com/google/uuid"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id          string                `json:"id"`
	Achievments map[string]Achievment `json:"achievments"`
	Runtastic   Credentials           `json:"-"`
}

type UserPersistence interface {
	read() map[string]User
	save(users map[string]User)
}
type Runtastic interface {
	ApiLogin(username string, password string) (*runtasticAPI.Session, error)
	GetExercises(session *runtasticAPI.Session) ([]api.Exercise, error)
}

type UserHandler struct {
	Pers   UserPersistence
	RunApi Runtastic
}

func (h *UserHandler) Create() string {
	id := uuid.New().String()
	users := h.Pers.read()
	users[id] = User{
		Id:          id,
		Achievments: make(map[string]Achievment),
	}
	h.Pers.save(users)
	return id
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
	Achievment
	Progress float64 `json:"progress"`
}

func (h *UserHandler) UserAchieved(achievments []Achievment, userId string) []Progress {
	session, err := h.RunApi.ApiLogin("g3483706@nwytg.com", "123456789")
	if err != nil {
		panic(err)
	}
	exercise, err := h.RunApi.GetExercises(session)
	if err != nil {
		panic(err)
	}

	users := h.Pers.read()
	user := users[userId]
	achieved := []Progress{}
	for _, achiev := range achievments {
		if _, found := user.Achievments[achiev.Id]; found {
			continue
		}
		prog := Progress{achiev, 0.0}
		for _, ex := range exercise {
			if achiev.Type == ex.Type {
				switch achiev.Unit {
				case "Kilometers":
					prog.Progress += float64(ex.Distance)
				case "Calories":
					prog.Progress += float64(ex.Calories)
				case "Minutes":
					prog.Progress += ex.Duration
				}
			}
		}
		achieved = append(achieved, prog)
		if prog.Progress >= prog.Value {
			user.Achievments[achiev.Id] = achiev
			users[userId] = user
		}
	}
	h.Pers.save(users)

	return achieved
}
func (h *UserHandler) RuntasticLogin(cred Credentials, userId string) error {
	// _, err := h.RunApi.ApiLogin(cred.Username, cred.Password)
	// if err != nil {
	// 	return err
	// }

	// users := h.Pers.read()
	// foundUser, found := users[userId]
	// if !found {
	// 	return errors.New("user not found")
	// }

	// foundUser.Runtastic = cred

	// users[userId] = foundUser

	// h.Pers.save(users)

	// return nil
	return nil
}
