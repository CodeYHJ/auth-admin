package util

type ResponseModal struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func CreateRes() *ResponseModal {
	r := &ResponseModal{Code: 200, Data: nil, Msg: ""}
	return r
}

func CreateFailRes(code int, msg string) *ResponseModal {
	r := &ResponseModal{Code: code, Msg: msg, Data: nil}
	return r
}
