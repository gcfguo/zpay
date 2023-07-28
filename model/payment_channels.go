package model

type (
	GetPaymentChannelsReq struct {
		PrepayID string `json:"prepay_id"`
	}
	GetPaymentChannelsRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List  []*PaymentChannels `json:"list"`
			Total int                `json:"total"`
		} `json:"data"`
	}
)

type PaymentChannels struct {
	PayWayId    uint64 `json:"pay_way_id"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
}
