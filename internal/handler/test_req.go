package handler

import (
	"ariqt-auth/internal/models"
	"ariqt-auth/pkg/auth"
	"ariqt-auth/pkg/helpers"
	"ariqt-auth/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TestReqHandler godoc
// @Summary Test Req Handler
// @Description Test Req
// @Tags Test Req Data
// @Produce json
// @Param credentials body models.TestReq true "User req"
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/test-req [post]
func (h *Handler) TestReqHandler(c *gin.Context) {
	logger.Logger.Info("TestReqHandler: API Called")

	authToken := helpers.GetAuthToken(c)

	var req models.TestReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("TestReqHandler: failed to bind payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid test req parameters"})
		return
	}
	logger.Logger.Info("TestReqHandler", zap.Any("req", req))

	// Verfying Token
	err := auth.VerifyToken(authToken)
	if err != nil {
		if err.Error() == "Token is expired" {
			logger.Logger.Error("TestReqHandler: token is expired", zap.Any("req", req), zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			return
		}
		logger.Logger.Error("TestReqHandler: failed to validate token", zap.Any("req", req), zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to validate token"})
		return
	}

	// Checking for Auth Token
	_, ok := h.apprepo.AuthToken[authToken]
	if !ok {
		logger.Logger.Error("TestReqHandler: token does not exist in system", zap.Any("req", req))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth token does not exist in system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "data": req.Data})
}
