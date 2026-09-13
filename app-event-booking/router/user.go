package router

import (
	"net/http"

	"example.com/app-event-booking/models"
	"example.com/app-event-booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user!",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func sigin(c *gin.Context) {
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Missing field.",
		})
		return
	}
	userDB, result := models.ValidateCredentials(&user)
	if !result {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isLogin": false,
			"message": "Invalid credentials.",
		})
		return
	}

	token, err := utils.GenerateToken(userDB.Email, userDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"isLogin": result,
		"token":   token,
	})
}
