package main

import (
	"ariqt-auth/config"
	_ "ariqt-auth/docs"
	"ariqt-auth/internal/models"
	"ariqt-auth/internal/models/apprepo"
	"ariqt-auth/internal/server"
	"ariqt-auth/pkg/logger"

	"go.uber.org/zap"
)

// @title Amit Kaushik Auth API-Assignment
// @version 1.0
// @description an Auth REST API service using Swagger with Gin-Gonic.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:7001
// @BasePath /api/v1
func main() {
	logger.InitializeLogger()
	logger.Logger.Info("main: logger initialized")

	appRepo := apprepo.AppRepo{}
	appRepo.DBUser = make(map[string]models.User)
	appRepo.AuthToken = make(map[string]models.User)

	config := config.GetConfig()
	appRepo.Config = config
	logger.Logger.Info("main: config loaded", zap.Any("config", appRepo.Config))

	server := server.New(&appRepo)
	server.Routes()
	server.StartServer()
}
