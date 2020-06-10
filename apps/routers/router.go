package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanopoker/minisns/apps/controllers"
	wrapper "github.com/nanopoker/minisns/libs/controller_wrapper"
)

func InitializeRouter(router *gin.Engine) {
	authorized := router.Group("/api")
	authorized.Use(wrapper.AuthenticationRequired)
	{
		authorized.POST("/logout", controller.LogoutHandler)
		authorized.POST("/edit_user", controller.EditUserHandler)
		authorized.POST("/follow", controller.FollowHandler)
		authorized.GET("/follow_list", controller.FollowlistHandler)
	}

	router.POST("/login", controller.LoginHandler)
	router.POST("/register", controller.RegisterHandler)
}
