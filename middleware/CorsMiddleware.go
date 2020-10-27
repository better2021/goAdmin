package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// * 表示所有的域名都可以访问
		ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
		// 设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age","86400")
		// 设置可以访问的请求方法 POST，GET等，* 表示允许所有的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods","*")
		// 允许请求带的header头信息
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","*")
		// 响应头表示是否可以将对请求的响应暴露给页面。返回true则可以，其他值均不可以
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials","true")

		if ctx.Request.Method == http.MethodOptions{
			ctx.AbortWithStatus(200)
		}else {
			ctx.Next()
		}
	}
}
