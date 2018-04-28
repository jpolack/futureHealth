package login

import "github.com/google/uuid"

type LoginToken struct {
	Token string `json:"token"`
}

func Login() LoginToken {
	return LoginToken{uuid.New().String()}
}
