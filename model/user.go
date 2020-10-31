package model

type BasicModel struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt Time `json:"createAt" example:"创建时间"`
	UpdatedAt Time `json:"updateAt" example:"更新时间"`
}

type User struct {
	BasicModel
	Name string `json:"name" gorm:"type:varchar(20);not null" example:"用户名称"`
	Telephone string `json:"telephone" gorm:"varchar(110);not null;unique" example:"手机号码"`
	Password string `json:"password" gorm:"size:255;not null" example:"密码"`
}

type UserDto struct {
	BasicModel
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user User) UserDto{
	return UserDto{
		Name:     user.Name,
		Telephone: user.Telephone,
	}
}