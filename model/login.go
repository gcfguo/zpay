package model

import "fmt"

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			ZPayMchID      uint64 `json:"zpay_mch_id"`
			TokenType      string `json:"token_type"`
			AccessToken    string `json:"access_token"`
			AuthorizedTime string `json:"authorized_time"`
		}
	}
)

func (r *LoginRes) Ok() bool {
	return r.Code == 0
}

func (r *LoginRes) Error() error {
	return fmt.Errorf(r.Msg)
}
