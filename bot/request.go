package bot

import (
	"github.com/gin-gonic/gin"
)

func (b *Bot) request(api string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors())     //跨域
	r.Use(recovery()) //错误处理
	r.Use(console())
	//*******************************路由设置*******************************
	v1 := r.Group(api)
	v1.GET("/exp", b.verify)
	v1.POST("/exp", b.message)
	//********************************************************************
	r.NoRoute(func(c *gin.Context) {
		serverError(c, "无效接口")
	})
	return r
}
