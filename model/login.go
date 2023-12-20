package model

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginRes struct {
		ZPayMchID      uint64 `json:"zpay_mch_id"`
		TokenType      string `json:"token_type"`
		AccessToken    string `json:"access_token"`
		AuthorizedTime string `json:"authorized_time"`
	}
)
