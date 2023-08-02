package zpay

import "net/http"

type ClientOption interface {
	Apply(settings *DialSettings) error
}

type httpClientOption struct {
	HttpClient *http.Client
}

func (h *httpClientOption) Apply(settings *DialSettings) error {
	settings.HTTPClient = h.HttpClient
	return nil
}

// WithHttpClient 使用自定的http client
func WithHttpClient(httpClient *http.Client) ClientOption {
	return &httpClientOption{HttpClient: httpClient}
}

type accessTokenOption struct {
	Token string
}

func (a *accessTokenOption) Apply(settings *DialSettings) error {
	settings.AccessToken = a.Token
	return nil
}

// WithAccessToken 已有token,直接使用
func WithAccessToken(token string) ClientOption {
	return &accessTokenOption{Token: token}
}

type instantAuthOption struct {
	AppID     string
	AppSecret string
}

func (a *instantAuthOption) Apply(settings *DialSettings) error {
	settings.AppID = a.AppID
	settings.AppSecret = a.AppSecret
	return nil
}

// WithInstantAuth 没有token,即时授权,根据app_id和app_secret获取token
func WithInstantAuth(appID, appSecret string) ClientOption {
	return &instantAuthOption{AppID: appID, AppSecret: appSecret}
}

type urlOption struct {
	URL string
}

func (u *urlOption) Apply(settings *DialSettings) error {
	settings.URL = u.URL
	return nil
}

func WithURL(url string) ClientOption {
	return &urlOption{URL: url}
}
