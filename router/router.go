package router

import (
	"app/controller"
	"app/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "3000"
	}
	r.GET("/token", middleware.AuthMiddleware())
	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", controller.SignIn)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Print(err.Error())
	}
}
