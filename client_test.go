package zpay

import (
	"context"
	"github.com/gcfguo/zpay/model"
	"github.com/gcfguo/zpay/util"
	"net/http"
	"testing"
	"time"
)

const appID = "your app_id"
const appSecret = "your app_secret"
const accessToken = "your access_token"

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
				ZPayMchID:      0,
				SubOrderNo:     "202308071358",
				Description:    "测试",
				Attach:         "测试",
				InherentAmount: "0.01",
				AdditionalFee:  "0",
				CallbackURL:    "https://localhost:8888/v1/open/testcallback",
			},
		},
		PayWayID: 2837253544757690459,
		SceneInfo: &model.SceneInfo{
			SceneType: "NATIVE",
		},
	})
	if err != nil {
		t.Fatal("支付失败:", err)
	}
	t.Log("支付成功:", util.JSON(got))
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
	got, err := client.GetPayWays(&model.GetPayWaysReq{
		Currency: "CNY",
	})
	if err != nil {
		t.Fatal("获取支付方式失败:", err)
	}
	t.Log("获取到的支付方式:", util.JSON(got))
}
