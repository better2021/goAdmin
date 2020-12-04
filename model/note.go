package model

// 留言板 model
type Note struct {
	BasicModel
	UserName    string `gorm:"not null" json:"userName" example:"用户名称"`
	Images  []Image `json:"images" example:"图片集"`
	Icon string `json:"icon" example:"用户头像"`
	Context string `json:"context" example:"留言"`
}

type  Image struct {
	Img  string `json:"img" example:"图片"`
}

