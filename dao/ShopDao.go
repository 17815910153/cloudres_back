package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
	"fmt"
)

const MAX_NEAR_DISTANCE = 5

type ShopDao struct {
	*tool.Orm
}

// 初始化ShopDao
func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}


// 查询附近的商家信息
func (sd *ShopDao) QueryNearShop(longi, lati float64) ([]model.Shop, error) {
	Max_longi := longi + MAX_NEAR_DISTANCE
	Min_longi := longi - MAX_NEAR_DISTANCE
	Max_lati := lati + MAX_NEAR_DISTANCE
	Min_lati := lati - MAX_NEAR_DISTANCE
	var shops []model.Shop


	fmt.Println(Max_longi,Min_longi,Max_lati,Min_lati)
	err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?", Min_longi, Max_longi, Min_lati, Max_lati).Find(&shops)

	if err !=nil {
		return nil,err
	}
	return shops, nil
}