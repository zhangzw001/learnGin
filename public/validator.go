package public

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zhangzw001/learnGin/config"
	"strings"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	// 获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	// 获取翻译器
	//trans := new(ut.Translator)
	trans, err := GetTranslate(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e :=range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

// 获取验证器
func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(config.ValidatorKey)
	if !ok {
		return nil, errors.New("未设置验证器")
	}
	validate , ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validate, nil
}

// 获取翻译器
func GetTranslate(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get(config.TranslatorKey)
	if !ok {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}
