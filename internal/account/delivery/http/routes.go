package http

import (
	"net/http"

	"github.com/ranefattesingh/bank-application-api/internal/router"
)

func (h *accountHandle) Routes() map[string]router.Route {
	return map[string]router.Route{
		"CreateAccount": {
			Method:      http.MethodPost,
			Path:        bankIDParam + accountRouterPath,
			HandlerFunc: h.CreateAccount,
		},
		"GetAccount": {
			Method:      http.MethodGet,
			Path:        bankIDParam + accountRouterPath + accountIDParam,
			HandlerFunc: h.GetAccount,
		},
	}
}
