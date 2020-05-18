package model

import "net/http"

type FeedBack struct {
	Dist http.ResponseWriter `json:"-"` //不进行序列化
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data interface{}         `json:"data,omitempty"`
}
