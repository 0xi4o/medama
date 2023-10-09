package services

import (
	"github.com/medama-io/medama/db/sqlite"
	"github.com/medama-io/medama/util"
)

type Handler struct {
	auth *util.AuthService
	db   *sqlite.Client
}

// NewService returns a new instance of the ogen service handler.
func NewService(auth *util.AuthService, db *sqlite.Client) *Handler {
	return &Handler{
		auth: auth,
		db:   db,
	}
}
