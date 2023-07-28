package model

type Goods struct {
	GoodsID   string `json:"goods_id"`
	SkuID     string `json:"sku_id"`
	UnitPrice string `json:"unit_price"`
	Quantity  int    `json:"quantity"`
}
