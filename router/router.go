package router

import (
	"application/app/controller"
	"application/router/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, contextPath string) {
	// 跨域
	r.Use(middleware.Cors())
	e := r.Group(fmt.Sprintf("/%s", contextPath))
	{
		// 区域信息
		e.POST("/index", controller.Index()) // 获取所有树编码
	}
}
