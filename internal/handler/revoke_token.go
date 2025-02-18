package handler

import (
	"ariqt-auth/pkg/auth"
	"ariqt-auth/pkg/helpers"
	"ariqt-auth/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RevokeTokenHandler godoc
// @Summary Revoke Token Handler
// @Description Revoke Token
// @Tags Revoke Token
// @Produce json
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/revoke-token [delete]
func (h *Handler) RevokeTokenHandler(c *gin.Context) {
	logger.Logger.Info("RevokeTokenHandler: API Called")

	authToken := helpers.GetAuthToken(c)

	// Verfying Token
	err := auth.VerifyToken(authToken)
	if err != nil {
		if err.Error() == "Token is expired" {
			logger.Logger.Error("RevokeTokenHandler: token is expired", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			return
		}
		logger.Logger.Error("RevokeTokenHandler: failed to validate token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to validate token"})
		return
	}

	// Checking for Auth Token
	_, ok := h.apprepo.AuthToken[authToken]
	if !ok {
		logger.Logger.Error("RevokeTokenHandler: token does not exist in system")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth token does not exist in system"})
		return
	}

	// Delete token from cache
	delete(h.apprepo.AuthToken, authToken)

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "token revoked"})
}
