package route

import (
	"github.com/gin-gonic/gin"
	"goAdmin/controller"
	"goAdmin/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine{
	r.Use(middleware.CorsMiddleware(),middleware.RecoveryMiddleware()) // 使用跨域中间件 和 cover() 中间件
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users",controller.UserList)
		v1.POST("/auth/register",controller.Register)
		v1.POST("/auth/login",controller.Login)
		v1.POST("/auth/info",middleware.AuthMiddleware(),controller.Info)
	}

	return r
}
