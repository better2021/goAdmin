package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/controller"
	"net/http"
	"time"
)

// IP白名单
func IPAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now() // 当前时间
		var ips = controller.Ips()
		fmt.Println(ips,"---")

		var ipList = []string {"127.0.0.1","172.17.0.1"} // IP白名单,首先有默认ip为 127.0.0.1
		for _,v := range ips{
			ipList = append(ipList, v.Ip)
		}
		elapsed := time.Since(start)
		fmt.Println("该函数执行完成耗时：", elapsed)
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
