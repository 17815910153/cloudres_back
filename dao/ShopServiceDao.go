package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
)

type ShopServiceDao struct {
	*tool.Orm
}
//实例化Dao对象
func NewShopServiceDao() *ShopServiceDao {
	return &ShopServiceDao{tool.DbEngine}
}
// 根据shopID,查询ShopService
func (dao *ShopServiceDao ) QueryAllByShopIDs(shopId []int64) ([]model.ShopService) {

	shopServices := []model.ShopService{}
	dao.Engine.Table("shop_service").In("shop_id",shopId).Find(&shopServices)

	return shopServices

}
