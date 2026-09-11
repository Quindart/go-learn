package main

import (
	"example.com/app-event-booking/db"
	"example.com/app-event-booking/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()
	router.RegisterRouter(server)
	server.Run(":3001")

}
