package zpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gcfguo/zpay/model"
)

type Client struct {
	ctx         context.Context
	url         string
	httpClient  *http.Client
	accessToken string
	appID       string
	appSecret   string
}

func NewClient(
	ctx context.Context,
	opts ...ClientOption) (*Client, error) {
	var settings = &DialSettings{
		HTTPClient: http.DefaultClient,
		URL:        "http://127.0.0.1:8888",
	}
	for _, v := range opts {
		err := v.Apply(settings)
		if err != nil {
			return nil, err
		}
	}
	return &Client{
		ctx:         ctx,
		url:         settings.URL,
		httpClient:  settings.HTTPClient,
		accessToken: settings.AccessToken,
		appID:       settings.AppID,
		appSecret:   settings.AppSecret,
	}, nil
}

func (c *Client) check() error {
	if c.httpClient == nil {
		return errors.New("http client is must for DialSettings")
	}
	if c.accessToken == "" && (c.appID == "" || c.appSecret == "") {
		return errors.New("access_token is must for DialSettings")
	}
	return nil
}

func (c *Client) newRequest(
	httpMethod string,
	reqURI string,
	data any) (*http.Request, error) {
	reqBody, err := c.writeRequestBody(data)
	if err != nil {
		return nil, err
	}
	reqURL := c.url + reqURI
	req, err := http.NewRequest(httpMethod, reqURL, reqBody)
	return req, err
}

func (c *Client) doRequest(
	httpMethod string,
	reqURI string,
	data any,
	scanner any) (*string, error) {
	req, err := c.newRequest(httpMethod, reqURI, data)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	content := string(b)
	if scanner == nil {
		return &content, nil
	}
	err = json.Unmarshal(b, scanner)
	return &content, err
}

func (c *Client) doRequestWithToken(
	httpMethod string,
	reqURI string,
	data any,
	scanner any) (*string, error) {
	err := c.check()
	if err != nil {
		return nil, err
	}
	if c.accessToken == "" {
		authClient := AuthClient{c}
		token, err := authClient.GetAccessToken()
		if err != nil {
			return nil, err
		}
		c.accessToken = token.AccessToken
	}
	req, err := c.newRequest(httpMethod, reqURI, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	content := string(b)
	if scanner == nil {
		return &content, nil
	}
	err = json.Unmarshal(b, scanner)
	return &content, err
}

func (c *Client) writeRequestBody(data any) (io.Reader, error) {
	var rBody io.Reader
	switch data.(type) {
	case string:
		rBody = strings.NewReader(data.(string))
	case []byte:
		rBody = bytes.NewReader(data.([]byte))
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		rBody = bytes.NewReader(b)
	}
	return rBody, nil
}

func (c *Client) handleResponse(content []byte, result Result) error {
	err := json.Unmarshal(content, result)
	if err != nil {
		return err
	}
	if !result.Ok() {
		return result.Error()
	}
	return nil
}

// Register
// 注册
func (c *Client) Register(req *model.RegisterReq) (*model.RegisterRes, error) {
	resContent, err := c.doRequest(http.MethodPost, "/v1/open/merchant/register", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.RegisterRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Login
// 登录
func (c *Client) Login(req *model.LoginReq) (*model.LoginRes, error) {
	resContent, err := c.doRequest(http.MethodPost, "/v1/open/merchant/login", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.LoginRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetApps
// 获取应用
func (c *Client) GetApps(req *model.GetAppsReq) (*model.GetAppsRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodGet, "/v1/api/app/get", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.GetAppsRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateApp
// 创建应用
func (c *Client) CreateApp(req *model.CreateAppReq) (*model.CreateAppRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/app/create", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.CreateAppRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AppSettings
// 应用设置
func (c *Client) AppSettings(req *model.AppSettingsReq) (*model.AppSettingsRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/app/settings", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.AppSettingsRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreatePayPalChannel
// 创建paypal支付渠道
func (c *Client) CreatePayPalChannel(req *model.CreatePayPalChannelReq) (*model.CreatePayPalChannelRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/paypal/create", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.CreatePayPalChannelRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreatePayWay
// 创建支付方式
func (c *Client) CreatePayWay(req *model.CreatePayWayReq) (*model.CreatePayWayRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/create", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.CreatePayWayRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Payment
// 支付
func (c *Client) Payment(req *model.PaymentReq) (*model.PaymentRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/order/payment", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.PaymentRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Refund
// 退款
func (c *Client) Refund(req *model.RefundReq) (*model.RefundRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/order/refund", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.RefundRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPayWays
// 获取支付方式
func (c *Client) GetPayWays(req *model.GetPayWaysReq) (
	*model.GetPayWaysRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/get", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.GetPayWaysRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CustomsDeclare
// 订单推送海关
func (c *Client) CustomsDeclare(req *model.CustomsDeclareReq) (
	*model.CustomsDeclareRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/customs/declare", req, nil)
	if err != nil {
		return nil, err
	}
	res := new(model.CustomsDeclareRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CustomsRedeclare
// 订单重新推送海关
func (c *Client) CustomsRedeclare() {

}

// CustomsQuery
// 订单推送查询
func (c *Client) CustomsQuery() {

}
