package handler

import (
	"ariqt-auth/internal/models/apprepo"
	"ariqt-auth/pkg/server"
)

type Handler struct {
	apprepo *apprepo.AppRepo
	srv     *server.Server
}

func New(apprepo *apprepo.AppRepo, srv *server.Server) *Handler {
	return &Handler{
		apprepo: apprepo,
		srv:     srv,
	}
}
