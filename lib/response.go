package lib

import "encoding/json"

type JsonResponse struct {
	Code    int         `json:"code"`    //错误码
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    //返回数据
}

func (jr *JsonResponse) String() string {
	val, _ := json.Marshal(jr)
	return string(val)
}
