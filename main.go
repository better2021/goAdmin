package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goAdmin/common"
	_ "goAdmin/docs" // 注意这个一定要引入自己的docs
	"goAdmin/route"
	"io"
	"net/http"
	"os"
	"time"
)

// @title Golang Gin API
// @version 2.0
// @description An example of gin
// @termsOfService 运行地址：http://localhost/swagger/index.html
// @license.name MIT //localhost:80
func main() {
	f, _ := os.Create("gin.log")               // 创建gin.log日志文件
	gin.DefaultErrorWriter = io.MultiWriter(f) // 错误信息写入gin.log日志文件

	var db = common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = route.CollectRouter(r)

	url := ginSwagger.URL("80/swagger.doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello golang",
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run(port)
}