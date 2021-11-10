package main

import (
	"net/http"

	"covid_admission_api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	covidAdmissionRoute := server.Group("/covid-admission")
	covidAdmissionRoute.Use(middlewares.UploadXrayImageTimeout())
	{
		covidAdmissionRoute.GET("/image-classification/upload-xray-image", uploadXrayImageHandler)
	}

	server.Run()
}

func uploadXrayImageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Admitted",
	})
}
