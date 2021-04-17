package router

import (
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4/middleware"
	"DGM-B/handler"
)

func GetUserAPI(e *echo.Echo) {
	g := e.Group("/user")
	g.POST("/login", handler.Login)
	g.POST("/register",handler.Register)
}

func GetUserInfoAPI(e *echo.Echo) {
	g := e.Group("/userinfo",middleware.BasicAuth(handler.Auth))
	g.GET("/getUserInfo",handler.GetUserInfo)
	g.POST("/changeUserInfo",handler.ChangeUserInfo)
	g.POST("/changePWD",handler.ChangePwd)
	g.POST("/changeAvatar",handler.ChangeAvatar)
}
