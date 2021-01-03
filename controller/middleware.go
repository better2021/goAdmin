package controller

import (
	"fmt"
	"goAdmin/model"
	"goAdmin/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 查找ip白名单列表并返回
func Ips() []model.IpWhite {
	var ips []model.IpWhite
	db.Find(&ips)
	return ips
}

// 根据userId查找到对应好的user信息
func FindUser(userId uint) model.User {
	var user model.User
	db.First(&user, userId)
	return user
}

func FindApi(c *gin.Context) {
	host := c.Request.Host
	fmt.Println(host, "host")
	err := qrcode.WriteFile(host+"/swagger/index.html", qrcode.Medium, 256, "./uploadFiles/qrcode.png")
	if err != nil {
		fmt.Println(err)
	}

	var visit model.Visit
	db.Find(&visit) // 等价于 	db.Raw("SELECT * FROM visits").Scan(&visit)
	visit.VisitNum++
	db.Save(&visit)

	//ip := util.GetClientIp()
	//serverIp := util.GetServerIP()
	//RemoteIP := util.RemoteIP(c.Request)

	c.JSON(http.StatusOK, gin.H{
		"message":  "hello golang",
		"time":     time.Now().Format("2006-01-02 15:04:05"),
		"week":     util.Getweek(),
		"qrcode":   "http://" + host + "/static/qrcode.png",
		"visitNum": visit.VisitNum,
		"ip":       c.ClientIP(),
		//"serverIp": serverIp,
		//"RemoteIP": RemoteIP,
	})
}
