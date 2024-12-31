package model

type (
	GetPayWaysReq struct {
		Currency       string `json:"currency"`
		CustomsCountry string `json:"customs_country"`
		TradingCountry string `json:"trading_country"`
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
		TradingCountries []string `json:"trading_countries"`
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
		Name     string         `json:"name"`
		Email    string         `json:"email"`
		Phone    string         `json:"phone"`
		Password string         `json:"password"`
		Wechat   *PayWayWechat  `json:"wechat,omitempty"`
		Paypal   *PayWayPaypal  `json:"paypal,omitempty"`
		Alipay   *PayWayAlipay  `json:"alipay,omitempty"`
		Useepay  *PayWayUseepay `json:"useepay,omitempty"`
		Stripe   *PayWayStripe  `json:"stripe,omitempty"`
		Onepay   *PayWayOnepay  `json:"onepay,omitempty"`
	}
	PayWayWechat struct {
		SubMchID         string   `json:"sub_mch_id"`
		TradingCountries []string `json:"trading_countries"`
	}
	PayWayPaypal struct {
		ClientID         string   `json:"client_id"`
		ClientSecret     string   `json:"client_secret"`
		IsSandBox        bool     `json:"is_sand_box"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
	}
	PayWayAlipay struct {
		AppID            string   `json:"app_id"`
		MchID            string   `json:"mch_id"`
		PrivateKey       string   `json:"private_key"`
		PublicKey        string   `json:"public_key"`
		SandboxSwitch    int      `json:"sandbox_switch"`
		TradingCountries []string `json:"trading_countries"`
	}
	PayWayUseepay struct {
		MerchantNo       string   `json:"merchant_no"`
		AppID            string   `json:"app_id"`
		SecretKey        string   `json:"secret_key"`
		IsSandBox        bool     `json:"is_sand_box"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
	}
	PayWayStripe struct {
		SecretKey        string   `json:"secret_key"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
	}
	PayWayOnepay struct {
		AppID            string   `json:"app_id"`
		MerchantSelfCode string   `json:"merchant_self_code"`
		ShopSelfCode     string   `json:"shop_self_code"`
		PrivateKey       string   `json:"private_key"`
		PublicKey        string   `json:"public_key"`
		GatewayUrl       string   `json:"gateway_url"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
	}
	AddPayWayRes struct {
		ZPayMchID uint64 `json:"zpay_mch_id"`
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}
)

type (
	RemovePaymentWayReq struct {
		PaymentWays []string `json:"payment_ways"`
	}
	RemovePaymentWayRes struct {
	}
)

type (
	ShowPaymentChannelRes struct {
		Paypal  *ChannelPaypal  `json:"paypal,omitempty"`
		Alipay  *ChannelAlipay  `json:"alipay,omitempty"`
		Wechat  *ChannelWechat  `json:"wechat,omitempty"`
		Useepay *ChannelUseepay `json:"useepay,omitempty"`
		Stripe  *ChannelStripe  `json:"stripe,omitempty"`
	}
	ChannelPaypal struct {
		ClientId         string   `json:"client_id"      `
		ClientSecret     string   `json:"client_secret"  `
		IsSandbox        bool     `json:"is_sandbox" `
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
	ChannelAlipay struct {
		MchId            string   `json:"mch_id"         `
		AppId            string   `json:"app_id"         `
		PrivateKey       string   `json:"private_key"    `
		PublicKey        string   `json:"public_key"     `
		IsSandbox        bool     `json:"is_sandbox" `
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
	ChannelWechat struct {
		MchId            string   `json:"mch_id"`
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
	ChannelUseepay struct {
		MerchantNo       string   `json:"merchant_no"    `
		AppId            string   `json:"app_id"         `
		SignType         string   `json:"sign_type"      `
		SecretKey        string   `json:"secret_key"     `
		IsSandbox        bool     `json:"is_sandbox" `
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
	ChannelStripe struct {
		SecretKey        string   `json:"secret_key"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
	ChannelOnePay struct {
		ZpayMchId        string   `json:"zpay_mch_id"`
		AppID            string   `json:"app_id"`
		MerchantSelfCode string   `json:"merchant_self_code"`
		ShopSelfCode     string   `json:"shop_self_code"`
		PrivateKey       string   `json:"private_key"`
		PublicKey        string   `json:"public_key"`
		GatewayUrl       string   `json:"gateway_url"`
		Currencies       []string `json:"currencies"`
		TradingCountries []string `json:"trading_countries"`
		InUse            bool     `json:"in_use"`
	}
)
