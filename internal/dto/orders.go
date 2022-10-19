package dto

type OrdersDetailDto struct {
	ID        uint   `json:"id"`
	GoodsName string `json:"goodsName"`
	Barcode   string `json:"barcode"`
	Price     string `json:"price"`
	Count     uint   `json:"count"`
	Supplier  string `json:"supplier"`
	Standard  string `json:"standard"`
	Spec      string `json:"spec"`
}

type OrdersDto struct {
	Page
	Position uint               `json:"position"` //在列表中的位置
	ID       uint               `json:"id"`
	Money    string             `json:"money"`
	Count    uint               `json:"count"`
	Detail   []*OrdersDetailDto `json:"detail"`
}
