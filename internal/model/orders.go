package model

import "github.com/qauzy/mynahs/internal/dto"

type Orders struct {
	BaseModel
	Money string `gorm:"column:money" json:"money"`
	Count uint   `gorm:"column:count" json:"count"`
}

func (Orders) TableName() string {
	return "orders"
}

type OrdersDetail struct {
	BaseModel
	OrdersID  uint   `gorm:"column:orders_id" json:"ordersID"` // 订单ID
	GoodsID   uint   `gorm:"column:goods_id" json:"goodsID"`   // 商品ID
	GoodsName string `gorm:"column:goods_name" json:"goodsName"`
	Barcode   string `gorm:"column:barcode" json:"barcode"`
	Brand     string `gorm:"column:brand" json:"brand"`
	Price     string `gorm:"column:price" json:"price"`
	Spec      string `gorm:"column:spec" json:"spec"`
	Count     uint   `gorm:"column:count" json:"count"` //商品计数
}

func (OrdersDetail) TableName() string {
	return "orders_detail"
}

func NewOrders() *Orders {
	return &Orders{}
}

func (u *Orders) List(p *dto.OrdersDto) (list []*Orders, err error) {
	if err = u.DB().Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&list).Error; err != nil {
		return
	}
	return
}

// 根据id获取订单详细信息
func (g *OrdersDetail) Detail(id uint) *OrdersDetail {
	if err := g.DB().Where("id", id).First(g).Error; err != nil {
		return nil
	}
	return g
}

// 创建订单信息
func (g *Orders) Create(req *dto.OrdersDto) (err error) {
	g.Money = req.Money
	g.Count = req.Count
	result := g.DB().Create(g)
	if result.Error != nil {
		return
	}
	var data []*OrdersDetail
	for _, v := range req.Detail {
		item := &OrdersDetail{
			OrdersID:  g.ID,
			GoodsID:   v.ID,
			GoodsName: v.GoodsName,
			Barcode:   v.Barcode,
			Price:     v.Price,
			Spec:      v.Spec,
			Count:     v.Count,
		}
		data = append(data, item)
	}

	//插入订单详细信息
	err = g.DB().CreateInBatches(data, len(data)).Error
	if err != nil {
		return
	}
	return result.Error
}

// 删除订单信息
func (g *Orders) Destory() error {
	result := g.DB().Delete(g)
	return result.Error
}

// 更新订单 信息
func (g *Orders) Update() error {
	result := g.DB().Save(g)
	return result.Error
}
