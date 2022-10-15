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

type GoodBean struct {
	ID        uint   `json:"id"`
	Position  uint   `json:"position"` //在列表中的位置
	GoodsName string `json:"goodsName"`
	Barcode   string `json:"barcode"`
	Price     string `json:"price"`
	Brand     string `json:"brand"`
	Supplier  string `json:"supplier"`
	Standard  string `json:"standard"`
	Spec      string `json:"spec"`
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
func (g *Goods) GetGoods(barcode string) *Goods {
	if err := g.DB().Where("barcode", barcode).First(g).Error; err != nil {
		return nil
	}
	return g
}

// 注册商品信息
func (g *Goods) Register() error {
	result := g.DB().Create(g)
	go Goods2xlsx()
	return result.Error
}

// 删除商品信息
func (g *Goods) Destory() error {
	result := g.DB().Delete(g)
	go Goods2xlsx()
	return result.Error
}

// 更新商品信息
func (g *Goods) UpdatePrice(price string) error {
	result := g.DB().Model(g).Where("id = ?", g.ID).Update("price", price)
	go Goods2xlsx()
	return result.Error
}

// 更新商品信息
func (g *Goods) Update() error {
	result := g.DB().Save(g)
	go Goods2xlsx()
	return result.Error
}
