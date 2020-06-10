package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nanopoker/minisns/apps/rpcclient"
	"github.com/nanopoker/minisns/apps/types"
	"github.com/nanopoker/minisns/config"
	wrapper "github.com/nanopoker/minisns/libs/controller_wrapper"
)

func LoginHandler(c *gin.Context) {
	var json types.LoginParam
	wrapper.ParamsCheck(c, &json)
	rsp, identity := rpcclient.Login(json.Username, json.Password)
	if identity != "" {
		c.SetCookie("identity", identity, config.COOKIE_DURATION, "/", "localhost", false, true)
	}
	wrapper.HttpServerResponse(c, rsp, nil)
}

func LogoutHandler(c *gin.Context) {
	identity, _ := c.Cookie("identity")
	rsp := rpcclient.Logout(identity)
	wrapper.HttpServerResponse(c, rsp, nil)
}

func RegisterHandler(c *gin.Context) {
	var json types.RegisterParam
	wrapper.ParamsCheck(c, &json)
	rsp := rpcclient.Register(json.Username, json.Password, json.Email, json.Nickname)
	wrapper.HttpServerResponse(c, rsp, nil)
}

func EditUserHandler(c *gin.Context) {
	var json types.EditParam
	wrapper.ParamsCheck(c, &json)
	identity, _ := c.Cookie("identity")
	rsp, userid, nickname, email, username := rpcclient.EditUser(json.Nickname, json.Password, json.Email, identity)
	wrapper.HttpServerResponse(c, rsp, map[string]interface{}{"username": username, "userid": userid, "email": email, "nickname": nickname})
}
