package zpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
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
	data any) (*string, error) {
	req, err := c.newRequest(httpMethod, reqURI, data)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	for t := 0; t < 3; t++ {
		resp, _ = c.httpClient.Do(req)
		if resp != nil {
			break
		}
		var err net.Error
		if errors.As(err, &err) {
			continue
		}
	}
	if resp == nil {
		return nil, fmt.Errorf("too many failures")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	content := string(b)
	return &content, err
}

func (c *Client) doRequestWithToken(
	httpMethod string,
	reqURI string,
	data any) (*string, error) {
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

	var resp *http.Response
	for t := 0; t < 3; t++ {
		resp, _ = c.httpClient.Do(req)
		if resp != nil {
			break
		}
		var err net.Error
		if errors.As(err, &err) {
			continue
		}
	}
	if resp == nil {
		return nil, fmt.Errorf("too many failures")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	content := string(b)
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

func (c *Client) handleResponse(content []byte, data any) error {
	var result APIResult
	result.Data = data
	err := json.Unmarshal(content, &result)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf(result.Msg)
	}
	return nil
}

func (c *Client) GetAccessToken() string {
	return c.accessToken
}

// Register
// 注册
func (c *Client) Register(req *model.RegisterReq) (*model.RegisterRes, error) {
	resContent, err := c.doRequest(http.MethodPost, "/v1/open/merchant/register", req)
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
	resContent, err := c.doRequest(http.MethodPost, "/v1/open/merchant/login", req)
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
	resContent, err := c.doRequestWithToken(http.MethodGet, "/v1/api/app/get", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/app/create", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/app/settings", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/paypal/create", req)
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

// CreateWechatChannel
// 创建微信支付渠道
func (c *Client) CreateWechatChannel(req *model.CreateWechatSubChannelReq) (*model.CreateWechatSubChannelRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/wechat/sub/create", req)
	if err != nil {
		return nil, err
	}
	res := new(model.CreateWechatSubChannelRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreatePayWay
// 创建支付方式
func (c *Client) CreatePayWay(req *model.CreatePayWayReq) (*model.CreatePayWayRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/create", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/order/payment", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/order/refund", req)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/get", req)
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

func (c *Client) AddPayWay(req *model.AddPayWayReq) (
	*model.AddPayWayRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/add", req)
	if err != nil {
		return nil, err
	}
	res := new(model.AddPayWayRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) RemovePayWay(req *model.RemovePaymentWayReq) (
	*model.RemovePaymentWayRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/remove", req)
	if err != nil {
		return nil, err
	}
	res := new(model.RemovePaymentWayRes)
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
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/customs/declare", req)
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

// ShowPaymentChannel
// 展示支付渠道
func (c *Client) ShowPaymentChannel() (*model.ShowPaymentChannelRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/payway/channels", nil)
	if err != nil {
		return nil, err
	}
	res := new(model.ShowPaymentChannelRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) LogisticsUpload(req *model.LogisticsUploadReq) (*model.LogisticsUploadRes, error) {
	resContent, err := c.doRequestWithToken(http.MethodPost, "/v1/api/order/logistics/upload", req)
	if err != nil {
		return nil, err
	}
	res := new(model.LogisticsUploadRes)
	err = c.handleResponse([]byte(*resContent), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
