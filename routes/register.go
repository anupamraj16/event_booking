package routes

import (
	"net/http"
	"strconv"

	"github.com/anupamraj16/models"
	"github.com/gin-gonic/gin"
)

func getRegistrations(context *gin.Context) {
	registrations, err := models.GetAllRegistrations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error."})
		return
	}
	context.JSON(http.StatusOK, registrations)
}

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered."})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	e, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if e.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to cancel registration."})
		return
	}

	err = e.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled."})
}
