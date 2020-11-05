package model

type Book struct {
	BasicModel
	Title	string `json:"title" gorm:"type:varchar(100);not null" example:"书籍名称"`
	Year	string `json:"year" gorm:"type:varchar(20);not null" example:"年份"`
	Author	string `json:"author" gorm:"not null" example:"作者"`
	Desc	string `json:"desc" example:"描述"`
}
