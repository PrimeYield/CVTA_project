package middleware

import (
	jwt "exercise/pkg/JWT"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return authMiddleware
}

func authMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("jwt_token")
	if err != nil {
		if err == http.ErrNoCookie {
				log.Println("No 'myAuthToken' cookie found, unauthorized.")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized: No authentication token found.",
				})
				return
			}
			log.Printf("Error getting cookie: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request.",
			})
			return
	}

	token, err := jwt.ValidateToken(tokenStr)
	log.Printf("Middlerware:%v", token)
	if err != nil || token == nil {
		fmt.Printf("tokenStr is a wrong request %v", err)
		return
	}
	// username, err := token.Get("username")
	createdBy, ok := token.Get("username")
	if !ok {
		fmt.Println("Middlleware: <key: username> is not exists")
		return
	}
	c.Set("username", createdBy)
	c.Next()
}