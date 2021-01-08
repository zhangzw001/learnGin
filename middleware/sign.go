package middleware

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	config "github.com/zhangzw001/learnGin/config/dev"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// 获取当前时间
func GetTimeUnix() int64  {
	return time.Now().Unix()
}

// MD5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return fmt.Sprintf("%x",s.Sum(nil))
	//return hex.EncodeToString(s.Sum(nil ))
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(config.AppName+ config.AppSecret))
	return sign
}

func SignMiddleware() gin.HandlerFunc {
	return VerifySign
}
// 验证签名
func VerifySign(c *gin.Context) {
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
	if ts > GetTimeUnix() || GetTimeUnix() - ts >= exp {
		ResponseError(c,10001,errors.New("Ts Error"))
		return
	}

	// 验证签名
	if sn == "" || sn != CreateSign(req) {
		ResponseError(c,10001,errors.New("Sn Error"))
		return
	}
	//
	c.Next()
}



func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(GetTimeUnix(), 10)
	res := map[string]interface{}{}
	name := c.Query("name")
	price := c.Query("price")
	params := url.Values{
		"name"  : []string{name},
		"price" : []string{price},
		"ts"    : []string{ts},
	}
	res["sn"] = CreateSign(params)
	res["ts"] = ts
	ResponseSuccess(c,res)
}
