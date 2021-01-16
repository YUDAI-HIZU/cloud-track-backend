package controllers

import (
	"app/interfaces/repository"
	"app/usecase/interactor"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	Interactor interactor.UserInteractor
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}
}

func (u *UserController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.Interactor.GetByID(id)
	fmt.Println(user, err)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, user)
}
