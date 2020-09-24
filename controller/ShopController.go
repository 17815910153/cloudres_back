package controller

import (
	"CloudRes/service"
	"CloudRes/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct {

}

func (sc *ShopController ) Router(e *gin.Engine) {
	e.GET("/api/shops",sc.GetNearShop)
	e.GET("/api/search_shops",sc.SearchShops)

}
// 根据关键字查询附近商家。
func (sc *ShopController) SearchShops(c *gin.Context)  {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")
	keywords := c.Query("keywords")

	if keywords == "" {
		tool.Failed(c,"请重新输入商铺名称")
		return

	}
	// 设置默认位置
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "113.29" //珠海
		latitude = "22.20"
	}
	shopService := service.ShopService{}
	shopsVO, err := shopService.GetNearShops(longitude, latitude, keywords)
	if err != nil || len(shopsVO) == 0 {
		tool.Failed(c, "暂未获取到商户信息")
		return
	}
	tool.Success(c, shopsVO)

}
// 查询附近商家
func (sc *ShopController) GetNearShop(c *gin.Context)  {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "113.29" //珠海
		latitude = "22.20"
	}
	//fmt.Println("longitude,latitude",longitude,latitude)
	shopService := service.ShopService{}
	shopsVO, err := shopService.GetNearShops(longitude, latitude,"")
	if err != nil || len(shopsVO) == 0 {
		tool.Failed(c, "暂未获取到商户信息")
		return
	}
	tool.Success(c, shopsVO)
}

