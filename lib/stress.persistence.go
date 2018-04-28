package lib

import (
	"encoding/json"
)

func CreateStressPersistence(path string) stressPersistence {
	return stressPersistence{path}
}

type stressPersistence struct {
	path string
}

func (pers *stressPersistence) read() map[int]StressLevel {
	bytes := Read(pers.path)
	users := make(map[int]StressLevel)

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

func (pers *stressPersistence) save(users map[int]StressLevel) {
	bytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	Write(bytes, pers.path)
}
