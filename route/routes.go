package route

import (
	"github.com/gin-gonic/gin"
	"goAdmin/controller"
	"goAdmin/middleware"
	"net/http"
)

func CollectRouter(r *gin.Engine) *gin.Engine{
	// 静态资源文件
	r.StaticFS("/static",http.Dir("uploadFiles"))
	r.Use(middleware.CorsMiddleware(),middleware.IPAuthMiddleWare())
	r.GET("/api/getCode",controller.GenerateCaptchaHandler)
	r.POST("/api/auth/login",controller.Login)

	r.Use(middleware.CorsMiddleware(),middleware.AuthMiddleware(),middleware.RecoveryMiddleware(),middleware.IPAuthMiddleWare()) // 使用跨域中间件 和 cover()及ip白名单 中间件
	v1 := r.Group("/api")
	{
		v1.GET("/users",controller.UserList)
		v1.POST("/auth/register",controller.Register)
		v1.PUT("/users/:id",controller.ChangePassword)
		v1.GET("/auth/info",controller.Info)
		v1.DELETE("/users/:id",controller.UserDelete)

		v1.GET("/films",controller.FilmList)
		v1.POST("/films",controller.FilmCreate)
		v1.PUT("/films/:id",controller.FilmUpdate)
		v1.DELETE("/films/:id",controller.FilmDelete)

		v1.GET("/books",controller.BookList)
		v1.POST("/books",controller.BookCreate)
		v1.PUT("/books/:id",controller.BookUpdate)
		v1.DELETE("/books/:id",controller.BookDelete)

		v1.POST("/upload",controller.UploadFile)
	}

	// 未知路由处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"code":http.StatusNotFound,
			"msg-zh":"路由地址错误或请求方法错误",
			"msg-en":"Not router or Method",
		})
	})

	return r
}
