package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/zhangzw001/learnGin/config/dev"
	"github.com/zhangzw001/learnGin/routers"
)

func main() {

	gin.SetMode(gin.DebugMode)
	e := gin.Default()
	routers.InitRouter(e)
	e.Run(config.PORT)
}
