package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/public"
)

type ApiOutput struct {
	Name string
	Price int
	Ts string
	Sn string
}

type ApiAddInput struct {
	Name string `json:"name" form:"name" comment:"名称" example:"abc" validate:"required,test_tag,max=24,min=6"`
	Price int `json:"price" form:"price" comment:"价格" example:"abc"  validate:"required,min=0"`
	Ts string `json:"ts" form:"ts" comment:"有效期" example:"" validate:"required,min=1"`
	Sn string `json:"sn" form:"sn" comment:"签名" example:"" validate:"required,min=32,max=32"`
}

// 绑定结构体
func (param *ApiAddInput) BindValueParam(c *gin.Context) error  {
	return public.DefaultGetValidParams(c,param)
}



type ApiUpdateInput struct {
	Name string		`json:"name" form:"name" comment:"名称" example:"abc" validate:"required,test_tag,max=24,min=6"`
	Price int	`json:"price" form:"price" comment:"价格" example:"10" validate:"required,min=0"`
}

// 绑定结构体
func (param *ApiUpdateInput) BindValueParam(c *gin.Context) error  {
	return public.DefaultGetValidParams(c,param)
}
