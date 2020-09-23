package controller

import (
	"CloudRes/service"
	"CloudRes/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopController struct {

}

func (sc *ShopController ) Router(e *gin.Engine) {
	e.GET("/api/shops",sc.GetNearShop)

}
func (sc *ShopController) GetNearShop(c *gin.Context)  {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "113.29" //珠海
		latitude = "22.20"
	}
	fmt.Println("longitude,latitude",longitude,latitude)
	shopService := service.ShopService{}
	shops, err := shopService.GetNearShop(longitude, latitude)
	if err != nil || len(shops) == 0 {
		tool.Failed(c, "暂未获取到商户信息")
		return
	}
	tool.Success(c, shops)



}

