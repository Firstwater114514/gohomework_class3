package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}
func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func Question(c *gin.Context, message, tip string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip":     tip,
	})
}
func AnswerRight(c *gin.Context, message, tip1, tip2 string, checkCode int) {
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"message":   message,
		"tip1":      tip1,
		"tip2":      tip2,
		"checkCode": checkCode,
	})
}
func Comments(c *gin.Context, comments, _1, _2, _3, _4, _5, _6 string) {
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"comments": comments,
		"1":        _1,
		"2":        _2,
		"3":        _3,
		"4":        _4,
		"5":        _5,
		"6":        _6,
	})
}
func LoginSuccess(c *gin.Context, message, tip1, tip2, tip3, tip4, tip5 string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip1":    tip1,
		"tip2":    tip2,
		"tip3":    tip3,
		"tip4":    tip4,
		"tip5":    tip5,
	})
}
