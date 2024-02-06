package controllers

import (
	"net/http"

	"github.com/ono5/jwt-with-golang/utils"
)

type Controller struct{}

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "Yes")
	}
}
