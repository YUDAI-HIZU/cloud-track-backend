package controllers

import (
	"app/domain"
	"app/interfaces/repository"
	"app/usecase/interactor"
	"net/http"
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
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) Create(c *gin.Context) {
	var user *domain.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := u.Interactor.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (u *UserController) SignIn(c *gin.Context) {
	var input *domain.SignInInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := u.Interactor.SignIn(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
