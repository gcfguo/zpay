package model

type (
	CreatePayPalChannelReq struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		IsSandBox    bool   `json:"is_sand_box"`
	}
	CreatePayPalChannelRes struct {
		ChannelID uint64 `json:"channel_id"`
	}
)

type (
	CreateWechatSubChannelReq struct {
		ZPayMchID   uint64 `json:"zpay_mch_id"`
		ChannelID   uint64 `json:"channel_id"`
		WechatMchID string `json:"wechat_mch_id"`
	}
	CreateWechatSubChannelRes struct {
		ChannelID uint64 `json:"channel_id"`
	}
)
