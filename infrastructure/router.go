package infrastructure

import (
	"app/infrastructure/middleware"
	"app/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	DB  *gorm.DB
	Gin *gin.Engine
}

func NewRouter(db *gorm.DB) *Router {
	r := &Router{
		DB:  db,
		Gin: gin.Default(),
	}
	r.Gin.Use(gin.Logger())
	r.Gin.Use(gin.Recovery())
	r.setRouter()
	return r
}

func (r *Router) setRouter() {
	userController := controllers.NewUserController(r.DB)

	r.Gin.POST("/sign-up", func(c *gin.Context) {
		userController.Create(c)
	})
	r.Gin.POST("sign-in", func(c *gin.Context) {
		userController.SignIn(c)
	})

	users := r.Gin.Group("/users")
	{
		users.GET("/:id", func(c *gin.Context) {
			userController.GetByID(c)
		})
	}

	account := r.Gin.Group("/account")
	{
		account.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			userController.CurrentUser(c)
		})
		account.GET("/profiles", middleware.AuthMiddleware(), func(c *gin.Context) {

		})
		account.POST("/profiles", middleware.AuthMiddleware(), func(c *gin.Context) {

		})
		account.PUT("/profiles", middleware.AuthMiddleware(), func(c *gin.Context) {

		})
	}
}

func (r *Router) Run() {
	r.Gin.Run()
}
