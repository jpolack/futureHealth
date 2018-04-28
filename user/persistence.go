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

func (pers *jsonPersistence) read() map[string]User {
	f, err := ioutil.ReadFile(pers.path)
	if err != nil {
		err = ioutil.WriteFile(pers.path, []byte("{}"), 0777)
		if err != nil {
			panic(err)
		}
		f, err = ioutil.ReadFile(pers.path)
		if err != nil {
			panic(err)
		}
	}

	users := make(map[string]User)
	err = json.Unmarshal(f, &users)
	if err != nil {
		panic(err)
	}
	return users
}

func (pers *jsonPersistence) save(users map[string]User) {
	bytes, err := json.Marshal(users)

	err = ioutil.WriteFile(pers.path, bytes, 0777)
	if err != nil {
		panic(err)
	}
}
