package model

type (
	PaymentReq struct {
		MasterOrderNo    string      `json:"master_order_no"`
		TotalAmount      string      `json:"total_amount"`
		Attach           string      `json:"attach"`
		Description      string      `json:"description"`
		Currency         string      `json:"currency"`
		ExpiresInSeconds int         `json:"expires_in_seconds"`
		SubOrders        []*SubOrder `json:"sub_orders"`
		PayWayID         string      `json:"pay_way_id"`
		SceneInfo        *SceneInfo  `json:"scene_info"`
		BillingAddress   *Address    `json:"billing_address"`
		ShippingAddress  *Address    `json:"shipping_address"`
	}
	PaymentRes struct {
		H5     *H5Result     `json:"h5,omitempty"`
		Native *NativeResult `json:"native,omitempty"`
		JSAPI  *JSAPIResult  `json:"jsapi,omitempty"`
	}
)

type SubOrder struct {
	ZPayMchID        uint64  `json:"zpay_mch_id"`
	SubOrderNo       string  `json:"sub_order_no"`
	Description      string  `json:"description"`
	Attach           string  `json:"attach"`
	GoodsTotalAmount string  `json:"goods_total_amount"`
	Discount         string  `json:"discount"`
	ShippingFee      string  `json:"shipping_fee"`
	TaxFee           string  `json:"tax_fee"`
	InsuranceFee     string  `json:"insurance_fee"`
	CallbackURL      string  `json:"callback_url"`
	Items            []*Item `json:"items"`
}

type Item struct {
	SkuID       string `json:"sku_id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UnitPrice   string `json:"unit_price"`
	Quantity    int    `json:"quantity"`
}

type SceneInfo struct {
	//下单场景类型
	SceneType string `json:"scene_type"`
	//DeviceID 商户端设备号
	DeviceID string `json:"device_id"`
	//PayerClientIP 用户终端IP
	PayerClientIP string `json:"payer_client_ip"`
	//H5Type H5支付场景类型
	H5Type string `json:"h5_type"`
	//H5AppName 应用名称
	H5AppName string `json:"h5_app_name"`
	//H5AppURL 网站URL
	H5AppURL string `json:"h5_app_url"`
	//ReturnURL 支付后跳转地址
	ReturnURL string `json:"return_url"`
	//CancelURL 支付取消后跳转地址
	CancelURL string `json:"cancel_url"`
}

type H5Result struct {
	H5URL string `json:"h5_url"`
}

type NativeResult struct {
	QRCodeSummary string `json:"qr_code_summary"`
}

type JSAPIResult struct {
	AuthLink  string `json:"auth_link,omitempty"`
	AppID     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

type Address struct {
	HouseNo    string `json:"house_no,omitempty"`
	Email      string `json:"email"`
	PhoneNo    string `json:"phone_no,omitempty"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city,omitempty"`
	State      string `json:"state"`
	Country    string `json:"country"`
}
