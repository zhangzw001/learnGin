package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zhangzw001/learnGin/controller/api"
	"github.com/zhangzw001/learnGin/controller/sign"
	_ "github.com/zhangzw001/learnGin/docs"
	"github.com/zhangzw001/learnGin/middleware"
)

func InitRouter(e *gin.Engine) {
	//e.GET("/sn",middleware.SignDemo)
	e.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	// sign
	Sign := e.Group("/sn")
	{
		sign.RegisterSign(Sign)
	}


	// v1版本
	GroupV1 := e.Group("/v1")
	GroupV1.Use(middleware.SignMiddleware(),middleware.LoggerToFile())
	{
		api.Register(GroupV1)
	}

	// v2版本
	GroupV2 := e.Group("/v2")
	{
		api.Register2(GroupV2)
	}


}


