package request

import "github.com/Edward-Alphonse/saywo_pkg/utils"

type Headers struct {
	Authorization string `json:"authorization"`
}

func (h *Headers) JWT() string {
	return h.parseBearerAuth()
}

func (h *Headers) parseBearerAuth() (token string) {
	const prefix = "Bearer "
	auth := h.Authorization
	if len(auth) < len(prefix) || !utils.EqualFold(auth[:len(prefix)], prefix) {
		return ""
	}
	return auth[len(prefix):]
}
