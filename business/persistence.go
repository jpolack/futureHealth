package business

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CreateJsonPersistence(path string) jsonPersistence {
	err := ioutil.WriteFile(path, []byte("[]"), 0777)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	return jsonPersistence{f, path}
}

type jsonPersistence struct {
	file *os.File
	path string
}

func (pers *jsonPersistence) read() []Achievment {
	f, err := ioutil.ReadFile(pers.path)
	if err != nil {
		panic(err)
	}

	achievs := []Achievment{}
	err = json.Unmarshal(f, &achievs)
	if err != nil {
		panic(err)
	}
	return achievs
}

func (pers *jsonPersistence) save(achievments []Achievment) {
	bytes, err := json.Marshal(achievments)

	_, err = pers.file.Write(bytes)
	if err != nil {
		panic(err)
	}
}
