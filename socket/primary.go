package socket

import (
	"goAdmin/controller"
	"net/http"
	"strconv"

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
	userInfo := controller.Info(ctx)
	rooms := []map[string]interface{}{
		{"id": 1, "num": OnlineRoomUserCount(1)},
		{"id": 2, "num": OnlineRoomUserCount(2)},
		{"id": 3, "num": OnlineRoomUserCount(3)},
		{"id": 4, "num": OnlineRoomUserCount(4)},
		{"id": 5, "num": OnlineRoomUserCount(5)},
		{"id": 6, "num": OnlineRoomUserCount(6)},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rooms":     rooms,
		"user_info": userInfo,
	})

}

/**
聊天室
*/
func Room(ctx *gin.Context) {
	roomId := ctx.Param("room_id")

	userInfo := controller.Info(ctx)
	msgList := controller.GetLimitMsg(roomId, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"user_info":      userInfo,
		"msg_list":       msgList,
		"msg_list_count": len(msgList),
		"room_id":        roomId,
	})
}

/**
私聊
*/
func PrivateChat(ctx *gin.Context) {
	roomId := ctx.Query("room_id")
	toUid := ctx.Query("uid")

	userInfo := controller.Info(ctx)
	uid := strconv.Itoa(int(userInfo["uid"].(uint)))
	msgList := controller.GetLimitPrivateMsg(uid, toUid, 0)

	ctx.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
		"msg_List":  msgList,
		"room_id":   roomId,
	})
}
