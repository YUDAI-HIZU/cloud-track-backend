package router

import (
	"app/controller"
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

	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", controller.SignIn)
	g := r.Group("/")
	{
		g.GET("/auth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": "auth ok"})
		})
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Print(err)
	}
}
