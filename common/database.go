package common

import (
	"fmt"
	"goAdmin/model"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	InitConfig() // 获取数据库配置
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	// root:709463253@/goadmin?charset=utf8&parseTime=True&loc=Local

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(driverName, args)
	fmt.Println(driverName, args, "---")
	if err != nil {
		panic("faild to connect database,err" + err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
<<<<<<< HEAD
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate(&model.User{}, &model.UserDto{}, &model.Film{}, &model.Visit{})
=======
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;")
	db.AutoMigrate(&model.User{},&model.UserDto{},&model.Film{},&model.Visit{})
>>>>>>> a9d031bbff8e794fb30a2d9c90f9eb4b0c38afc1
	return db
}

// 获取数据库初始化配置
func InitConfig() {
	// 获取当前的工作目录
	workDir, _ := os.Getwd()
	fmt.Println("当前文件的路劲:" + workDir)
	// 设置要读取的文件名
	viper.SetConfigName("application")
	// 设置要读取的文件的类型
	viper.SetConfigType("yml")
	// 添加读取文件的路劲
	viper.AddConfigPath(workDir + "/config")
	// 读取文件配置
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err, "---")
	}
}
