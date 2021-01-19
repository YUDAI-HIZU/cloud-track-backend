package infrastructure

import (
	"app/infrastructure/middleware"
	"app/interfaces/controllers"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	DB   *gorm.DB
	Gin  *gin.Engine
	App  *firebase.App
	Port string
}

func NewRouter(db *gorm.DB, app *firebase.App) *Router {
	r := &Router{
		DB:   db,
		Gin:  gin.Default(),
		App:  app,
		Port: "3000",
	}
	r.Gin.Use(gin.Logger())
	r.Gin.Use(gin.Recovery())
	r.setRouter()
	return r
}

func (r *Router) setRouter() {
	userController := controllers.NewUserController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) {
		userController.GetByID(c)
	})
	r.Gin.POST("/sign-up", func(c *gin.Context) {
		userController.Create(c)
	})
	r.Gin.POST("sign-in", func(c *gin.Context) {
		userController.SignIn(c)
	})
	r.Gin.GET("/auth", middleware.AuthMiddleware(r.App))
}

func (r *Router) Run() {
	r.Gin.Run(":" + r.Port)
}
