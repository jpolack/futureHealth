package lib

import (
	"io/ioutil"
	"os"
)

func createFile(path string) {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func Read(path string) []byte {
	createFile(path)

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return fileContent
}

func Write(bytes []byte, path string) {
	createFile(path)
	err := ioutil.WriteFile(path, bytes, 0777)
	if err != nil {
		panic(err)
	}
}
