package controllers

import (
	"fmt"
	"labpro/single-service/initializers"
	"net/http"
	// "strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsClaimsValid(claims jwt.MapClaims) bool {
	return claims["authorized"] != nil && claims["user_id"] != nil && claims["exp"] != nil
}

func ParseJWT(token string) jwt.MapClaims {
	// token = strings.Split(token, " ")[1]

	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.Cfg.SECRET_TOKEN), nil
	})

	return claims
}

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		// CreateResponse(c, http.StatusUnauthorized, "error", "Authorization header is missing.", nil)
		CreateResponse(c, http.StatusUnauthorized, "error", "You are unauthorized.", nil)
		c.Abort()
		return
	}
	
	claims := ParseJWT(token)
	
	if !IsClaimsValid(claims) {
		// CreateResponse(c, http.StatusUnauthorized, "error", "Invalid or expired token.", nil)
		CreateResponse(c, http.StatusUnauthorized, "error", "You are unauthorized.", nil)
		c.Abort()
		return
	}

	if claims["authorized"] == false {
		CreateResponse(c, http.StatusUnauthorized, "error", "You are unauthorized.", nil)
		c.Abort()
		return
	}
	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		CreateResponse(c, http.StatusUnauthorized, "error", "Invalid or expired token.", nil)
		c.Abort()
		return
	}

	c.Set("user_id", claims["user_id"])
	fmt.Println("Authorization access granted.")

	c.Next()
}
