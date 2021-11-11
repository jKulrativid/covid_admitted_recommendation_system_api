package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	c.JSON(http.StatusOK, gin.H{
		"timeoutMessage": "In Time",
	})
}

func UploadXrayImageTimeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(200*time.Millisecond),
		timeout.WithHandler(emptySuccessResponse),
	)

}
