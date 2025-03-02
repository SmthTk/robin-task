package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseStringToUint(st string) uint {
	uintP, _ := strconv.ParseUint(st, 10, 64)
	return uint(uintP)
}

func GetUserIDFromContext(c *gin.Context) uint {
	userIDIntf, _ := c.Get("userID")
	return userIDIntf.(uint)
}

func ResponseData(status bool, message string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
