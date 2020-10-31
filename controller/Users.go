package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goAdmin/common"
	"goAdmin/model"
	"goAdmin/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

// 查询手机号
func isTelephoneExis(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}

// 判断手机号和密码的长度是否正确
func isRight(telephone string,password string,ctx *gin.Context) bool {
	if len(telephone) != 11{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"手机号必须为11位",
		})
		return false
	}

	if len(password) < 6{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":423,
			"msg":"密码不能少于6位",
		})
		return false
	}
	return true
}

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Param telephone query string true "telephone"
// @Param password query string true "password"
// @Success 200 {object} model.User
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/auth/register [post]
func Register(ctx *gin.Context) {
	db := common.InitDB()
	var user = model.User{}
	err := ctx.Bind(&user) // Bind绑定后传json格式
	if err != nil{
		ctx.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}

	// 获取参数
	name := user.Name
	telephone := user.Telephone
	password := user.Password
	// 获取参数
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}

	// 如果没有写名称,就用随机字符串代替
	if len(name) == 0{
		name = util.RandomString(10)
	}
	log.Println(name,password,telephone)
	// 判断手机号是否存在
	if isTelephoneExis(db,telephone){
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":424,
			"msg":"用户已存在",
		})
		return
	}

	// 用户不存在就创建用户
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"msg":"加密发送错误",
		})
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK,gin.H{"msg":"注册成功"})
}

// @Summary 用户登陆
// @Description 用户登陆
// @Tags 用户
// @Accept json
// @Produce json
// @Param telephone query string true "telephone"
// @Param password query string true "password"
// @Success 200 {object} model.User
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/auth/login [post]
func Login(ctx *gin.Context)  {
	db := common.InitDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	isReturn := isRight(telephone,password,ctx)
	if !isReturn {
		return
	}

	// 判断手机号是否存在
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID == 0{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"msg":"用户不存在"})
		return
	}

	// 判断密码是否正确
	fmt.Println(user.Password,"---")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest,gin.H{"msg":"密码错误"})
		return
	}

	// 发送token
	token,err := common.ReleaseToken(user)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"msg":"系统异常",
		})
		fmt.Println(err)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK,gin.H{
		"data":gin.H{"token":token},
		"msg":"登录成功",
	})
}

// @Summary 单个用户信息
// @Description 用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param telephone query string false "telephone"
// @Param token query string true "token"
// @Success 200 {object} model.UserDto
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/auth/info [post]
func Info(ctx *gin.Context) {
	user,_ := ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{
		"data":gin.H{"user":model.ToUserDto(user.(model.User))},
	})
}

// @Summary 获取用户列表
// @Description 用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param telephone query string false "telephone"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} model.User
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/users/ [get]
func UserList(ctx *gin.Context){
	var users []model.User
	name := ctx.Query("name")
	pageNum,_ := strconv.Atoi(ctx.DefaultPostForm("pageNum","1"))
	pageSize,_ := strconv.Atoi(ctx.DefaultPostForm("pageSize","10"))
	fmt.Println(name,pageNum,pageSize,"--")
	/*
		迷糊搜索，name为搜索的条件，根据电影的名称name来搜索
		Offset 其实条数
		Limit	 每页的条数
		Order("id desc") 根据id倒序排序
		总条数 Count(&count)
	*/
	db := common.InitDB()
	var count int
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Where("name LIKE?","%" + name + "%").Order("created_at desc").Find(&users).Count(&count)

	ctx.JSON(http.StatusOK,gin.H{
		"msg":"请求成功",
		"data":users,
		"attr":gin.H{
			"page":pageNum,
			"total":count,
		},
	})
}

// @Summary 删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} model.User
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/userList/{id} [delete]
func UserDelete(ctx *gin.Context)  {
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		panic(err)
	}

	fmt.Println(id,"--")
	db := common.InitDB()
	db.Where("id=?",id).Delete(model.User{})
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"删除成功",
	})
}
