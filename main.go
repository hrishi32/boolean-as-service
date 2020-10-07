package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hrishi32/boolean-as-service/models"
	"github.com/hrishi32/boolean-as-service/routes"
)

func main() {
	server := gin.Default()
	defaultRepo := models.RepoImplement{}
	models.SetRepo(&defaultRepo)
	models.Migrate()
	routes.Init(server)

	server.Run(":8000")
}
