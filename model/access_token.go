package model

type (
	GetAccessTokenReq struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}
	GetAccessTokenRes struct {
		ZPayMchID      uint64 `json:"zpay_mch_id"`
		TokenType      string `json:"token_type"`
		AccessToken    string `json:"access_token"`
		AuthorizedTime string `json:"authorized_time"`
	}
)
