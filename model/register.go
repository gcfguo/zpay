package model

import "fmt"

type (
	RegisterReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	RegisterRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func (r *RegisterRes) Ok() bool {
	return r.Code == 0
}

func (r *RegisterRes) Error() error {
	return fmt.Errorf(r.Msg)
}
