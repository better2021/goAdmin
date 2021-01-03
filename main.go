package main

import (
	"fmt"
	"goAdmin/controller"
	_ "goAdmin/docs" // 注意这个一定要引入自己的docs
	"goAdmin/middleware"
	"goAdmin/route"
	"goAdmin/socket"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Gin API
// @version 2.0
// @description An example of gin
// @termsOfService 运行地址：http://localhost/swagger/index.html
// @license.name MIT //localhost:80

func main() {
	f, _ := os.Create("gin.log")                          // 创建gin.log日志文件
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout) // 错误信息写入gin.log日志文件

	var db *gorm.DB
	defer db.Close()

	r := gin.Default()

	url := ginSwagger.URL("80/swagger.doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(middleware.CorsMiddleware())
	r.GET("/api", controller.FindApi)
	// r.GET("/ws", socket.WsHandler)

	r.GET("/ws", socket.Start)
	r.GET("/onlineCount", socket.OnlineCount)
	r.GET("/roomList", socket.RoomList)
	r.GET("/room/:room_id", socket.Room)
	r.GET("/private-chart", socket.PrivateChat)

	r = route.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		fmt.Println(r.Run(":" + port))
	}

	r.Run(port)
}
