package controllers

const (
	SUCCESS_CODE      = 0
	PARAM_ERROR_CODE  = 1000
	SERVER_ERROR_CODE = 1001
)

type Response struct {
	Data    interface{} `json:"data"`
	ErrMsg  string      `json:"errmsg"`
	ErrCode int         `json:"errcode"` //0:正确
}
