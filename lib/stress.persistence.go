package lib

//DEPRECATED

import (
	"encoding/json"
)

func CreateStressPersistence(path string) stressPersistence {
	return stressPersistence{path}
}

type stressPersistence struct {
	path string
}

func (pers *stressPersistence) read() []StressLevel {
	bytes := Read(pers.path)
	users := []StressLevel{}

	if len(bytes) == 0 {
		pers.save(users)
		return users
	}

	err := json.Unmarshal(bytes, &users)
	if err != nil {
		panic(err)
	}
	return users
}

func (pers *stressPersistence) save(users []StressLevel) {
	bytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	Write(bytes, pers.path)
}
