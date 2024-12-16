package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorDetection(client *gin.Context, err error) {
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
