package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	DB   *gorm.DB
	Gin  *gin.Engine
	Port string
}

func NewRouter(db *gorm.DB) *Router {
	r := &Router{
		DB:   db,
		Gin:  gin.Default(),
		Port: "3000",
	}
	r.Gin.Use(gin.Logger())
	r.Gin.Use(gin.Recovery())
	r.setRouter()
	return r
}

func (r *Router) setRouter() {
	r.Gin.GET("/users/:id", func(c *gin.Context) {
		fmt.Print("====set router func====")
	})
}

func (r *Router) Run() {
	r.Gin.Run(":" + r.Port)
}
