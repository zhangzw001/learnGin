package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode ResponseCode = 200 + iota
	InternalServerError ResponseCode = 500
	CustomizeCode           = 1000
	)

type Response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{
		Code: SuccessCode,
		Msg:  "ok",
		Data: data,
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func ResponseError(c *gin.Context, code ResponseCode,err error) {
	resp := &Response{
		Code: code,
		Msg:  err.Error(),
		Data: "",
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200,err)
}
