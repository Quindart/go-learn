package router

import (
	"database/sql"
	"errors"
	"fmt"
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
	userID := c.GetInt64("userId")
	var event models.Event
	err := c.ShouldBindBodyWithJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No event found"})
		return
	}
	event.UserID = int(userID)
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
	userID := c.GetInt64("userId")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No id found",
		})
		return
	}
	event, err := models.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find this event, try again",
		})
		return
	}
	if int64(event.UserID) != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You can only update your own events."})
		return
	}
	var updateEvent models.Event
	fmt.Println("updateEvent", updateEvent)
	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse this event",
		})
		return
	}
	updateEvent.ID = id
	err = models.UpdateByID(id, updateEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch event",
		})
	}
}

func deleteEvent(c *gin.Context) {
	userID := c.GetInt64("userId")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "id not found",
		})
		return
	}
	event, err := models.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event."})
		}
		return
	}
	if int64(event.UserID) != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You can only delete your own events."})
		return
	}
	err = models.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to deleted by id",
		})
		return
	}

}
