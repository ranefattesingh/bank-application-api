package main

import (
	"net/http"

	accountHttp "github.com/ranefattesingh/bank-application-api/internal/account/delivery/http"
	bankHttp "github.com/ranefattesingh/bank-application-api/internal/bank/delivery/http"
	"github.com/ranefattesingh/bank-application-api/internal/router"
)

func main() {
	accountHandler := accountHttp.NewAccountHandler()
	bankHandler := bankHttp.NewBankHandler()

	router := router.NewRouter(bankHandler, accountHandler)
	http.ListenAndServe("localhost:8080", router)
}
