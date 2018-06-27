package entity

import "github.com/everywan/go-web/config"

type Result_web struct {
	Code int
	Msg  interface{}
	Data interface{}
}

func (r *Result_web) TransCode(statusCode *config.Code) {
	r.Code = statusCode.Code
	r.Msg = statusCode.Msg
}
