package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Middleware(engine *gin.Engine) {
}

//// 处理跨域请求,支持options访问
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		if method == "OPTIONS" {
//			c.Header("Access-Control-Allow-Origin", "*")
//			c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken,Token")
//			c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
//			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//			c.Header("Access-Control-Allow-Credentials", "true")
//		}
//		// 放行所有OPTIONS方法
//		if method == "GET" {
//			// 处理GET请求
//			c.Header("Content-Type", "text/plain; charset=utf-8")
//			c.Header("Access-Control-Allow-Origin", "*") // 允许所有源
//		}
//		c.Next()
//	}
//}

//func Cors() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		// 允许 Origin 字段中的域发送请求
//		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
//		// 设置预验请求有效期为 86400 秒
//		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
//		// 设置允许请求的方法
//		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
//
//		// context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		// context.Header("Access-Control-Allow-Credentials", "true")
//		// OPTIONS请求返回200
//		if context.Request.Method == "OPTIONS" {
//			context.AbortWithStatus(http.StatusNoContent)
//		}
//	}
//}

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type", "Access-Token", "client-id"},
		MaxAge:          6 * time.Hour,
	}

	return cors.New(c)
}
