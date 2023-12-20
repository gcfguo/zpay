package zpay

import (
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type AuthClient struct {
	*Client
}

func (a *AuthClient) GetAccessToken() (*model.GetAccessTokenRes, error) {
	var res model.GetAccessTokenRes
	_, err := a.doRequest(http.MethodPost,
		"/v1/open/merchant/authorize",
		&model.GetAccessTokenReq{
			AppID:     a.appID,
			AppSecret: a.appSecret,
		},
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
