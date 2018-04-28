package login

import (
	"fmt"
	"testing"
)

func Test_Login_GetUuid(t *testing.T) {
	uid := Login()
	fmt.Println(uid)
}
