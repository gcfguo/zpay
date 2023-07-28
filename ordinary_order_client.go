package zpay

import (
	"errors"
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type OrdinaryOrderClient struct {
	*Client
}

// CreateOrder 创建订单
func (c *OrdinaryOrderClient) CreateOrder(
	data *model.CreateOrdinaryOrderReq) (*model.CreateOrdinaryOrderRes, error) {
	var res model.CreateOrdinaryOrderRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/order/ordinary/create",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}

// PayOrder 支付订单
func (c *OrdinaryOrderClient) PayOrder(
	data *model.PayOrdinaryOrderReq) (*model.PayOrdinaryOrderRes, error) {
	var res model.PayOrdinaryOrderRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/order/ordinary/pay",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}
