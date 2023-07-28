package zpay

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gcfguo/zpay/model"
)

const appID = "your app_id"
const appSecret = "your app_secret"

func TestAuthClient_GetAccessToken(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAuthGet(appID, appSecret),
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

func TestOrdinaryOrderClient_CreateOrder(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAuthGet(appID, appSecret),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	ordinaryClient := &OrdinaryOrderClient{client}
	got, err := ordinaryClient.CreateOrder(&model.CreateOrdinaryOrderReq{
		OrderID:         "zpaytest11111",
		OrderFee:        "0.01",
		Currency:        "CNY",
		CallbackURL:     "http://127.0.0.1:8888",
		Description:     "测试订单",
		ExpiresInSecond: 7200,
		Goods: []*model.Goods{
			{
				GoodsID:   "122222",
				SkuID:     "211111111",
				UnitPrice: "0.01",
				Quantity:  1,
			},
		},
	})
	if err != nil {
		t.Fatal("创建普通订单失败:", err)
	}
	t.Log("预支付单号:", got.Data.PrepayID)
}

func TestPaymentChannelsClient_GetPaymentChannels(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAuthGet(appID, appSecret),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	channelsClient := &PaymentChannelsClient{client}
	got, err := channelsClient.GetPaymentChannels(&model.GetPaymentChannelsReq{
		PrepayID: "2TBBip3QDrubsP2FmEaJKKbFqcv",
	})
	if err != nil {
		t.Fatal("获取支付渠道失败:", err)
	}
	t.Log("支付渠道信息:", got.Data.List[0].PayWayId)
}

func TestOrdinaryOrderClient_PayOrder(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAuthGet(appID, appSecret),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	ordinaryClient := &OrdinaryOrderClient{client}
	got, err := ordinaryClient.PayOrder(&model.PayOrdinaryOrderReq{
		PrepayID:  "2TBBip3QDrubsP2FmEaJKKbFqcv",
		PayWayID:  2835659989484307265,
		SceneType: "NATIVE",
	})
	if err != nil {
		t.Fatal("普通订单支付失败:", err)
	}
	t.Log("收银台请求信息:", got.Data.Native.QRCodeSummary)
}
