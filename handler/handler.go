package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 封装通用响应体
type Response struct {
	Code int         `json: "code"`
	Msg  string      `json: "msg"`
	Data interface{} `json: "data"`
}

// SendResponse 响应
func SendResponse(c *gin.Context, err error, data interface{}) {
	code := 0
	msg := "ok"

	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
