package routers

import (
	"github.com/NurbekDos/funk/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetIssuerRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1/iss")

	v1.POST("/login", handlers.IssuerLogin)
}
