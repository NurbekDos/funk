package routers

import (
	"github.com/NurbekDos/funk/internal/handlers"
	"github.com/NurbekDos/funk/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
	v1.GET("/me", middlewares.AuthMiddleware(), handlers.Me)
}
