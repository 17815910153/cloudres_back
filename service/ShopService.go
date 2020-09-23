package service

import (
	"CloudRes/dao"
	"CloudRes/model"
	"strconv"
)

type ShopService struct {

}
/**
获取附近的商家
 */
func (shopService *ShopService ) GetNearShop(longi, lati string ) ([]model.Shop, error) {
	long, err := strconv.ParseFloat(longi, 10)
	if err != nil {
		return nil, err
	}
	latitu, err := strconv.ParseFloat(lati, 10)
	if err != nil {
		return nil, err
	}
	shopDao := dao.NewShopDao()
	shops, err := shopDao.QueryNearShop(long, latitu)
	if err != nil{
		return nil, err
	}
	return shops, nil

}
