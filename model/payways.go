package model

import "fmt"

type (
	GetPayWaysReq struct {
		Currency       string `json:"currency"`
		CustomsCountry string `json:"customs_country"`
	}
	GetPayWaysRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List  []*PayWays `json:"list"`
			Total int        `json:"total"`
		}
	}
)

func (r *GetPayWaysRes) Ok() bool {
	return r.Code == 0
}

func (r *GetPayWaysRes) Error() error {
	return fmt.Errorf(r.Msg)
}

type PayWays struct {
	PayWayId    uint64 `json:"pay_way_id"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}
