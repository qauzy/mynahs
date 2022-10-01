package model

type Goods struct {
	BaseModel
	GoodsName string `gorm:"column:goods_name" json:"goodsName"`
	Barcode   string `gorm:"column:barcode" json:"barcode"`
	Brand     string `gorm:"column:brand" json:"brand"`
	Price     string `gorm:"column:price" json:"price"`
	Spec      string `gorm:"column:spec" json:"spec"`
	Supplier  string `gorm:"column:supplier" json:"supplier"`
}

func (Goods) TableName() string {
	return "goods"
}

func NewGoods() *Goods {
	return &Goods{}
}

func NewGoodsEx(GoodsName, Barcode, Brand, Price, Spec, Supplier string) *Goods {
	return &Goods{
		GoodsName: GoodsName,
		Barcode:   Barcode,
		Price:     Price,
		Brand:     Brand,
		Spec:      Spec,
		Supplier:  Supplier,
	}
}
func (u *Goods) GetGoodsList(p *Page) (list []*Goods, err error) {
	if err = u.DB().Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&list).Error; err != nil {
		return
	}
	return
}

// GetUserInfo 根据barcode获取商品信息
func (u *Goods) GetGoods(barcode string) *Goods {
	if err := u.DB().Where("barcode", barcode).First(u).Error; err != nil {
		return nil
	}
	return u
}

// 注册商品信息
func (g *Goods) Register() error {
	result := g.DB().Create(g)
	return result.Error
}

// 删除商品信息
func (g *Goods) Destory() error {
	result := g.DB().Delete(g)
	return result.Error
}

// 更新商品信息
func (g *Goods) UpdatePrice(id int64, price string) error {
	result := g.DB().Where("id = ?", id).Update("price", price)
	return result.Error
}

// 更新商品信息
func (g *Goods) Update() error {
	result := g.DB().Save(g)
	return result.Error
}
