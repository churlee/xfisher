package dto

import (
	"github.com/gin-gonic/gin"
	"lilith/common"
)

type Response struct {
	C *gin.Context
}

//返回结构体
type RestfulDto struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//成功返回调用方法
func (r Response) Success(httpCode int, data interface{}) {
	r.C.JSON(httpCode, RestfulDto{
		Data: data,
	})
}

//失败返回调用方法
func (r Response) Error(httpCode int, code string) {
	r.C.JSON(httpCode, RestfulDto{
		Code: code,
		Msg:  common.GetMsg(code),
	})
}
