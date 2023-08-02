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
const accessToken = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiIwMDczMDY5MERBQTFDMEZFQTNGNDIxRDg4MTUwQ0Q1QiIsIm1jaF9pZCI6MjgzNTYyODg3OTk2MjQ0MDUxMywic3ViIjoielBheSIsImV4cCI6MTY5MDk0ODI2MH0.ezoo4IV6RxBwjA1jQs8yqDn1nR7B5uARz8YrMYLAAMq9SA9EZ-0jxKiiHl8FXM_O6ODKQ6E6UDj1igAvzxQ2nGdvPDlOP2A-xRDuU7qNVbPtveb4ya6nCkjjWz1bEafWoY_92fceO2V1ezZ4VH8vozSaRrfweHhuW4fQT6aqueY"

// 测试获取access_token
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

// 测试创建普通订单
func TestOrdinaryOrderClient_CreateOrder(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	ordinaryClient := &OrdinaryOrderClient{client}
	got, err := ordinaryClient.CreateOrder(&model.CreateOrdinaryOrderReq{
		OrderID:          "zpaytest11116",
		OrderFee:         "0.01",
		Currency:         "CNY",
		CallbackURL:      "http://127.0.0.1:8888",
		Description:      "测试订单",
		ExpiresInSeconds: 7200,
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

// 测试获取普通订单支付渠道
func TestPaymentChannelsClient_GetPaymentChannels(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	channelsClient := &PaymentChannelsClient{client}
	got, err := channelsClient.GetOrdinaryPaymentChannels(
		&model.GetOrdinaryPaymentChannelsReq{
			PrepayID: "2TPPlOfPIdPXpBuGt8ByBBx8ZUX",
		})
	if err != nil {
		t.Fatal("获取支付渠道失败:", err)
	}
	t.Log("支付渠道信息:", got.Data.List[0].PayWayId)
}

// 测试普通订单支付
func TestOrdinaryOrderClient_PayOrder(t *testing.T) {
	client, err := NewClient(
		context.Background(),
		WithHttpClient(&http.Client{Timeout: 5 * time.Second}),
		WithAccessToken(accessToken),
	)
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	ordinaryClient := &OrdinaryOrderClient{client}
	got, err := ordinaryClient.PayOrder(&model.PayOrdinaryOrderReq{
		PrepayID: "2TPPlOfPIdPXpBuGt8ByBBx8ZUX",
		PayWayID: 2835659989484307265,
		SceneInfo: &model.SceneInfo{
			SceneType: "NATIVE",
		},
	})
	if err != nil {
		t.Fatal("普通订单支付失败:", err)
	}
	t.Log("收银台请求信息:", got.Data.Native.QRCodeSummary)
}

// 测试普通订单退款
func TestOrdinaryOrderClient_Refund(t *testing.T) {
	client, err := NewClient(context.Background(), WithAccessToken(accessToken))
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	ordinaryClient := &OrdinaryOrderClient{client}
	got, err := ordinaryClient.Refund(&model.RefundOrdinaryOrderReq{
		OrderID:      "zpaytest11116",
		RefundNo:     "zpaytest11117",
		RefundAmount: "0.01",
		Reason:       "不想要了",
		CallbackURL:  "https://www.baidu.com",
	})
	if err != nil {
		t.Fatal("普通订单退款失败:", err)
	}
	t.Log("退款结果:", util.JSON(got.Data))
}

// 测试创建合单
func TestCombinedOrderClient_CreateOrder(t *testing.T) {
	client, err := NewClient(context.Background(), WithAccessToken(accessToken))
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	combinedClient := &CombinedOrderClient{client}
	got, err := combinedClient.CreateOrder(&model.CreateCombinedOrderReq{
		CombinedOrderID:  "202308020942",
		Description:      "合单支付测试",
		ExpiresInSeconds: 7200,
		Currency:         "CNY",
		OrderFee:         "0.01",
		SubOrders: []*model.SubOrder{
			{
				ZPayMchID:   2836652182160277614,
				SubOrderID:  "202308020942",
				Description: "合单支付子订单",
				Attach:      "无备注",
				OrderFee:    "0.01",
				CallbackURL: "https://www.baidu.com",
				Goods: []*model.Goods{
					{
						GoodsID:   "xyz001",
						UnitPrice: "0.01",
						Quantity:  1,
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal("合单创建失败:", err)
	}
	t.Log("合单创建返回信息:", got.Data.PrepayID)
}

// 测试获取合单支付渠道
func TestPaymentChannelsClient_GetCombinedPaymentChannels(t *testing.T) {
	client, err := NewClient(context.Background(), WithAccessToken(accessToken))
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	channelClient := PaymentChannelsClient{client}
	got, err := channelClient.GetCombinedPaymentChannels(&model.GetCombinedPaymentChannelsReq{
		PrepayID: "2TPOLwWTtjMKSNTiwWSY7s78n0c",
	})
	if err != nil {
		t.Fatal("获取支付渠道失败:", err)
	}
	t.Log("支付渠道信息:", util.JSON(got.Data.List))
}

// 测试支付合单
func TestCombinedOrderClient_PayOrder(t *testing.T) {
	client, err := NewClient(context.Background(), WithAccessToken(accessToken))
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	combinedClient := &CombinedOrderClient{client}
	got, err := combinedClient.PayOrder(&model.PayCombinedOrderReq{
		PrepayID: "2TPOLwWTtjMKSNTiwWSY7s78n0c",
		PayWayID: 2836805454443578177,
		SceneInfo: &model.SceneInfo{
			SceneType: "NATIVE",
		},
	})
	if err != nil {
		t.Fatal("合单支付失败:", err)
	}
	t.Log("合单支付结果:", util.JSON(got.Data))
}

// 测试合单退款
func TestCombinedOrderClient_Refund(t *testing.T) {
	client, err := NewClient(context.Background(), WithAccessToken(accessToken))
	if err != nil {
		t.Fatal("初始化失败:", err)
	}
	combinedClient := &CombinedOrderClient{client}
	got, err := combinedClient.Refund(&model.RefundCombinedOrderReq{
		ZPaySubMchID: 0,
		OrderID:      "202308020942",
		RefundNo:     "202308020943",
		RefundAmount: "0.01",
		Reason:       "不想要了",
		CallbackURL:  "https://www.baidu.com",
		RefundGoods:  nil,
	})
	if err != nil {
		t.Fatal("合单退款失败:", err)
	}
	t.Log("退款结果:", util.JSON(got.Data))
}
