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
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "error": err.Error()})
		return
	}
	UserService := service.UserService{}
	err := UserService.Create(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}

func SignIn(c *gin.Context) {
	info := model.UserInfo{}
	if err := c.ShouldBindJSON(&info); err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "error": err.Error()})
		return
	}
	UserService := service.UserService{}
	user, err := UserService.GetByEmail(info.Email)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "ユーザーが存在しません", "error": err.Error()})
	}
	err = user.PasswordVerify(user.Password, info.Password)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "パスワードが違います", "error": err.Error()})
	}
	token, _ := user.GenerateToken()
	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
		"data":   user,
		"token":  token,
	})
}
