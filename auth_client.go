package zpay

import (
	"errors"
	"net/http"

	"github.com/gcfguo/zpay/model"
)

type AuthClient struct {
	*Client
}

func (a *AuthClient) GetAccessToken() (*model.AccessToken, error) {
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
	//这儿可以优化为集中处理逻辑,暂时没时间
	if res.Code != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res.Data, nil
}
