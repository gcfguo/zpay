package model

type (
	GetAppsReq struct {
	}
	GetAppsRes struct {
		List  []*AppInfo `json:"list"`
		Total int        `json:"total"`
	}
	AppInfo struct {
		AppId        string `json:"app_id"         `
		AppSecret    string `json:"app_secret"     `
		CreatorMchId uint64 `json:"creator_mch_id" `
		OwnerMchId   uint64 `json:"owner_mch_id"   `
		Description  string `json:"description"    `
		CreateTime   string `json:"create_time"    `
		Settings     struct {
			PayWaySettings      string `json:"pay_way_settings"`
			PayModeSettings     int    `json:"pay_mode_settings"`
			IpBlacklistSettings string `json:"ip_blacklist_settings"`
			OrderSceneSettings  string `json:"order_scene_settings"`
			PaidConfirmType     string `json:"paid_confirm_type"`
			RefundedConfirmType string `json:"refunded_confirm_type"`
		} `json:"settings"`
	}
)

type (
	CreateAppReq struct {
		OwnerMchID  uint64 `json:"owner_mch_id"`
		Description string `json:"description"`
		PaymentMode int    `json:"payment_mode"`
	}
	CreateAppRes struct {
	}
)

type (
	AppSettingsReq struct {
		AppID               string   `json:"app_id"`
		PayModeSettings     int      `json:"pay_mode_settings"`
		IPBlacklistSettings []string `json:"ip_blacklist_settings"`
		OrderSceneSettings  []string `json:"order_scene_settings"`
		PayWaySettings      []string `json:"pay_way_settings"`
	}
	AppSettingsRes struct {
	}
)
