package server

import (
	"fmt"
	"log"
	"time"

	"github.com/NurbekDos/funk/internal/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port = 8081 // TODO: alyp tastap, env ga sal wws

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

	err := engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("GIN: engine.Run error: %s\n", err.Error())
	}
}
