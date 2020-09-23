package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}
//实例化Dao对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}
// 查询所有的食物分类列表
func (dao *FoodCategoryDao ) QueryAllCategory() ([]model.FoodCategory, error) {
	var categories[] model.FoodCategory
	err := dao.Find(&categories)
	if err !=nil {
		return nil, err
	}
	return categories, nil
}
