package model

import "fmt"

type (
	PaymentReq struct {
		MasterOrderNo    string      `json:"master_order_no" v:"bail|required|min-length:1"`
		TotalAmount      string      `json:"total_amount" v:"bail|required|min-length:1"`
		Attach           string      `json:"attach"`
		Description      string      `json:"description" v:"bail|required|min-length:1"`
		Currency         string      `json:"currency" v:"bail|required|min-length:1"`
		ExpiresInSeconds int         `json:"expires_in_seconds" v:"bail|required|min:1"`
		SubOrders        []*SubOrder `json:"sub_orders" v:"bail|required|min-length:1"`
		PayWayID         uint64      `json:"pay_way_id" v:"bail|required|min:1"`
		SceneInfo        *SceneInfo  `json:"scene_info" v:"bail|required"`
	}
	PaymentRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			H5     *H5Result     `json:"h5,omitempty"`
			Native *NativeResult `json:"native,omitempty"`
			JSAPI  *JSAPIResult  `json:"jsapi,omitempty"`
		} `json:"data"`
	}
)

func (s *PaymentRes) Ok() bool {
	return s.Code == 0
}

func (s *PaymentRes) Error() error {
	return fmt.Errorf(s.Msg)
}

type SubOrder struct {
	ZPayMchID      uint64 `json:"zpay_mch_id"`
	SubOrderNo     string `json:"sub_order_no" v:"bail|required|min-length:1"`
	Description    string `json:"description" v:"bail|required|min-length:1"`
	Attach         string `json:"attach"`
	InherentAmount string `json:"inherent_amount" v:"bail|required|min-length:1"`
	AdditionalFee  string `json:"additional_fee" v:"bail|required|min-length:1"`
	CallbackURL    string `json:"callback_url" v:"bail|required|min-length:1"`
}

type SceneInfo struct {
	//下单场景类型
	SceneType string `json:"scene_type" v:"bail|required|in:NATIVE"`
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
}

type H5Result struct {
	H5URL string `json:"h5_url"`
}

type NativeResult struct {
	QRCodeSummary string `json:"qr_code_summary"`
}

type JSAPIResult struct {
	AppID     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}
