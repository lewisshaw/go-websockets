package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ping": true,
	})
}
