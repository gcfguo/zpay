package zpay

import (
	"errors"
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type CombinedOrderClient struct {
	*Client
}

func (c *CombinedOrderClient) CreateOrder(
	data *model.CreateCombinedOrderReq) (*model.CreateCombinedOrderRes, error) {
	var res model.CreateCombinedOrderRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/order/combined/create",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}

func (c *CombinedOrderClient) PayOrder(
	data *model.PayCombinedOrderReq) (*model.PayCombinedOrderRes, error) {
	var res model.PayCombinedOrderRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/order/combined/pay",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}

// Refund 订单退款
func (c *CombinedOrderClient) Refund(
	data *model.RefundCombinedOrderReq) (*model.RefundCombinedOrderRes, error) {
	var res model.RefundCombinedOrderRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/order/combined/refund",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}
