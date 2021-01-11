package controller

import (
	"goAdmin/model"
	"sort"
	"strconv"
)

/**
ws的方法
*/

func AddUser(value interface{}) model.User {
	var u model.User
	u.Name = value.(map[string]interface{})["name"].(string)
	u.Password = value.(map[string]interface{})["password"].(string)
	u.ImgUrl = value.(map[string]interface{})["img_url"].(string)
	db.Create(&u)
	return u
}

func SaveImgUrl(ImgUrl string, u model.User) model.User {
	u.ImgUrl = ImgUrl
	db.Save(&u)
	return u
}

func GetOnlineUserList(uids []float64) []map[string]interface{} {
	var results []map[string]interface{}
	db.Where("id IN ?", uids).Find(&results)

	return results
}

// 存消息
func SaveContent(value interface{}) model.Message {
	var m model.Message
	m.UserId = value.(map[string]interface{})["user_id"].(int)
	m.ToUserId = value.(map[string]interface{})["to_user_id"].(int)
	m.Content = value.(map[string]interface{})["content"].(string)
	m.ImgUrl = value.(map[string]interface{})["img_url"].(string)
	m.UserName = value.(map[string]interface{})["username"].(string)

	roomIdStr := value.(map[string]interface{})["room_id"].(string)
	roomIdInt, _ := strconv.Atoi(roomIdStr)
	m.RoomId = roomIdInt

	if _, ok := value.(map[string]interface{})["image_url"]; ok {
		m.ImageUrl = value.(map[string]interface{})["image_url"].(string)
	}

	db.Create(&m)
	return m
}

// 房间聊天记录
func GetLimitMsg(roomId string, offset int) []model.Message {
	var results []model.Message

	db.Model(&results).Select("messages.*, users.name ,users.img_url").
		Joins("INNER Join users on users.id = messages.user_id").
		Where("messages.room_id = ? AND messages.to_user_id = ?", roomId, 0).
		Order("messages.id desc").Offset(offset).
		Limit(20).Scan(&results)

	if offset == 0 {
		sort.Slice(results, func(i, j int) bool { // 排序
			return results[i].ID < results[j].ID
		})
	}

	return results
}

// 私人聊天记录
func GetLimitPrivateMsg(uid, toUId string, offset int) []model.Message {
	var results []model.Message

	db.Model(&results).
		Select("messages.*,users.name,users.img_url").
		Joins("INNER Join users on users.id = messages.user_id").
		Where("(" +
			"(" + "messages.user_id = " + uid + " and messages.to_user_id=" + toUId + ")" +
			" or " +
			"(" + "messages.user_id = " + toUId + " and messages.to_user_id=" + uid + ")" +
			")").
		Order("messages.id desc").
		Offset(offset).
		Limit(20).
		Scan(&results)

	if offset == 0 {
		sort.Slice(results, func(i, j int) bool {
			return results[i].ID < results[j].ID
		})
	}
	return results
}
