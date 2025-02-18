package server

import (
	"ariqt-auth/internal/constants"
	"ariqt-auth/internal/handler"
	"ariqt-auth/internal/models/apprepo"
	"ariqt-auth/pkg/server"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	appRepo *apprepo.AppRepo
	srv     *server.Server
	handler *handler.Handler
}

func New(appRepo *apprepo.AppRepo) *Server {
	srv := server.New()
	handler := handler.New(appRepo, srv)
	return &Server{
		appRepo: appRepo,
		srv:     srv,
		handler: handler,
	}
}

func (s *Server) StartServer() {
	s.srv.StartPort(s.appRepo.Config.Port)
}

func (s *Server) Routes() {
	s.srv.Routes(http.MethodGet, constants.API_STATUS, s.handler.StatusHandler)
	s.srv.Routes(http.MethodPost, constants.API_SIGNUP, s.handler.SignUpHandler)
	s.srv.Routes(http.MethodPost, constants.API_SIGNIN, s.handler.SignInHandler)
	s.srv.Routes(http.MethodPost, constants.API_TEST_REQ, s.handler.TestReqHandler)
	s.srv.Routes(http.MethodDelete, constants.API_REVOKE_TOKEN, s.handler.RevokeTokenHandler)
	s.srv.Routes(http.MethodGet, constants.API_REFRESH_TOKEN, s.handler.RefreshTokenHandler)
	s.srv.Routes(http.MethodGet, "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
