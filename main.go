package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goAdmin/common"
	_ "goAdmin/docs" // 注意这个一定要引入自己的docs
	"goAdmin/model"
	"goAdmin/route"
	"goAdmin/util"
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

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	url := ginSwagger.URL("80/swagger.doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/api", func(c *gin.Context) {
		host := c.Request.Host
		fmt.Println(host,"host")
		err := qrcode.WriteFile(host+"/swagger/index.html", qrcode.Medium, 256, "./uploadFiles/qrcode.png")
		if err != nil {
			fmt.Println(err)
		}

		var visit model.Visit
		db.Find(&visit) // 等价于 	db.Raw("SELECT * FROM visits").Scan(&visit)
		visit.VisitNum++
		db.Save(&visit)

		ip := util.GetClientIp()
		serverIp := util.GetServerIP()
		RemoteIP := util.RemoteIP(c.Request)

		c.JSON(http.StatusOK, gin.H{
			"message":  "hello golang",
			"time":     time.Now().Format("2006-01-02 15:04:05"),
			"week":     util.Getweek(),
			"qrcode":   "http://" + host + "/static/qrcode.png",
			"visitNum": visit.VisitNum,
			"ip":       ip,
			"serverIp": serverIp,
			"RemoteIP": RemoteIP,
		})
	})

	r = route.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		fmt.Println(r.Run(":" + port))
	}

	r.Run(port)
}
