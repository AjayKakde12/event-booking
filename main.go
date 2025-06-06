package main

import (
	"event-booking.com/root/db"
	"event-booking.com/root/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
