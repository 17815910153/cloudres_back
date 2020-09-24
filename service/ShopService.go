package service

import (
	"CloudRes/VO"
	"CloudRes/dao"
	"CloudRes/model"
	"strconv"
)

type ShopService struct {

}
/**
获取附近的商家
 */
//func (shopService *ShopService ) GetNearShop(longi, lati string ) ([]model.Shop, error) {
//	long, err := strconv.ParseFloat(longi, 10)
//	if err != nil {
//		return nil, err
//	}
//	latitu, err := strconv.ParseFloat(lati, 10)
//	if err != nil {
//		return nil, err
//	}
//	shopDao := dao.NewShopDao()
//	shops, err := shopDao.QueryNearShop(long, latitu,"")
//	if err != nil{
//		return nil, err
//	}
//	return shops, nil
//
//}

/**
根据输入名称获取附近商家
 */
func (shopService *ShopService) GetNearShops(longitude string, latitude string, keywords string) ([]VO.ShopInfoVO, error) {
	long, err := strconv.ParseFloat(longitude, 10)
	if err != nil {
		return nil, err
	}
	latitu, err := strconv.ParseFloat(latitude, 10)
	if err != nil {
		return nil, err
	}
	shopDao := dao.NewShopDao()

	var shopVOS  []VO.ShopInfoVO


	// 1.先从shop数据库找出符合条件的商家信息。
	shops, err := shopDao.QueryNearShop(long, latitu,keywords)

	if err != nil {
		return nil, err
	}

	// 2.从shop_service数据库 找出相对应的 shop_id
	// 需要从shops中评出一个 shop_id 切片
	var shopId []int64
	for _, shop := range shops{
		shopId = append(shopId, shop.Id)
	}

	// 查出对应的Shop_Service
	shopserviceDao := dao.NewShopServiceDao()
	shopServices := shopserviceDao.QueryAllByShopIDs(shopId)

	// 3. 从service数据库中，找对对应的信息
	// 需要从shopService中评出一个service_id 切片
	var serviceID []int64
	for _, shopService := range shopServices{
		serviceID = append(serviceID,shopService.ServiceId )
	}
	// 查出对应的service
	serviceDao := dao.NewServiceDao()
	servcies := serviceDao.QueryServiceByServiceID(serviceID)

	// 4.将获取到的service和shop信息拼接成一个VO,返回
	for _,shop := range shops {
		var shopVO VO.ShopInfoVO
		shopTOshopVO(&shop, &shopVO)
		// 判断是否有对应的服务
		for _, shopService := range shopServices{
			for _, service := range servcies{
				if  shopVO.Id == shopService.ShopId && service.Id == shopService.ServiceId {
					shopVO.Supports = append(shopVO.Supports, service)

				}
			}
		}
		shopVOS = append(shopVOS, shopVO)
	}

	return shopVOS,nil

}

// 实现shop->shopVO结构体的复制
func shopTOshopVO(shop *model.Shop, vo *VO.ShopInfoVO)  {
	vo.Id =	shop.Id
	vo.Name	= shop.Name
	vo.PromotionInfo = shop.PromotionInfo
	vo.Address = shop.Address
	vo.Phone = shop.Phone
	vo.Status =	shop.Status
	vo.Longitude = shop.Longitude
	vo.Latitude	= shop.Latitude
	vo.ImagePath = shop.ImagePath
	vo.IsNew = shop.IsNew
	vo.IsPremium = shop.IsPremium
	vo.Rating = shop.Rating
	vo.RatingCount = shop.RatingCount
	vo.RecentOrderNum = shop.RecentOrderNum
	vo.MinimumOrderAmount = shop.MinimumOrderAmount
	vo.DeliveryFee = shop.DeliveryFee
	vo.OpeningHours	= shop.OpeningHours
}
