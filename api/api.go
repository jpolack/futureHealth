package api

import (
	"context"

	"github.com/Metalnem/runtastic/api"
)

type Exercise struct {
	Calories int32
	Distance int32
	Duration float64
	Type     string
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
	availEx := map[string]int{"Running": 0}

	for _, act := range activities {
		if _, ok := availEx[act.Type]; ok {
			exercises = append(exercises, Exercise{
				Calories: act.Calories,
				Distance: act.Distance,
				Duration: act.Duration.Seconds(),
				Type:     act.Type,
			})
		}
	}

	return exercises, nil
}
