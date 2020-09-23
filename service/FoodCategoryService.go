package service

import (
	"CloudRes/dao"
	"CloudRes/model"
)

type FoodCategoryService struct {

}
//定义全局变量
//var categoryDao = dao.FoodCategoryDao{tool.DbEngine}

//var foodCategoryDao = dao.NewFoodCategoryDao()
func (service *FoodCategoryService) GetAllCategory() ([]model.FoodCategory, error)  {
	foodCategoryDao := dao.NewFoodCategoryDao()
	categories, err := foodCategoryDao.QueryAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}