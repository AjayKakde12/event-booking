package routes

import (
	"net/http"
	"strconv"

	"event-booking.com/root/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	var registerEvent models.Registration
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	registerEvent.EventId = eventId
	registerEvent.UserId = userId
	registrationId, err := registerEvent.RegisterForEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Registration Failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registration Successful", "registrationId": registrationId})

}

func cancelRegistration(context *gin.Context) {
	var registrationModel models.Registration
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id"})
		return
	}
	registrationModel.EventId = eventId
	registrationModel.UserId = userId
	registrationId, err := registrationModel.GetRegistration()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Registration not found"})
		return
	}
	err = models.CancelRegistration(registrationId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled"})

}
