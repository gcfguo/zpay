package zpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
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
