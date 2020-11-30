package model

type Music struct {
	BasicModel
	Title  string `json:"title" gorm:"not null" example:"歌曲名称"`
	Year   string `json:"year" example:"年份"`
	Actor  string `json:"actor" gorm:"not null" example:"作曲"`
	Desc   string `json:"desc" example:"描述"`
	UserId uint   `json:"userId" example:"用户id"`
}
