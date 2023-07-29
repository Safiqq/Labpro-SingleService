package controllers

import "github.com/gin-gonic/gin"

func CreateResponse(c *gin.Context, statusCode int, status string, message string, data interface{}) {
	c.IndentedJSON(statusCode, gin.H{
		"status": status,
		"message": message,
		"data": data,
	})
}