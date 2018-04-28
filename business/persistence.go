package business

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

func (pers *jsonPersistence) read() []Achievment {
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

	achievs := []Achievment{}
	err = json.Unmarshal(f, &achievs)
	if err != nil {
		panic(err)
	}
	return achievs
}

func (pers *jsonPersistence) save(achievments []Achievment) {
	bytes, err := json.Marshal(achievments)

	err = ioutil.WriteFile(pers.path, bytes, 0777)
	if err != nil {
		panic(err)
	}
}
