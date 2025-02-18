package handler

import (
	"ariqt-auth/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusHandler godoc
// @Summary Status Handler
// @Description Responds with "version"
// @Tags Status
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/status [get]
func (h *Handler) StatusHandler(c *gin.Context) {
	logger.Logger.Info("StatusHandler: API Called")
	c.JSON(http.StatusOK, gin.H{"version": "v1"})
}
