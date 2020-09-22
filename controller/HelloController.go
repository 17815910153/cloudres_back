package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {

}

//简单的测试是否可以通gin
func (hello *HelloController) Router(engine *gin.Engine)  {
	engine.GET("/hello", hello.Hello)

}
//解析 /hello
func (hello *HelloController) Hello(context *gin.Context)  {
	context.JSON(http.StatusOK,"hello world!")

}
