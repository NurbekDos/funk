package routers

import (
	"github.com/NurbekDos/funk/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
}
