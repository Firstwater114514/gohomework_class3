package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Change struct {
	NewPassword string `form:"new password" json:"new password" binding:"required"`
}
type Forget struct {
	Username string `form:"username" json:"username" binding:"required"`
}
type Question struct {
	Answer string `form:"answer" json:"answer" binding:"required"`
}
type Code struct {
	Code        string `form:"code" json:"code" binding:"required"`
	Username    string `form:"username" json:"username" binding:"required"`
	NewPassword string `form:"new password" json:"new password" binding:"required"`
}
type AddComment struct {
	Comment string `form:"comment" json:"comment" binding:"required"`
}
type DeleteComment struct {
	Num string `form:"num" json:"num" binding:"required"`
}
