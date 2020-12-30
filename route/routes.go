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
	r.POST("/api/auth/register",controller.Register)

	r.Use(middleware.AuthMiddleware(),middleware.RecoveryMiddleware()) // 使用跨域中间件 和 cover()及ip白名单 中间件
	v1 := r.Group("/api")
	{
		v1.GET("/users",controller.UserList)
		v1.PUT("/users/:id",controller.UserUpdate)
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

		v1.GET("/musics",controller.MusicList)
		v1.POST("/musics",controller.MusicCreate)
		v1.PUT("/musics/:id",controller.MusicUpdate)
		v1.DELETE("/musics/:id",controller.MusicDelete)

		v1.GET("/notes",controller.NoteList)
		v1.POST("/notes",controller.NoteCreate)
		v1.DELETE("/notes/:id",controller.NoteDelete)

		v1.GET("/ipWhite",controller.IpList)
		v1.POST("/ipWhite",controller.IpsCreate)
		v1.DELETE("/ipWhite/:id",controller.IpsDelete)

		v1.POST("/upload",controller.UploadFile) 	// 单文件上传
		v1.POST("/uploads",controller.UploadFiles)	// 多文件上传
	}

	// 未知路由处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"code":http.StatusNotFound,
			"msg":"路由地址错误或请求方法错误",
		})
	})

	return r
}
