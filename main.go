package main

import (
	"fmt"
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
func main()  {
	f,_ := os.Create("gin.log") // 创建gin.log日志文件
	gin.DefaultErrorWriter = io.MultiWriter(f) // 错误信息写入gin.log日志文件

	InitConfig()
	db := common.InitDB() // 初始化数据库
	defer db.Close()

	r := gin.Default()
	r = route.CollectRouter(r)

	url := ginSwagger.URL("80/swagger.doc.json")
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler,url))

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"hello golang",
			"time":time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	port := viper.GetString("server.port")
	if port != ""{
		panic(r.Run(":" + port))
	}
	r.Run(port)
}

func InitConfig(){
	// 获取当前的工作目录
	workDir,_:= os.Getwd()
	fmt.Println("当前文件的路劲:" + workDir)
	// 设置要读取的文件名
	viper.SetConfigName("application")
	// 设置要读取的文件的类型
	viper.SetConfigType("yml")
	// 添加读取文件的路劲
	viper.AddConfigPath(workDir + "/config")
	// 读取文件配置
	err := viper.ReadInConfig()
	if err !=nil{
		fmt.Println(err,"---")
	}
}
