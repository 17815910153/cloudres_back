package controller

import (
	"CloudRes/service"
	"CloudRes/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
	
}

var cateryService = service.FoodCategoryService{}

//进行路由转发，定义router
func (cc *FoodCategoryController) Router(e *gin.Engine) {
	e.GET("/api/categories", cc.getAllFoodCategory)
	
}


func (cc *FoodCategoryController) getAllFoodCategory(c *gin.Context)  {
	categories, err := cateryService.GetAllCategory()
	if err != nil {
		tool.Failed(c,"获取食物类别失败")
	}
	tool.Success(c,categories)
}
