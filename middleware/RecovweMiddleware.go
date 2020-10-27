package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/* panic恢复中间件*/
func RecoveryMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover();err!=nil{
				ctx.JSON(http.StatusOK,gin.H{"code":400,"msg":err})
			}
		}()

		ctx.Next()
	}
}