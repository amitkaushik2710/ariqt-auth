package handler

import (
	"ariqt-auth/internal/models"
	"ariqt-auth/pkg/auth"
	"ariqt-auth/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler godoc
// @Summary User Sign-up
// @Description Creates New User
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.AuthReq true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/signup [post]
func (h *Handler) SignUpHandler(c *gin.Context) {
	logger.Logger.Info("SignUpHandler: API Called")

	var req models.AuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("SignUpHandler: failed to bind payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signup parameters"})
		return
	}
	logger.Logger.Info("SignUpHandler", zap.Any("req", req))

	// Checking if user already exist with email in DB
	exUser, ok := h.apprepo.DBUser[req.Email]
	if ok {
		logger.Logger.Error("SignUpHandler: user already exist with email", zap.Any("existing user", exUser))
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist with email"})
		return
	}

	hashedPwd, err := auth.HashPassword(req.Password)
	if err != nil {
		logger.Logger.Error("SignUpHandler: failed to hash user password", zap.Any("req", req), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server failed to signup user"})
		return
	}

	newUser := models.User{
		Email:    req.Email,
		Password: hashedPwd,
	}

	h.apprepo.DBUser[newUser.Email] = newUser

	logger.Logger.Info("SignUpHandler", zap.Any("db user", h.apprepo.DBUser))

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "user created successfully"})
}
