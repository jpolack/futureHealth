package lib

import (
	"fmt"
	"os"
	"testing"
)

func removeFile() {
	err := os.Remove("../data/test.json")
	if err != nil {
		panic(err)
	}
}

func TestCreate(t *testing.T) {
	removeFile()
	createFile("../data/test.json")
}
func TestRead(t *testing.T) {
	removeFile()
	content := Read("../data/test.json")
	if len(content) != 0 {
		fmt.Println("File is not empty")
		t.Fail()
	}
}
func TestWrite(t *testing.T) {
	removeFile()
	Write([]byte("abc"), "../data/test.json")
}
func TestReadWrite(t *testing.T) {
	removeFile()
	Write([]byte("abc"), "../data/test.json")
	content := Read("../data/test.json")

	if string(content[:]) != "abc" {
		fmt.Println("File is not empty")
		t.Fail()
	}
}
