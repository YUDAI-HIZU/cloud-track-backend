package controllers

import (
	"app/usecase/interactor"
	"app/usecase/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	interactor interactor.UserInteractor
}

func NewUserController(db database.DB) *userController {
	return &userController{
		interactor: interactor.UserInteractor{
			UserRepository: repository.UserRepository{
				db: db,
			},
		},
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
