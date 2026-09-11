package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)
	server.PUT("/events", putEvent)
}
