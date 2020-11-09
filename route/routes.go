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
		v1.PUT("/users/:id",controller.ChangePassword)
		v1.POST("/auth/login",controller.Login)
		v1.GET("/auth/info",middleware.AuthMiddleware(),controller.Info)
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

	return r
}
