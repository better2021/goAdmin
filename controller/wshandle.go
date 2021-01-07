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

func SaveContent(value interface{}) model.Message {
	var m model.Message
	m.UserId = value.(map[string]interface{})["user_id"].(int)
	m.ToUserId = value.(map[string]interface{})["to_user_id"].(int)
	m.Content = value.(map[string]interface{})["content"].(string)

	roomIdStr := value.(map[string]interface{})["room_id"].(string)
	roomIdInt, _ := strconv.Atoi(roomIdStr)

	m.RoomId = roomIdInt

	if _, ok := value.(map[string]interface{})["image_url"]; ok {
		m.ImgUrl = value.(map[string]interface{})["img_url"].(string)
	}
	db.Create(&m)
	return m
}

func GetLimitMsg(roomId string, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	var m model.Message
	db.Model(&m).
		Select("messages.*,users.name,users.img_url").
		Joins("INNER Join users on users.id = messages.user_id").
		Where("messages.room_id="+roomId).
		Where("messages.to_user_id = ?", 0).
		Order("messages.id desc").
		Offset(100).
		Scan(&results)

	if offset == 0 {
		sort.Slice(results, func(i, j int) bool {
			return results[i]["id"].(uint32) < results[j]["id"].(uint32)
		})
	}
	return results
}

func GetLimitPrivateMsg(uid, toUId string, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	db.Model(&model.Message{}).
		Select("messages.*,users.name,users.img_url").
		Joins("INNER Join users on users.id = messages.user_id").
		Where("(" +
			"(" + "messages.user_id = " + uid + " and messages.to_user_id=" + toUId + ")" +
			" or " +
			"(" + "messages.user_id = " + toUId + " and messages.to_user_id=" + uid + ")" +
			")").
		Order("messages.id desc").
		Offset(offset).
		Limit(100).
		Scan(&results)

	if offset == 0 {
		sort.Slice(results, func(i, j int) bool {
			return results[i]["id"].(uint32) < results[j]["id"].(uint32)
		})
	}
	return results
}
