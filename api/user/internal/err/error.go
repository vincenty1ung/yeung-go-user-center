package err

import "net/http"

// 返回的结构体，json格式的body
type Message struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

// 定义错误处理函数
func ErrorHandler(err error) (int, interface{}) {
	return http.StatusConflict, Message{
		Code: -1,
		Desc: err.Error(),
	}
}
