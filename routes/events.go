package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaustubsingh09/go-auth-gin/models"
)

func postEvents(context *gin.Context) {

	userId := context.GetInt64("userId")

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save the event, try again later"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data try again later"})
		return
	}
	context.JSON(200, events)
}

func getUniqueEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event try again later"})
		return
	}
	event, err := models.GetUniqueEvent(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event does not exists"})
		return
	}
	context.JSON(200, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse int"})
		return
	}

	userId := context.GetInt("userId")
	event, err := models.GetUniqueEvent(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event does not exists"})
		return
	}

	if int(event.UserId) != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return

	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "request failed please try again after some time"})
		return
	}

	updatedEvent.ID = eventId

	_, err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func deleteEvent(context *gin.Context) {
	userId := context.GetInt("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse event id"})
		return
	}

	event, err := models.GetUniqueEvent(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event does not exists"})
		return
	}

	if event.UserId != int64(userId) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	//logic to delete event
	var deleteEvent models.Event

	deleteEvent.ID = eventId

	err = deleteEvent.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})

}
