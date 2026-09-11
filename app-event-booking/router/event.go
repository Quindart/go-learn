package router

import (
	"net/http"
	"strconv"

	"example.com/app-event-booking/models"
	"github.com/gin-gonic/gin"
)

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch!"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindBodyWithJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No event found"})
		return
	}
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event!",
		})
		return
	}
	c.JSON(http.StatusCreated, event)

}

func getEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "id not found "})
		return
	}
	// handler get by id
	event, err := models.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func putEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No id found",
		})
		return
	}
	_, err = models.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find this event, try again",
		})
		return
	}
	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse this event",
		})
		return
	}
	updateEvent.ID = id
	
}
