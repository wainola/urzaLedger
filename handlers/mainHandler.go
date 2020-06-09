package handlers

import (
	"net/http"
	"urza/models"
)

type UrzaEnvironment struct {
	DB *models.UrzaDB
}

type UrzaApp struct {
	env *UrzaEnvironment
	fn  func(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error
}

func (urzaHandler UrzaApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := ctx.Err(); err != nil {
		return
	}

	urzaHandler.fn(urzaHandler.env, w, r)
}
