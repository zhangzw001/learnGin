package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	config "github.com/zhangzw001/learnGin/config/dev"
	"strings"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	//
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e :=range errs {
			sliceErrs = append(sliceErrs, e.Error())
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
	validator , ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validator, nil
}
