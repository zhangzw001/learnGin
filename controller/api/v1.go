package controllerapi

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/dto"
	"github.com/zhangzw001/learnGin/middleware"
)

type Controller struct{}

func Register(router *gin.RouterGroup) {
	api := Controller{}
	router.GET("/product/add", api.Add)
	router.POST("/product/update", api.Update)

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
// @Success 200 {object} middleware.Response{data=dto.ApiOutput} "success"
// @Router /v1/product/add [get]
func (api Controller) Add(c *gin.Context) {
	// 获取 Get 参数
	params := &dto.ApiAddInput{}
	err := params.BindValueParam(c)
	if err != nil {
		middleware.ResponseError(c,10001,err)
		return
	}

	//name := c.Query("name")
	//price := c.DefaultQuery("price", "100")
	//ts := c.Query("ts")
	//sn := c.Query("sn")
	data := dto.ApiOutput{
		Name:  params.Name,
		Price: params.Price,
		Ts:    params.Ts,
		Sn:    params.Sn,
	}
	middleware.ResponseSuccess(c, data)

}




// Add godoc
// @Summary api接口
// @Description api接口
// @Tags v1 update接口
// @ID /v1/product/update
// @Accept  json
// @Produce  json
// @Param body body dto.ApiUpdateInput true "body"
// @Success 200 {object} middleware.Response{data=dto.ApiOutput} "success"
// @Router /v1/product/update [POST]
func (api Controller) Update(c *gin.Context) {

	input := &dto.ApiUpdateInput{}
	err := input.BindValueParam(c)
	if err != nil {
		middleware.ResponseError(c,10001,err)
		return
	}

	//
	//name := c.Query("name")
	//price := c.DefaultQuery("price", "100")
	//ts := c.Query("ts")
	//sn := c.Query("sn")
	//
	data := dto.ApiOutput{
		Name:  input.Name,
		Price: input.Price,
		//Ts:    ts,
		//Sn:    sn,
	}
	middleware.ResponseSuccess(c, data)

}
