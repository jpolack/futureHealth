package api

import (
	"fmt"
	"testing"
)

func Test_Login_CorrectUser_Passed(t *testing.T) {
	_, err := ApiLogin("g3483706@nwytg.com", "123456789")
	if err != nil {
		t.Error(err)
	}
}

func Test_GetExercises(t *testing.T) {
	session, _ := ApiLogin("g3483706@nwytg.com", "123456789")
	act, err := GetExercises(session)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Result %#v\n", act)
}
