package model

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
