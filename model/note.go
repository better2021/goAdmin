package model

// 留言板 model
type Note struct {
	BasicModel
	UserName    string `gorm:"not null" json:"userName" example:"用户名称"`
	Images  string `json:"images" example:"图片集"`
	Icon string `json:"icon" example:"用户头像"`
	Context string `json:"context" example:"留言"`
	UserId uint `json:"userId" example:"用户id"`
}

