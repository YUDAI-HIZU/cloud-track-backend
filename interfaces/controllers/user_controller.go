package controllers

import (
	"strconv"

	"app/usecase/interactor"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type userController struct {
	interactor interactor.UserInteractor
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{
		interactor: interactor.UserInteractor{},
	}
}

func (u *userController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.interactor.GetByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
