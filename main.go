package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goAdmin/common"
	"goAdmin/route"
	"net/http"
	"os"
	"time"
)

func main()  {
	InitConfig()
	db := common.InitDB() // 初始化数据库
	defer db.Close()

	r := gin.Default()
	r = route.CollectRouter(r)
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
