package main

import (
	"fmt"
	"net/http"

	"event-booking.com/root/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	error := context.ShouldBindJSON(&event) //This will save request body in event pointer
	if error != nil {
		fmt.Println(error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}
	event.ID = 1
	event.UserId = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
