package model

type (
	LogisticsUploadReq struct {
		OutOrderNo string             `json:"out_order_no"`
		Couriers   []LogisticsCourier `json:"couriers"`
	}
	LogisticsCourier struct {
		CourierNumber string `json:"courier_number"`
		CourierCode   string `json:"courier_code"`
		URL           string `json:"url"`
	}
	LogisticsUploadRes struct{}
)
