package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
)

const MAX_NEAR_DISTANCE = 5

type ShopDao struct {
	*tool.Orm
}

// 初始化ShopDao
func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}


// 查询关键字附近的商家信息,
func (sd *ShopDao) QueryNearShop(longi, lati float64, keywords string) ([]model.Shop, error) {
	Max_longi := longi + MAX_NEAR_DISTANCE
	Min_longi := longi - MAX_NEAR_DISTANCE
	Max_lati := lati + MAX_NEAR_DISTANCE
	Min_lati := lati - MAX_NEAR_DISTANCE

	var shops []model.Shop
	// 1.先从shop数据库找出符合条件的商家信息。
	if keywords == "" {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? ",
				Min_longi, Max_longi, Min_lati, Max_lati).Find(&shops)
		if err !=nil {
			return nil,err
		}
	} else {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and shop.name  like  CONCAT('%',?,'%')",
				Min_longi, Max_longi, Min_lati, Max_lati, keywords).Find(&shops)
		if err !=nil {
			return nil,err
		}
	}

	return shops, nil

}

