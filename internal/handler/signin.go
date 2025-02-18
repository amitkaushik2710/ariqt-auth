package handler

import (
	"ariqt-auth/internal/models"
	"ariqt-auth/pkg/auth"
	"ariqt-auth/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignInHandler godoc
// @Summary User Sign-in
// @Description Authenticates a user and returns access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.AuthReq true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/signin [post]
func (h *Handler) SignInHandler(c *gin.Context) {
	logger.Logger.Info("SignInHandler: API Called")

	var req models.AuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("SignInHandler: failed to bind payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signin parameters"})
		return
	}
	logger.Logger.Info("SignInHandler", zap.Any("req", req))

	// Checking if user already exist with email in DB
	user, ok := h.apprepo.DBUser[req.Email]
	if !ok {
		logger.Logger.Error("SignInHandler: user does not exist in db", zap.Any("req", req))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not exist in system"})
		return
	}

	chk := auth.VerifyPassword(req.Password, user.Password)
	if !chk {
		logger.Logger.Error("SignInHandler: invalid user password", zap.Any("req", req))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user password"})
		return
	}

	authToken, exp, err := auth.CreateToken(user.Email)
	if err != nil {
		logger.Logger.Error("SignInHandler: failed to generate auth token for the user", zap.Any("user", user), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server failed to sign in user"})
		return
	}

	h.apprepo.AuthToken[authToken] = user

	logger.Logger.Info("SignInHandler", zap.Any("db user", h.apprepo.DBUser))
	logger.Logger.Info("SignInHandler", zap.Any("auth token", h.apprepo.AuthToken))

	c.JSON(http.StatusOK, gin.H{"status": 200, "token": authToken, "exp": exp})
}
