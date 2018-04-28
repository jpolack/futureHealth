package api

import (
	"fmt"
	"testing"
)

func tTest_Login_CorrectUser_Passed(t *testing.T) {
	api := RuntasticApi{}
	_, err := api.ApiLogin("g3483706@nwytg.com", "123456789")
	if err != nil {
		t.Error(err)
	}
}

func tTest_GetExercises(t *testing.T) {
	api := RuntasticApi{}
	session, _ := api.ApiLogin("g3483706@nwytg.com", "123456789")
	act, err := api.GetExercises(session)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Result %#v\n", act)
}
