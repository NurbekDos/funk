package routers

import (
	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/handlers"
	"github.com/NurbekDos/funk/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetAdminRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1/adm")

	v1.POST("/login", handlers.AdminLogin)

	v1.POST("/create", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminCreateAdmin)
	v1.POST("/issuers/create", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminCreateIssuer)

	v1.GET("/cases", handlers.Cases)
	v1.POST("/cases/create", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminCreateCase)
	v1.POST("/cases/update", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminUpdateCase)
	v1.POST("/cases/delete", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminDeleteCase)

	v1.GET("/tokens", handlers.AdminGetToken)
	v1.POST("/tokens/create", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminCreateToken)
	v1.POST("/tokens/update", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminUpdateToken)
	v1.POST("/tokens/delete", middlewares.AuthMiddleware(cfg.UserType_Admin), handlers.AdminDeleteToken)
}
