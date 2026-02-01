package routes

import (
	"eventsapi/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find event"})
		return
	}

	err = event.Register()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = id
	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	context.JSON(http.StatusInternalServerError, gin.H{"message": "cancelled registration"})
}
