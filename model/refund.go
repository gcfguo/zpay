package model

import "fmt"

type (
	RefundReq struct {
		ZPayMchID    uint64 `json:"zpay_mch_id"`
		OrderNo      string `json:"order_no" v:"bail|required|min-length:1"`
		RefundNo     string `json:"refund_no" v:"bail|required|min-length:1"`
		RefundAmount string `json:"refund_amount" v:"bail|required|min-length:1"`
		Reason       string `json:"reason" v:"bail|required|min-length:1"`
		CallbackURL  string `json:"callback_url" v:"bail|required|min-length:1"`
	}
	RefundRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderStatus string `json:"order_status"`
			CreateTime  string `json:"create_time"`
		} `json:"data"`
	}
)

func (r *RefundRes) Ok() bool {
	return r.Code == 0
}

func (r *RefundRes) Error() error {
	return fmt.Errorf(r.Msg)
}
