package router

import (
	"example.com/app-event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)

	// Các route trong nhóm này đều yêu cầu token hợp lệ.
	events := server.Group("/events")
	events.Use(middlewares.Authenticate)
	{
		events.POST("", createEvent)
		events.PUT("/:id", putEvent)
		events.DELETE("/:id", deleteEvent)
	}
	server.POST("/signup", signup)
	server.POST("/signin", sigin)
}
