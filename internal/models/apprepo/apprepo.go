package apprepo

import (
	"ariqt-auth/config"
	"ariqt-auth/internal/models"
)

type AppRepo struct {
	Config    *config.Config
	DBUser    map[string]models.User
	AuthToken map[string]models.User
}
