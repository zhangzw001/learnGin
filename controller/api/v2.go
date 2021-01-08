package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/dto"
	"github.com/zhangzw001/learnGin/middleware"
)

type Controller2 struct {}
func Register2(router *gin.RouterGroup) {
	api := Controller{}
	router.GET("/product/add", api.Add)
}

// Add godoc
// @Summary api接口
// @Description api接口
// @Tags v2 add接口
// @ID /v2/product/add
// @Accept  json
// @Produce  json
// @Param name query string true "名称"
// @Param price query string true "价格"
// @Success 200 {object} middleware.Response{data=dto.ApiAdd} "success"
// @Router /v2/product/add [get]
func (api Controller2) Add(c *gin.Context) {
	// 获取 Get 参数
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")

	data := dto.ApiAdd{
		Name:  name,
		Price: price,
	}
	middleware.ResponseSuccess(c,data)

}


