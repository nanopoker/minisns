package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nanopoker/minisns/apps/rpcclient"
	"github.com/nanopoker/minisns/apps/types"
	wrapper "github.com/nanopoker/minisns/libs/controller_wrapper"
)

func FollowHandler(c *gin.Context) {
	var json types.FollowParam
	wrapper.ParamsCheck(c, &json)
	identity, _ := c.Cookie("identity")
	rsp := rpcclient.Follow(identity, json.Followee)
	wrapper.HttpServerResponse(c, rsp, nil)
}

func FollowlistHandler(c *gin.Context) {
	var json types.FollowlistParam
	wrapper.ParamsCheck(c, &json)
	identity, _ := c.Cookie("identity")
	pageno := json.Pageno
	count := json.Count
	rsp, follows := rpcclient.Followlist(pageno, count, identity)
	wrapper.HttpServerResponse(c, rsp, map[string]interface{}{"list": follows, "total": len(follows)})
}
