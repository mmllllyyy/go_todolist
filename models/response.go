package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 200
	ERROR   = 500
)

type ResponseVo struct {
	Code     int32       `json:"code" form:"code"`
	Msg      string      `json:"msg" form:"msg"`
	DataName string      `json:"dataName" form:"dataName"`
	Data     interface{} `json:"data,omitempty" form:"data"`
}

func SuccessResponse(ctx *gin.Context, dataName string, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseVo{Code: 0, Msg: "success", DataName: dataName, Data: data})
}

func FailResponse(ctx *gin.Context, msg string, dataName string, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseVo{Code: 1, Msg: msg, DataName: dataName, Data: data})
}

func SuccessVo(dataName string, data interface{}) ResponseVo {
	return ResponseVo{Code: 0, Msg: "success", DataName: dataName, Data: data}
}

func FailVo(msg string, dataName string, data interface{}) ResponseVo {
	return ResponseVo{Code: 1, Msg: msg, DataName: dataName, Data: data}
}
