package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhangzw001/learnGin/config"
	"reflect"
	"regexp"
)

// 设置验证器
//func ValidatorMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 创建实例
//		val := validator.New()
//		val.RegisterValidation("test_tag",func(fl validator.FieldLevel) bool {
//			return  fl.Field().String() != "testName"
//		})
//
//		c.Set(config.ValidatorKey,val)
//		c.Next()
//
//	}
//}



// 设置验证器+翻译器
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建验证器实例
		val := validator.New()

		// 设置支持的语言
		en := en.New()
		zh := zh.New()
		// 创建翻译器, 支持zh,en两种
		uni := ut.New(zh, zh,en )
		// 根据参数取翻译器实例
		locale := c.DefaultQuery("locale","zh")
		// 获取默认的f.fallback
		trans ,_ := uni.GetTranslator(locale)


		// 翻译器注册到验证器
		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(val,trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			_ = zh_translations.RegisterDefaultTranslations(val,trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			// 自定义的验证器
			val.RegisterValidation("test_tag",func(fl validator.FieldLevel) bool {
				matched, _ := regexp.MatchString(`^[_0-9a-zA-Z]{6,24}$`,fl.Field().String())
				return matched
			})

			// 自定义的翻译器
			val.RegisterTranslation("test_tag",trans,func(ut ut.Translator) error {
				return ut.Add("test_tag","{0} 填写不正确,仅允许数字字母下划线,长度大于6小于24",true  )
			},func(ut ut.Translator, fe validator.FieldError) string {
				t , _ := ut.T("test_tag",fe.Field())
				return t
			})
			break
		}

		// 设置翻译器和验证器
		c.Set(config.TranslatorKey,trans)
		c.Set(config.ValidatorKey,val)
		c.Next()

	}
}
