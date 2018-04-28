package achievment

import (
	"fmt"
	"os"
	"testing"
)

func beforeEach() {
	err := os.Remove("../data/test.json")
	if err != nil {
		panic(err)
	}
}

func TestWriteRead(t *testing.T) {
	beforeEach()
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
	beforeEach()
	w := CreateJsonPersistence("../data/test.json")

	a := w.read()

	if len(a) != 0 {
		fmt.Println("Expected len of 0")
		t.Fail()
	}
}
