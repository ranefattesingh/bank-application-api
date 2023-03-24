package http

import (
	"net/http"
)

const (
	accountRouterPath = "/account"
	accountIDParam    = "/{accountId}"
	bankIDParam       = "/{bankId}"
)

type accountHandle struct {
}

func NewAccountHandler() *accountHandle {
	return &accountHandle{}
}

func (h *accountHandle) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateAccount"))
}

func (h *accountHandle) GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAccount"))
}
