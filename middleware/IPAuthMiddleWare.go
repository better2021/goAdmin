package middleware

import (
	"fmt"
	"goAdmin/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IP黑名单
func IPAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now() // 当前时间
		var ips = controller.Ips()
		fmt.Println(ips, "---")

		var ipList = []string{} // IP黑名单
		for _, v := range ips {
			ipList = append(ipList, v.Ip)
		}
		elapsed := time.Since(start)
		fmt.Println("该函数执行完成耗时：", elapsed)
		flag := true /*如果要改为ip白名单把flag:=true 改为false 调换即可*/
		clientIp := ctx.ClientIP()
		for _, host := range ipList {
			if clientIp == host {
				flag = false
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  fmt.Sprintf("%s 禁止访问，此ip地址已在黑名单中, 如要访问请联系管理员", clientIp),
			})
			ctx.Abort() // 中断请求
			return
		}
		ctx.Next() // 继续执行后面的代码
	}
}
