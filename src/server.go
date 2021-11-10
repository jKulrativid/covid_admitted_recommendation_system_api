package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/image-classification/upload-xray-image", uploadXrayImageHandler)

	server.Run()
}

func uploadXrayImageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Admitted",
	})
}
