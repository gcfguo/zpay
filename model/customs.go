package model

import "fmt"

type (
	CustomsDeclareReq struct {
		ZPayMchID  string `json:"zpay_mch_id"`
		OutOrderNo string `json:"out_order_no"`
		Customs    string `json:"customs"`
		ActionType string `json:"action_type"`
		CertType   string `json:"cert_type"`
		CertID     string `json:"cert_id"`
		Name       string `json:"name"`
		//拆单必传信息
		SubOrderNo   string `json:"sub_order_no"`
		FeeType      string `json:"fee_type"`
		SubOrderFee  string `json:"sub_order_fee"`
		TransportFee string `json:"transport_fee"`
		ProductFee   string `json:"product_fee"`
	}
	CustomsDeclareRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func (r *CustomsDeclareRes) Ok() bool {
	return r.Code == 0
}

func (r *CustomsDeclareRes) Error() error {
	return fmt.Errorf(r.Msg)
}
