package ctl

import (
	"AlittleRequire/pkg/e"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
}

func RespError(err error, data interface{}, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}
