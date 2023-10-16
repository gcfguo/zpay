package model

import "fmt"

type (
	GetPayWaysReq struct {
		Currency       string `json:"currency"`
		CustomsCountry string `json:"customs_country"`
	}
	GetPayWaysRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List  []*PayWays `json:"list"`
			Total int        `json:"total"`
		}
	}
	PayWays struct {
		PayWayId         string   `json:"pay_way_id"`
		Logo             string   `json:"logo"`
		Description      string   `json:"description"`
		Currencies       []string `json:"currencies"`
		CustomsCountries []string `json:"customs_countries"`
	}
)

func (r *GetPayWaysRes) Ok() bool {
	return r.Code == 0
}

func (r *GetPayWaysRes) Error() error {
	return fmt.Errorf(r.Msg)
}

type (
	CreatePayWayReq struct {
		OwnerMchID             uint64 `json:"owner_mch_id"`
		ChannelID              uint64 `json:"channel_id"`
		ChannelType            string `json:"channel_type"`
		CurrencySettings       string `json:"currency_settings"`
		CustomsCountrySettings string `json:"customs_country_settings"`
	}
	CreatePayWayRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			PayWayID uint64 `json:"pay_way_id"`
		}
	}
)

func (r *CreatePayWayRes) Ok() bool {
	return r.Code == 0
}

func (r *CreatePayWayRes) Error() error {
	return fmt.Errorf(r.Msg)
}
