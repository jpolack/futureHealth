package lib

import (
	"encoding/json"
)

func CreateUserPersistence(path string) userPersistence {
	return userPersistence{path}
}

type userPersistence struct {
	path string
}

func (pers *userPersistence) read() map[string]User {
	bytes := Read(pers.path)
	users := make(map[string]User)

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

func (pers *userPersistence) save(users map[string]User) {
	bytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	Write(bytes, pers.path)
}
