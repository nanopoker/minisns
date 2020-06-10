package controller_wrapper

import (
	"github.com/gin-gonic/gin"
	code "github.com/nanopoker/minisns/apps/retcode"
	"github.com/nanopoker/minisns/apps/rpcclient"
	"github.com/nanopoker/minisns/libs/logger"
	"net/http"
)

func HttpServerResponse(c *gin.Context, rsp code.Code, data interface{}) {
	var httpcode int
	if rsp == code.CodeSucc {
		httpcode = http.StatusOK
	} else if rsp == code.CodeInternalErr {
		httpcode = http.StatusInternalServerError
	} else {
		httpcode = http.StatusBadRequest
	}
	c.JSON(httpcode, gin.H{
		"retcode": int(rsp.CodeNum),
		"message": rsp.CodeMsg,
		"data":    data,
	})
}

func paramsCheck(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	return nil
}

func ParamsCheck(c *gin.Context, params interface{}) {
	err := paramsCheck(c, params)
	if err != nil {
		logger.Error("params check error, ", err.Error())
		errormsg := "params error, please check your params"
		HttpServerResponse(c, code.Code{CodeNum: uint32(http.StatusBadRequest), CodeMsg: errormsg}, nil)
	}
}

func AuthenticationRequired(c *gin.Context) {
	identity, err := c.Cookie("identity")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"retcode": http.StatusUnauthorized,
			"message": "you need to login",
			"data":    nil,
		})
		return
	} else {
		rsp := rpcclient.AuthIdentity(identity)
		if rsp != code.CodeSucc {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"retcode": http.StatusUnauthorized,
				"message": "you need to login",
				"data":    nil,
			})
		}
	}
}
