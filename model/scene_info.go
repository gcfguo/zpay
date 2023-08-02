package model

type SceneInfo struct {
	//下单场景类型
	SceneType string `json:"scene_type"`
	//DeviceID 商户端设备号
	DeviceID string `json:"device_id"`
	//PayerClientIP 用户终端IP
	PayerClientIP string `json:"payer_client_ip"`
	//H5Type H5支付场景类型
	H5Type string `json:"h5_type"`
	//H5AppName 应用名称
	H5AppName string `json:"h5_app_name"`
	//H5AppURL 网站URL
	H5AppURL string `json:"h5_app_url"`
}
