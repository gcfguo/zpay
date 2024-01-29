package zpay

import (
	"context"

	"net/http"
	"testing"
	"time"

	"github.com/gcfguo/zpay/model"
	"github.com/gcfguo/zpay/util"
)

const appID = "your app_id"
const appSecret = "your app_secret"
const accessToken = "your access_token"

func TestClient_Register(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	_, err = client.Register(&model.RegisterReq{
		Name:     "2227309180@qq.com",
		Email:    "2227309180@qq.com",
		Phone:    "13888888888",
		Password: "13999999999",
	})
	if err != nil {
		t.Fatal("注册失败:", err)
	}
	t.Log("注册成功")
}

func TestClient_Login(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.Login(&model.LoginReq{
		Email:    "2227309180@qq.com",
		Password: "13999999999",
	})
	if err != nil {
		t.Fatal("登录失败:", err)
	}
	t.Log("登录成功:", util.JSON(got))
}

func TestAuthClient_GetAccessToken(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithInstantAuth(appID, appSecret),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	authClient := &AuthClient{client}
	accessToken, err := authClient.GetAccessToken()
	if err != nil {
		t.Fatal("获取access_token失败:", err)
	}
	t.Log("access_token.access_token:", accessToken.AccessToken)
}

func TestClient_Payment(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.Payment(&model.PaymentReq{
		MasterOrderNo:    "2023080712357",
		TotalAmount:      "0.01",
		Attach:           "测试",
		Description:      "测试",
		Currency:         "CNY",
		ExpiresInSeconds: 7200,
		SubOrders: []*model.SubOrder{
			{
				ZPayMchID:   0,
				SubOrderNo:  "202308071358",
				Description: "测试",
				Attach:      "测试",
				CallbackURL: "https://localhost:8888/v1/open/testcallback",
			},
		},
		PayWayID: "2837253544757690459",
		SceneInfo: &model.SceneInfo{
			SceneType: "NATIVE",
		},
	})
	if err != nil {
		t.Fatal("支付失败:", err)
	}
	t.Log("支付成功:", util.JSON(got))
}

func TestClient_GetApps(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken("your_access_token"),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.GetApps(&model.GetAppsReq{})
	if err != nil {
		t.Fatal("获取应用失败:", err)
	}
	t.Log("获取到的应用:", util.JSON(got))
}

func TestClient_CreateApp(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken("your_access_token"),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	_, err = client.CreateApp(&model.CreateAppReq{
		OwnerMchID:  0,
		Description: "蜂洞独立站",
		PaymentMode: 1,
	})
	if err != nil {
		t.Fatal("创建失败:", err)
	}
	t.Log("创建应用成功")
}

func TestClient_AppSettings(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken("your_access_token"),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	_, err = client.AppSettings(&model.AppSettingsReq{
		PayModeSettings: 1,
		AppID:           "your_appid",
		PayWaySettings:  []string{"your_pay_way_id1", "your_pay_way_id2"},
	})
	if err != nil {
		t.Fatal("应用设置失败:", err)
	}
	t.Log("应用设置成功")
}

func TestClient_CreatePayPalChannel(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken("your_access_token"),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.CreatePayPalChannel(&model.CreatePayPalChannelReq{
		ClientID:     "your_paypal_client_id",
		ClientSecret: "your_paypal_client_secret",
		IsSandBox:    false,
	})
	if err != nil {
		t.Fatal("创建paypal支付渠道失败:", err)
	}
	t.Log("创建paypal支付渠道成功:", util.JSON(got))
}

func TestClient_CreatePayWay(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiIiLCJtY2hfaWQiOjI4NDc4MDUzNDE4NDI4MDE0NjIsInN1YiI6InpQYXkiLCJleHAiOjE2OTc0NDQ1NTh9.QGtpAfedPrArI2lVl5-yiRIZScNDYIqRZtW73kG0BWt23kkwy1kdHLsRsVczaDiRaZ88BHCxmatvUlDYHI0iYp25CranJXVSCKPl7-st-UiUiajeMHG-NWCoXWmGgjzy0bOuX6LXWoX0YDwPY73tnimecp_2YzpMsmfDc-i9zaE"),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.CreatePayWay(&model.CreatePayWayReq{
		OwnerMchID:             2847805341842801462,
		ChannelID:              4029586600,
		ChannelType:            "paypal",
		CurrencySettings:       `["CNY","USD"]`,
		CustomsCountrySettings: `["CN"]`,
	})
	if err != nil {
		t.Fatal("创建支付方式失败:", err)
	}
	t.Log("创建支付方式成功:", util.JSON(got))
}

func TestClient_GetPayWays(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	got, err := client.GetPayWays(&model.GetPayWaysReq{})
	if err != nil {
		t.Fatal("获取支付方式失败:", err)
	}
	t.Log("获取到的支付方式:", util.JSON(got))
}

func TestClient_ShowPaymentChannel(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("err0:", err)
	}
	got, err := client.ShowPaymentChannel()
	if err != nil {
		t.Fatal("err1:", err)
	}
	t.Log("got:", util.JSON(got))
}
