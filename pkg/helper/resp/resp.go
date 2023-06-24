package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/enum"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: code, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}

func HandleParamsErr(ctx *gin.Context) {
	resp := response{Code: enum.Error, Message: "参数错误", Data: nil}
	ctx.JSON(http.StatusBadRequest, resp)
}
