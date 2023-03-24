package account

import (
	"net/http"

	"github.com/ranefattesingh/bank-application-api/internal/router"
)

type Handlers interface {
	Routes() map[string]router.Route
	CreateAccount(http.ResponseWriter, *http.Request)
	GetAccount(http.ResponseWriter, *http.Request)
}
