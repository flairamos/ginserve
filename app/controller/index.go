package controller

import (
	"github.com/gin-gonic/gin"
)

func Index() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello World"})
	}
}
