package model

type Message struct {
	BasicModel
	UserId   int    `json:"user_id"`
	ToUserId int    `json:"to_user_id"`
	RoomId   int    `json:"room_id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"` //  图片消息
	ImgUrl   string `json:"img_url"`   // 用户头像
	UserName string `json:"user_name"`
}
