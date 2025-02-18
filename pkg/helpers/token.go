package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) string {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return ""
	}

	// Check if it's a Bearer token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return ""
	}

	token := tokenParts[1] // Extract the actual token
	fmt.Println("Extracted Token:", token)

	return token
}
