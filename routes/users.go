package routes

import (
	"net/http"

	"github.com/anupamraj16/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse incoming data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created."})
}