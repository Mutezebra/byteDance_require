package api

import (
	"AlittleRequire/pkg/ctl"
	"AlittleRequire/pkg/e"
	"encoding/json"
)

func ErrorResponse(err error, code int) *ctl.Response { // 用来处理错误信息
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "json类型不匹配")
	}
	return ctl.RespError(err, e.GetMsg(code))
}
