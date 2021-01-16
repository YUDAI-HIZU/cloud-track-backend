package controllers

import (
	"app/domain"
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

func (u *UserController) Create(c *gin.Context) {
	var user *domain.User
	c.Bind(&user)
	err := u.Interactor.Create(user)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(201, "")
}
