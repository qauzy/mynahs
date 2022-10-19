package vo

type OrdersVo struct {
	ID       uint              `json:"id"`
	Position uint              `json:"position"` //在列表中的位置
	Money    string            `json:"money"`
	Count    uint              `json:"count"`
	Detail   []*OrdersDetailVo `json:"detail"`
}

type OrdersDetailVo struct {
	ID        uint   `json:"id"`
	GoodsName string `json:"goodsName"`
	Barcode   string `json:"barcode"`
	Price     string `json:"price"`
	Count     uint   `json:"count"`
	Supplier  string `json:"supplier"`
	Standard  string `json:"standard"`
	Spec      string `json:"spec"`
}
