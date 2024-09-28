package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/NurbekDos/funk/internal/models"
	"github.com/NurbekDos/funk/internal/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 {
			log.Println("AuthMiddleware: Token parts")
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		claims := services.VerifyToken(tokenParts[1])
		if claims.ExpiresAt < time.Now().Unix() {
			log.Println("AuthMiddleware: Expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		user := models.User{
			ID:    claims.UserId,
			Email: claims.Email,
		}
		c.Set("user", user)
	}
}
