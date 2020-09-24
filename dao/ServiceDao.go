package dao

import (
	"CloudRes/model"
	"CloudRes/tool"
)

type ServiceDao struct {
	*tool.Orm
}
//实例化Dao对象
func NewServiceDao() *ServiceDao {
	return &ServiceDao{tool.DbEngine}
}
// 根据serviceID查询service
func (dao *ServiceDao ) QueryServiceByServiceID(serviceID []int64) ([]model.Service) {
	servcies := []model.Service{}
	dao.Engine.Table("service").In("id",serviceID).Find(&servcies)
	return servcies
}
