package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	}
}

func Failure(msg string) *Response {
	return &Response{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}
