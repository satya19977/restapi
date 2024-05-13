package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satya19977/restapi/db"
	"github.com/satya19977/restapi/models"
)

func GetEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"),10,64)   
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message": "Couldn't parse event id "})
		return
	}
	event,err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't fetch event"})
		return
	}
	context.JSON(http.StatusOK,event)
	
}

func GetEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func healthCheckHandler(context *gin.Context){
	//Check if the service is up
	status := http.StatusOK
	message := "OK"

	//Check if DB is reachable
	err := db.DB.Ping()
	if err != nil{
		status = http.StatusServiceUnavailable
		message = "Database is not reachable"
	}
	context.JSON(status,gin.H{
		"status": status,
		"message": message,
	})
}

func CreateEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Couldn't parse"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't create events"})
	    return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"event created","event": event})
}

func UpdateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"),10,64)   
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message": "Couldn't parse event id "})
		return
	}
	_,err = models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't create events"})
	    return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't update event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"),10,64)   
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message": "Couldn't parse event id "})
		return
	}
	event,err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "Couldn't fetch event. Error from GetEventByID"})
		return
	}
	event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message": "Couldn't delete the event. Error from event.Delete "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message":"Event Successfuly deleted"})

}
