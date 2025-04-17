package qo

// AccountRequest 请求结构体
type AccountRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password,omitempty"`
}
