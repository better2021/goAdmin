package model

type Film struct{
	BasicModel
	Name    *string `gorm:"unique;not null" json:"name" example:"电影名称"`
	Year    string  `json:"year" example:"年份"`
	Address string  `json:"address" example:"出品地区"`
	Actor   string  `json:"actor" example:"演员"`
	Desc    string  `json:"desc" example:"描述"`
}
