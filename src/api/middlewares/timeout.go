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

func Timeout(t time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(t*time.Millisecond),
		timeout.WithHandler(emptySuccessResponse),
	)

}
