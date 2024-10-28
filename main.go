package main

import (
	"github.com/balebbae/events-go-api/db"
	"github.com/balebbae/events-go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

