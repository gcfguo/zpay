package model

import "fmt"

type (
	CreatePayPalChannelReq struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		IsSandBox    bool   `json:"is_sand_box"`
	}
	CreatePayPalChannelRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			ChannelID uint64 `json:"channel_id"`
		}
	}
)

func (r *CreatePayPalChannelRes) Ok() bool {
	return r.Code == 0
}

func (r *CreatePayPalChannelRes) Error() error {
	return fmt.Errorf(r.Msg)
}
