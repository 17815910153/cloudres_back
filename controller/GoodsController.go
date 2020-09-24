package controller

import (
	"CloudRes/service"
	"CloudRes/tool"
	"github.com/gin-gonic/gin"
)

type GoodsController struct {

}

func (goods *GoodsController) Router(e *gin.Engine)  {
	e.GET("/api/goods/:id", goods.GetShopGood)

}

/**
根据商家id,查询在售的食物
 */
func (goods *GoodsController) GetShopGood(c *gin.Context)  {
	shopId := c.Param("id")

	if shopId == "" {
		tool.Failed(c, "该商铺还没开张，请稍等一会！")
		return
	}
	goodsService := service.GoodsService{}
	list, err := goodsService.GetGoodsByShopId(shopId)
	//fmt.Println(shopId)
	if err !=nil {
		tool.Failed(c, "网络异常，请稍后再试！")
		return
	}
	tool.Success(c,list)

}
