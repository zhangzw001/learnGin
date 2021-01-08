package public

import (
	"crypto/md5"
	"fmt"
	config "github.com/zhangzw001/learnGin/config/dev"
	"net/url"
	"sort"
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
