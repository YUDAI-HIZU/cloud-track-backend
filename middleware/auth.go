package middleware

import (
	"app/config"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
		fmt.Print(token, err)
		if err != nil {
			log.Print("failed to parse jwt token", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "認証に失敗しました", "error": err.Error()})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Print("failed to get claims", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "認証に失敗しました", "error": err.Error()})
			return
		}
		c.Set("userID", claims["userID"].(float64))
		c.Next()
	}
}
