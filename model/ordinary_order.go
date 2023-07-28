package model

type (
	CreateOrdinaryOrderReq struct {
		OrderID         string   `json:"order_id"`
		OrderFee        string   `json:"order_fee"`
		Currency        string   `json:"currency"`
		CallbackURL     string   `json:"callback_url"`
		Description     string   `json:"description"`
		Attach          string   `json:"attach"`
		ExpiresInSecond int      `json:"expires_in_second"`
		Payer           *Payer   `json:"payer"`
		Goods           []*Goods `json:"goods"`
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
		//SceneType 场景信息 JSAPI,H5,NATIVE,APP
		SceneType string `json:"scene_type"`
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
