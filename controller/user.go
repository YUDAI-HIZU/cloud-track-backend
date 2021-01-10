package controller

import (
	"app/model"
	"app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "error": err.Error()})
		return
	}
	UserService := service.UserService{}
	err := UserService.Create(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
