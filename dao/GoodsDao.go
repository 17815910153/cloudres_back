package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
)

type GoodsDao struct {
	*tool.Orm
}

func NewGoodsDao()  *GoodsDao {
	return &GoodsDao{Orm:tool.DbEngine}
}
func (goodDao *GoodsDao) QueryGoodsByShopId(shopId string) ([]model.Goods, error) {

	var goods []model.Goods
	err := goodDao.Where("shop_id = ?", shopId).Find(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
