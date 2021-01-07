package model

type Message struct {
	BasicModel
	UserId   int    `json:"user_id"`
	ToUserId int    `json:"to_user_id"`
	RoomId   int    `json:"room_id"`
	Content  string `json:"content"`
	ImgUrl   string `json:"img_url"`
}
