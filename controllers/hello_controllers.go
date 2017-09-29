package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/penguinn/penguin-test/services"
	"github.com/penguinn/penguin/component/log"
	"github.com/penguinn/penguin/component/router"
	"net/http"
)

type HelloController struct {
	GetHelloWorld func(*gin.Context) `path:"/hello"`
}

type HelloWorldRequest struct {
	LastName string `form:"name" binding:"required"`
}

func (HelloController) Name() string {
	return "HelloController"
}

func NewHelloController() HelloController {
	return HelloController{
		GetHelloWorld: GetHelloWorld,
	}
}

func GetHelloWorld(c *gin.Context) {
	var helloWorldRequest HelloWorldRequest
	err := router.BindRequest(c, &helloWorldRequest)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, Response{ErrCode: PARAM_ERROR_CODE, ErrMsg: err.Error()})
	}

	serviceStr, err := services.GetHelloWorld()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, Response{ErrCode: SERVER_ERROR_CODE, ErrMsg: err.Error()})
	}
	str := helloWorldRequest.LastName + "," + serviceStr
	c.JSON(http.StatusOK, Response{Data: str, ErrCode: SUCCESS_CODE})
}
