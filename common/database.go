package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"goAdmin/model"
	"net/url"
)

var db *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB{
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	// root:709463253@/gin?charset=utf8&parseTime=True&loc=Local
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db,err := gorm.Open(driverName,args)
	fmt.Println(driverName,args,"---")
	if err != nil {
		panic("faild to connect database,err" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}