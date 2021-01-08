package sign

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/dto"
	"github.com/zhangzw001/learnGin/middleware"
	"net/url"
	"strconv"
)

type Controller struct {}
func RegisterSign(router *gin.RouterGroup) {
	sign := Controller{}
	router.GET("/create", sign.Create)
}



// Add godoc
// @Summary sign接口
// @Description sign接口
// @Tags sign 创建接口
// @ID /sn/create
// @Accept  json
// @Produce  json
// @Param name query string true "名称"
// @Param price query string true "价格"
// @Success 200 {object} middleware.Response{data=dto.ApiAdd} "success"
// @Router /sn/create [get]
func (sign Controller) Create(c *gin.Context) {
	ts := strconv.FormatInt(middleware.GetTimeUnix(), 10)
	name := c.Query("name")
	price := c.Query("price")
	params := url.Values{
		"name"  : []string{name},
		"price" : []string{price},
		"ts"    : []string{ts},
	}

	data := dto.ApiAdd{
		Name:  name,
		Price: price,
		Ts:    ts,
		Sn:    middleware.CreateSign(params),
	}

	middleware.ResponseSuccess(c,data)

}
