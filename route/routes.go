package route

import (
	"github.com/gin-gonic/gin"
	"goAdmin/controller"
	"goAdmin/middleware"
	"net/http"
	"time"
)

func CollectRouter(r *gin.Engine) *gin.Engine{
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
			"time":time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.Use(middleware.CorsMiddleware(),middleware.RecoveryMiddleware()) // 使用跨域中间件 和 cover() 中间件
		v1.POST("/auth/register",controller.Register)
		v1.POST("/auth/login",controller.Login)
		v1.POST("/auth/info",middleware.AuthMiddleware(),controller.Info)
	}

	return r
}
