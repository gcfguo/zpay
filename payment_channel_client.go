package zpay

import (
	"errors"
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type PaymentChannelsClient struct {
	*Client
}

// GetPaymentChannels 获取支付渠道
func (c *PaymentChannelsClient) GetPaymentChannels(
	data *model.GetPaymentChannelsReq) (*model.GetPaymentChannelsRes, error) {
	var res model.GetPaymentChannelsRes
	_, err := c.doRequestWithToken(http.MethodPost,
		"/v1/api/payment/channels/get",
		data,
		&res,
	)
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, err
}
