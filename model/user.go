package model

type BasicModel struct {
	ID        uint     `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt JSONTime `json:"createAt" gorm:"datetime" example:"创建时间"`
	UpdatedAt JSONTime `json:"updateAt" gorm:"datetime" example:"更新时间"`
}

type User struct {
	BasicModel
	Name       string `json:"name" gorm:"type:varchar(20);not null;unique" example:"用户名称"`
	Telephone  string `json:"telephone" gorm:"varchar(110);not null;unique" example:"手机号码"`
	Password   string `json:"password" gorm:"size:255;not null" example:"密码"` // 字段不暴露给用户，则使用 `json:"-"` 修饰
	Desc       string `json:"desc" gorm:"varchar(225)"`
	IP         string `json:"ip" gorm:"varchar(20)"`
	ImgUrl     string `json:"imgUrl" gorm:"varchar(100)"`
	ThemeColor string `json:"themeColor" gorm:"varchar(20)"`
}

type Visit struct {
	VisitNum int `json:"visit_num" example:"访问次数"`
}

type IpWhite struct {
	BasicModel
	Ip string `json:"ip" gorm:"varchar(60);not null;unique" example:"ip黑名单"`
}
