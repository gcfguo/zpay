package model

type (
	CreateOrdinaryOrderReq struct {
		OrderID          string   `json:"order_id"`
		OrderFee         string   `json:"order_fee"`
		Currency         string   `json:"currency"`
		CallbackURL      string   `json:"callback_url"`
		Description      string   `json:"description"`
		Attach           string   `json:"attach"`
		ExpiresInSeconds int      `json:"expires_in_seconds"`
		Payer            *Payer   `json:"payer"`
		Goods            []*Goods `json:"goods"`
	}
	CreateOrdinaryOrderRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			PrepayID   string `json:"prepay_id"`
			CreateTime string `json:"create_time"`
		} `json:"data"`
	}
)

type (
	PayOrdinaryOrderReq struct {
		PrepayID string `json:"prepay_id"`
		PayWayID uint64 `json:"pay_way_id"`
		// SceneInfo 支付场景信息
		SceneInfo *SceneInfo `json:"scene_info"`
	}
	PayOrdinaryOrderRes struct {
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
	RefundOrdinaryOrderReq struct {
		OrderID      string         `json:"order_id"`
		RefundNo     string         `json:"refund_no"`
		RefundAmount string         `json:"refund_amount"`
		Reason       string         `json:"reason"`
		CallbackURL  string         `json:"callback_url"`
		RefundGoods  []*RefundGoods `json:"refund_goods"`
	}
	RefundOrdinaryOrderRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderStatus string `json:"order_status"`
			CreateTime  string `json:"create_time"`
		} `json:"data"`
	}
)
