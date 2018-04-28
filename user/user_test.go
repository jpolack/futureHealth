package user

import (
	"fmt"
	"testing"
)

func Test_Login_GetUuid(t *testing.T) {
	token := Login()
	fmt.Println(token)
}
