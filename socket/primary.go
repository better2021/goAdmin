package socket

import (
	"goAdmin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义 serve 的映射关系
var serveMap = map[string]ServeInterface{
	"GoServer": &GoServe{},
}

func Create() ServeInterface {
	return serveMap["GoServer"]
}

func Start(ctx *gin.Context) {
	Create().RunWs(ctx)
}

func OnlineUserCount() int {
	return Create().GetOnlineUserCount()
}

func OnlineRoomUserCount(roomId int) int {
	return Create().GetOnlineRoomUserCount(roomId)
}

/**
获取当前在线人数
*/
func OnlineCount(ctx *gin.Context) {
	onlineCount := OnlineUserCount()
	ctx.JSON(http.StatusOK, gin.H{
		"OnlineUserCount": onlineCount,
	})
}

/**
创建多个聊天室
*/
func RoomList(ctx *gin.Context) {
	var protocol string
	if ctx.Request.Proto == "HTTP/1.1" {
		protocol = "http://"
	} else {
		protocol = "https://"
	}

	rooms := []map[string]interface{}{
		{"id": 1, "num": OnlineRoomUserCount(1), "title": "聊天室", "imgUrl": protocol + ctx.Request.Host + "/static/" + "f44c6367717440a29056fffc3ba1abdc.jpeg"},
		//{"id": 2, "num": OnlineRoomUserCount(2), "title": "聊天室2", "imgUrl": protocol + ctx.Request.Host + "/static/" + "asd.jfif"},
		//{"id": 3, "num": OnlineRoomUserCount(3), "title": "聊天室3", "imgUrl": protocol + ctx.Request.Host + "/static/" + "qrcode.png"},
		//{"id": 4, "num": OnlineRoomUserCount(4), "title": "聊天室4", "imgUrl": protocol + ctx.Request.Host + "/static/" + "5b0000042eaffa033da6.gif"},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})

}

/**
聊天室历史记录
*/
func Room(ctx *gin.Context) {
	roomId := ctx.Param("room_id")
	msgList := controller.GetLimitMsg(roomId, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"msg_list":       msgList,
		"msg_list_count": len(msgList),
		"room_id":        roomId,
	})
}

/**
私聊历史记录
*/
func PrivateChat(ctx *gin.Context) {
	roomId := ctx.Query("room_id")
	uid := ctx.Query("uid")
	toUid := ctx.Query("to_uid")
	msgList := controller.GetLimitPrivateMsg(uid, toUid, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"msg_List": msgList,
		"room_id":  roomId,
	})
}
