package sign

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzw001/learnGin/dto"
	"github.com/zhangzw001/learnGin/middleware"
	"github.com/zhangzw001/learnGin/public"
	"log"
	"strconv"
)

type Controller struct{}

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
	ts := strconv.FormatInt(public.GetTimeUnix(), 10)

	params := &dto.ApiUpdateInput{}
	err := params.BindValueParam(c)
	if err != nil {
		middleware.ResponseError(c, 10001, err)
	}

	paramsSign := c.Request.Form
	paramsSign["ts"] = []string{ts}

	data := dto.ApiAdd{
		Name:  params.Name,
		Price: params.Price,
		Ts:    ts,
		Sn:    public.CreateSign(paramsSign),
	}
	log.Println(data)
	middleware.ResponseSuccess(c, data)

}
