package api

import (
	"github.com/gin-gonic/gin"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/change password", changePassword)
	r.POST("/forget password", forgetPassword)
	r.POST("/answer", answer)
	r.POST("/check code", checkCode)
	r.POST("/add comment", addComment)
	r.POST("/scan comments", scanComments)
	r.POST("delete comment", deleteComment)
	r.POST("/clear comments", clearComments)
	r.POST("/quit", quit)
	r.POST("/unsubscribe", unsubscribe)
	r.POST("/clear all", clearAll)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}
	r.Run(":11451")
}
