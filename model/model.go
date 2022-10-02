package model

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Err(msg string) Response {
	return Response{Code: -1, Msg: msg}
}

func Suc(data interface{}) Response {
	return Response{Code: 0, Data: data}
}
