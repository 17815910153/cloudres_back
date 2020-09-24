package service

import (
	"CloudRes/dao"
	"CloudRes/model"
)

type GoodsService struct {
	
}

func (service *GoodsService) GetGoodsByShopId(id string) ([]model.Goods,error)  {

	goodsDao := dao.NewGoodsDao()
	goods, err := goodsDao.QueryGoodsByShopId(id)
	if err != nil {
		return nil,err
	}
	return goods, nil
}
