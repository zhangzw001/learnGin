package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	config "github.com/zhangzw001/learnGin/config/dev"
)

// 设置验证器
func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建实例
		val := validator.New()
		val.RegisterValidation("testTag",func(fl validator.FieldLevel) bool {
			return  fl.Field().String() != "testName"
		})

		c.Set(config.ValidatorKey,val)
		c.Next()

	}
}
