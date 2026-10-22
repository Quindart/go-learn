package main

import (
	"fmt"
	"net/http"

	"example/fullstack-be/db"
	"example/fullstack-be/model"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("hello world!")
	db.InitDB()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	router.GET("/todos", func(c *gin.Context) {
		todos, err := model.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "No todos found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})
	
	router.POST("/todos", func(c *gin.Context) {
		var todo model.Todo
		err := c.ShouldBindBodyWithJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No todo found"})
			return
		}
		err = todo.Save()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create event!",
			})
			return
		}
		c.JSON(http.StatusCreated, todo)
	})
	router.Run(":8600")
}
