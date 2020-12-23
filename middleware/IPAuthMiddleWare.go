package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/common"
	"goAdmin/model"
	"net/http"
)

// IP白名单
func IPAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ips []model.IpWhite
		db := common.InitDB()
		db.Find(&ips)
		defer db.Close()
		// fmt.Println(ips,"ips")

		var ipList = []string {"127.0.0.1"} // IP白名单,首先有默认ip为 127.0.0.1
		for _,v := range ips{
			ipList = append(ipList, v.Ip)
		}

		flag := false	/*如果要改为ip黑名单把flag:=false 改为true 调换即可*/
		clientIp := ctx.ClientIP()
		for _, host := range ipList {
			if clientIp == host {
				flag = true
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  fmt.Sprintf("%s 禁止访问，此ip地址不在白名单中", clientIp),
			})
			ctx.Abort() // 中断请求
			return
		}
		ctx.Next() // 继续执行后面的代码
	}
}
