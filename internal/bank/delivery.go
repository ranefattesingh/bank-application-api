package bank

import (
	"net/http"

	"github.com/ranefattesingh/bank-application-api/internal/router"
)

type Handlers interface {
	Routes() map[string]router.Route
	CreateBank(http.ResponseWriter, *http.Request)
	ListBanks(http.ResponseWriter, *http.Request)
	DeleteBlank(http.ResponseWriter, *http.Request)
	GetBank(http.ResponseWriter, *http.Request)
	UpdateBank(http.ResponseWriter, *http.Request)
}
