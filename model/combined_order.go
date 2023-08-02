package model

type (
	CreateCombinedOrderReq struct {
		CombinedOrderID string `json:"combined_order_id"`
		Description     string `json:"description"`
		// ExpiresInSeconds 过期时间
		ExpiresInSeconds int `json:"expires_in_seconds"`
		// Currency 交易货币编码,如:CNY
		Currency string `json:"currency"`
		// OrderFee 合单订单金额
		OrderFee string `json:"order_fee"`
		// SubOrders 子订单信息
		SubOrders []*SubOrder `json:"sub_orders"`
		// Payer 支付人证件信息
		Payer *Payer `json:"payer"`
	}
	CreateCombinedOrderRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			PrepayID   string `json:"prepay_id"`
			CreateTime string `json:"create_time"`
		} `json:"data"`
	}
)

type (
	PayCombinedOrderReq struct {
		PrepayID string `json:"prepay_id"`
		PayWayID uint64 `json:"pay_way_id"`
		// SceneInfo 支付场景信息
		SceneInfo *SceneInfo `json:"scene_info"`
	}
	PayCombinedOrderRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			H5     *H5Result     `json:"h5,omitempty"`
			Native *NativeResult `json:"native,omitempty"`
			JSAPI  *JSAPIResult  `json:"jsapi,omitempty"`
		} `json:"data"`
	}
)

type (
	RefundCombinedOrderReq struct {
		ZPaySubMchID uint64         `json:"zpay_sub_mch_id"`
		OrderID      string         `json:"order_id"`
		RefundNo     string         `json:"refund_no"`
		RefundAmount string         `json:"refund_amount"`
		Reason       string         `json:"reason"`
		CallbackURL  string         `json:"callback_url"`
		RefundGoods  []*RefundGoods `json:"refund_goods"`
	}
	RefundCombinedOrderRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderStatus string `json:"order_status"`
			CreateTime  string `json:"create_time"`
		} `json:"data"`
	}
)

type SubOrder struct {
	ZPayMchID   uint64 `json:"zpay_mch_id"`
	SubOrderID  string `json:"sub_order_id"`
	Description string `json:"description"`
	Attach      string `json:"attach"`
	OrderFee    string `json:"order_fee"`
	CallbackURL string `json:"callback_url"`
	// Goods 商品信息,暂不校验金额
	// # 商品1单价*商品1数量+商品n单价*商品n数量=订单支付金额
	Goods []*Goods `json:"goods"`
}
