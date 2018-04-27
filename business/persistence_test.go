package business

import (
	"fmt"
	"testing"
)

func TestWriteRead(t *testing.T) {
	w := CreateJsonPersistence("../data/test.json")

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
func TestEmptyRead(t *testing.T) {
	w := CreateJsonPersistence("../data/test.json")

	a := w.read()

	if len(a) != 0 {
		fmt.Println("Expected len of 0")
		t.Fail()
	}
}
