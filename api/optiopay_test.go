package api

import (
	"fmt"
	"testing"
)

func Test_CreatePayment(t *testing.T) {
	op := optioPayment{}
	created := op.CreatePayment()

	fmt.Printf("Result %#v\n", created)
}
