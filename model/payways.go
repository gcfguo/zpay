package model

type (
	GetPayWaysReq struct {
		Currency       string `json:"currency"`
		CustomsCountry string `json:"customs_country"`
	}
	GetPayWaysRes struct {
		List  []*PayWays `json:"list"`
		Total int        `json:"total"`
	}
	PayWays struct {
		PayWayId         string   `json:"pay_way_id"`
		Logo             string   `json:"logo"`
		Description      string   `json:"description"`
		Currencies       []string `json:"currencies"`
		CustomsCountries []string `json:"customs_countries"`
	}
)

type (
	CreatePayWayReq struct {
		OwnerMchID             uint64 `json:"owner_mch_id"`
		ChannelID              uint64 `json:"channel_id"`
		ChannelType            string `json:"channel_type"`
		CurrencySettings       string `json:"currency_settings"`
		CustomsCountrySettings string `json:"customs_country_settings"`
	}
	CreatePayWayRes struct {
		PayWayID uint64 `json:"pay_way_id"`
	}
)

type (
	AddPayWayReq struct {
		Name     string        `json:"name"`
		Email    string        `json:"email"`
		Phone    string        `json:"phone"`
		Password string        `json:"password"`
		Wechat   *PayWayWechat `json:"wechat,omitempty"`
		Paypal   *PayWayPaypal `json:"paypal,omitempty"`
		Alipay   *PayWayAlipay `json:"alipay,omitempty"`
	}
	PayWayWechat struct {
		SubMchID string `json:"sub_mch_id"`
	}
	PayWayPaypal struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		IsSandBox    bool   `json:"is_sand_box"`
	}
	PayWayAlipay struct {
		AppID         string `json:"app_id"`
		MchID         string `json:"mch_id"`
		PrivateKey    string `json:"private_key"`
		PublicKey     string `json:"public_key"`
		SandboxSwitch int    `json:"sandbox_switch"`
	}
	AddPayWayRes struct {
	}
)
