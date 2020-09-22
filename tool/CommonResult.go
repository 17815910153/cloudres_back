package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int  = 0  //操作成功
	FAILED int  = 1  //操作失败

)

//普通成功返回
func Success(c *gin.Context, v interface{})  {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg": "成功",
		"data": v,
	})
}

// 普通失败返回
func Failed(c *gin.Context, v interface{})  {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"data": v,
	})

}
