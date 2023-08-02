package zpay

import (
	"errors"
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type PaymentChannelsClient struct {
	*Client
}

// GetOrdinaryPaymentChannels 获取普通订单支付渠道
func (c *PaymentChannelsClient) GetOrdinaryPaymentChannels(
	data *model.GetOrdinaryPaymentChannelsReq) (*model.GetOrdinaryPaymentChannelsRes, error) {
	var res model.GetOrdinaryPaymentChannelsRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/payment/ordinary/channels/get",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}

// GetCombinedPaymentChannels 获取合单支付渠道
func (c *PaymentChannelsClient) GetCombinedPaymentChannels(
	data *model.GetCombinedPaymentChannelsReq) (*model.GetCombinedPaymentChannelsRes, error) {
	var res model.GetCombinedPaymentChannelsRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/payment/combined/channels/get",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}
