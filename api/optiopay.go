package api

import (
	"net/http"
	"os"
	"strings"
)

type optioPayment struct{}

func (op *optioPayment) CreatePayment() bool {
	file, err := os.OpenFile("bankaccount.json", os.O_RDWR, 0777)

	if err != nil {
		return false
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.sandbox.optest.de/issuer/v1/payment-jobs", file)
	req.Header.Add("Authorization", "Bearer fue542x4kcbua7nblusfoiio2ljal2b4")

	res, err := client.Do(req)
	if err != nil || !strings.Contains(res.Status, "201") {
		return false
	}
	return true
}
