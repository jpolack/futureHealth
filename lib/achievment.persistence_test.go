package lib

import (
	"fmt"
	"testing"
)

func TestAchievWriteRead(t *testing.T) {
	removeFile()
	w := CreateAchievmentPersistence("../data/test.json")

	achievs := []Achievment{
		Achievment{},
	}

	w.save(achievs)

	a := w.read()

	if len(a) != 1 {
		fmt.Println("Expected len of 1")
		t.Fail()
	}
}
func TestAchievEmptyRead(t *testing.T) {
	removeFile()
	w := CreateAchievmentPersistence("../data/test.json")

	a := w.read()

	if len(a) != 0 {
		fmt.Println("Expected len of > 0")
		t.Fail()
	}
}
