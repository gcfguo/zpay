package model

type (
	RegisterReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	RegisterRes struct {
	}
)
