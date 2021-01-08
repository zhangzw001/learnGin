package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	config "github.com/zhangzw001/learnGin/config/dev"
	"github.com/zhangzw001/learnGin/public"
	"net/url"
	"strconv"
)



func SignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var method = c.Request.Method
		var ts int64
		var sn string
		var req url.Values

		if method == "GET" {
			req = c.Request.URL.Query()
			sn = c.Query("sn")
			ts, _  = strconv.ParseInt(c.Query("ts"), 10, 64)

		} else if method == "POST" {
			c.Request.ParseForm()
			req = c.Request.PostForm
			sn = c.PostForm("sn")
			ts, _  = strconv.ParseInt(c.PostForm("ts"), 10, 64)
		} else {
			ResponseError(c,10000,errors.New("Illegal requests"))
			return
		}

		exp, _ := strconv.ParseInt(config.ApiExpiry, 10, 64)

		// 验证过期时间
		if ts > public.GetTimeUnix() || public.GetTimeUnix() - ts >= exp {
			ResponseError(c,10001,errors.New("Ts Error"))
			return
		}

		// 验证签名
		if sn == "" || sn != public.CreateSign(req) {
			ResponseError(c,10001,errors.New("Sn Error"))
			return
		}
		//
		c.Next()
	}
}



//func SignDemo(c *gin.Context) {
//	ts := strconv.FormatInt(GetTimeUnix(), 10)
//	res := map[string]interface{}{}
//	name := c.Query("name")
//	price := c.Query("price")
//	params := url.Values{
//		"name"  : []string{name},
//		"price" : []string{price},
//		"ts"    : []string{ts},
//	}
//	res["sn"] = CreateSign(params)
//	res["ts"] = ts
//	ResponseSuccess(c,res)
//}
