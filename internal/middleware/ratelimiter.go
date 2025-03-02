package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/utils"
	"sync"
	"time"
)

var (
	requestCounts = make(map[string]int)
	mutex         sync.Mutex
)

func RateLimiter(limit int, period time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mutex.Lock()
		defer mutex.Unlock()

		if _, exists := requestCounts[ip]; !exists {
			requestCounts[ip] = 0
			go func() {
				time.Sleep(period)
				mutex.Lock()
				delete(requestCounts, ip)
				mutex.Unlock()
			}()
		}

		requestCounts[ip]++

		if requestCounts[ip] > limit {
			c.JSON(http.StatusTooManyRequests, utils.ResponseData(false, "Too many requests. Try again later.", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
