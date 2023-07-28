package zpay

import (
	"net/http"
)

// DialSettings Client所需的设置信息
type DialSettings struct {
	//HTTPClient 自定义所需的http.Client实例
	HTTPClient *http.Client
	//AccessToken 请求所需的token
	AccessToken string
	//AppID 获取access_token所需的AppID
	AppID string
	//AppSecret 获取access_token所需的AppSecret
	AppSecret string
	//URL 自定义请求地址
	URL string
}
