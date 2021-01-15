package router

import (
	"app/infrastructure/database"
	"app/interfaces/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	userController := controllers.NewUserController(database.Conn())
	Router.GET("/users/:id", func(c *gin.Context) { userController.GetByID(c) })
	if err := http.ListenAndServe(":"+"3000", Router); err != nil {
		log.Print(err.Error())
	}
}
