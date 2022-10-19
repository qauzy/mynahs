package service

import (
	"github.com/qauzy/mynahs/internal/dto"
	"github.com/qauzy/mynahs/internal/model"
)

type OrdersBo struct {
	Data []*OrderData
}
type OrderData struct {
	ID    uint   `json:"id"`
	Money string `json:"money"`
	Count uint   `json:"count"`
}

type OrdersService struct {
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}

func (o *OrdersService) Create(req *dto.OrdersDto) (err error) {
	obj := model.NewOrders()
	err = obj.Create(req)
	return
}

func (o *OrdersService) List(req *dto.OrdersDto) (result *OrdersBo, err error) {
	obj := model.NewOrders()
	var data []*model.Orders
	data, err = obj.List(req)
	if err != nil {
		return
	}

	//封装返回数据
	result = new(OrdersBo)
	if data == nil || len(data) == 0 {
		result.Data = make([]*OrderData, 0)
		return
	}
	for _, v := range data {
		item := &OrderData{
			ID:    v.ID,
			Money: v.Money,
			Count: v.Count,
		}
		result.Data = append(result.Data, item)
	}
	return
}
