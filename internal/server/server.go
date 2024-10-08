package server

import (
	"log"
	"time"

	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/NurbekDos/funk/internal/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Engine() {
	engine := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	engine.Use(cors.New(config))

	routers.SetUserRoutes(engine)
	routers.SetAdminRouter(engine)
	routers.SetIssuerRouter(engine)

	err := engine.Run(":" + cfg.GetConfig().Port)
	if err != nil {
		log.Printf("GIN: engine.Run error: %s\n", err.Error())
	}
}
