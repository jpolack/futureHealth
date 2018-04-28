package user

import (
	"encoding/json"
	"io/ioutil"
)

func CreateJsonPersistence(path string) jsonPersistence {
	return jsonPersistence{path}
}

type jsonPersistence struct {
	path string
}

func (pers *jsonPersistence) read() []User {
	f, err := ioutil.ReadFile(pers.path)
	if err != nil {
		err = ioutil.WriteFile(pers.path, []byte("[]"), 0777)
		if err != nil {
			panic(err)
		}
		f, err = ioutil.ReadFile(pers.path)
		if err != nil {
			panic(err)
		}
	}

	achievs := []User{}
	err = json.Unmarshal(f, &achievs)
	if err != nil {
		panic(err)
	}
	return achievs
}

func (pers *jsonPersistence) save(achievments []User) {
	bytes, err := json.Marshal(achievments)

	err = ioutil.WriteFile(pers.path, bytes, 0777)
	if err != nil {
		panic(err)
	}
}
