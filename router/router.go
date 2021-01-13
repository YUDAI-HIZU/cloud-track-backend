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
	authMiddleware, _ := middleware.AuthMiddleware()
	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", authMiddleware.LoginHandler, controller.SignIn)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	g := r.Group("/", authMiddleware.MiddlewareFunc())
	{
		g.GET("/auth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": "auth ok"})
		})
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Print(err)
	}
}
