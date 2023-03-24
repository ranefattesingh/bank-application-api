package http

import (
	"net/http"

	"github.com/ranefattesingh/bank-application-api/internal/router"
)

func (h *bankHandle) Routes() map[string]router.Route {
	return map[string]router.Route{
		"CreateBank": {
			Method:      http.MethodPost,
			Path:        bankRouterPath,
			HandlerFunc: h.CreateBank,
		},
		"GetBank": {
			Method:      http.MethodGet,
			Path:        bankRouterPath + bankIDParam,
			HandlerFunc: h.GetBank,
		},
		"ListBank": {
			Method:      http.MethodGet,
			Path:        bankRouterPath,
			HandlerFunc: h.ListBank,
		},
		"UpdateBank": {
			Method:      http.MethodPut,
			Path:        bankRouterPath + bankIDParam,
			HandlerFunc: h.UpdateBank,
		},
		"DeleteBank": {
			Method:      http.MethodDelete,
			Path:        bankRouterPath + bankIDParam,
			HandlerFunc: h.DeleteBlank,
		},
	}
}
