package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/common"
	"goAdmin/controller"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		// fmt.Print(tokenString,"token")
		// 验证token格式，token要Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":http.StatusUnauthorized,"msg":"token未携带或格式不对"})
			ctx.Abort() // 中断请求
			return
		}

		tokenString = tokenString[7:] // 截取7位以后的

		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			fmt.Println(err,"err",token.Valid)
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":http.StatusUnauthorized,"msg":"token已过期，请重新登录"})
			ctx.Abort()
			return
		}

		start := time.Now()
		// 验证通过获取claim 中的userId
		userId := claims.UserId
		user := controller.FindUser(userId)

		elapsed := time.Since(start)
		fmt.Println("耗时",elapsed)

		// 用户不存在
		if user.ID == 0{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":http.StatusUnauthorized,"msg":"没有权限哟"})
			ctx.Abort()
			return
		}

		// 用户存在，将user 的信息写入上下文
		ctx.Set("user",user)
		ctx.Next() // 继续执行后面的代码
	}
}
