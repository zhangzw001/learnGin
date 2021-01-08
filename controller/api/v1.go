package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/dto"
	"github.com/zhangzw001/learnGin/middleware"
)

type Controller struct{}

func Register(router *gin.RouterGroup) {
	api := Controller{}
	router.GET("/product/add", api.Add)
}

// Add godoc
// @Summary api接口
// @Description api接口
// @Tags v1 add接口
// @ID /v1/product/add
// @Accept  json
// @Produce  json
// @Param name query string true "名称"
// @Param price query string true "价格"
// @Param ts query string true "过期时间"
// @Param sn query string true "签名"
// @Success 200 {object} middleware.Response{data=dto.ApiAdd} "success"
// @Router /v1/product/add [get]
func (api Controller) Add(c *gin.Context) {
	// 获取 Get 参数
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")
	ts := c.Query("ts")
	sn := c.Query("sn")

	data := dto.ApiAdd{
		Name:  name,
		Price: price,
		Ts:    ts,
		Sn:    sn,
	}
	middleware.ResponseSuccess(c, data)

}
