package lib

import (
	"encoding/json"
)

func CreateAchievmentPersistence(path string) achievmentPersistence {
	return achievmentPersistence{path}
}

type achievmentPersistence struct {
	path string
}

func (pers *achievmentPersistence) read() []Achievment {
	bytes := Read(pers.path)
	achievs := []Achievment{}

	if len(bytes) == 0 {
		pers.save(achievs)
		return achievs
	}

	err := json.Unmarshal(bytes, &achievs)
	if err != nil {
		panic(err)
	}
	return achievs
}

func (pers *achievmentPersistence) save(users []Achievment) {
	bytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	Write(bytes, pers.path)
}
