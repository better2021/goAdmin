package socket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// http升级websocket协议的配置
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(ctx *gin.Context) {
	// 升级get请求为websocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// 读取ws中的数据,mt是消息类型，message是具体的消息
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(mt, string(message))
		if string(message) == "ping" {
			message = []byte("pong")
		}
		// 写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
