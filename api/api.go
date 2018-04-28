package api

import (
	"context"

	"github.com/Metalnem/runtastic/api"
)

type Exercise struct {
	Calories int32
	Distance int32
	Duration float64
}

func ApiLogin(username string, password string) (*api.Session, error) {
	return api.Login(context.Background(), username, password)
}

func GetExercises(session *api.Session) ([]Exercise, error) {
	activities, err := session.GetActivities(context.Background())
	if err != nil {
		return nil, err
	}

	exercises := []Exercise{}

	for _, act := range activities {
		if act.Type == "Running" {
			exercises = append(exercises, Exercise{
				Calories: act.Calories,
				Distance: act.Distance,
				Duration: act.Duration.Seconds(),
			})
		}
	}

	return exercises, nil
}
