package http

import (
	"net/http"
)

const (
	bankRouterPath = "/bank"
	bankIDParam    = "/{bankId}"
)

type bankHandle struct {
}

func NewBankHandler() *bankHandle {
	return &bankHandle{}
}

func (h *bankHandle) CreateBank(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateBank"))
}

func (h *bankHandle) ListBank(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ListBank"))
}

func (h *bankHandle) DeleteBlank(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteBank"))
}

func (h *bankHandle) GetBank(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBank"))
}

func (h *bankHandle) UpdateBank(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateBank"))
}
