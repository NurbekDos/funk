package routers

import (
	"github.com/NurbekDos/funk/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetAdminRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1/adm")

	v1.POST("/login", handlers.AdminLogin)
}
