package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
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

var db = common.InitDB() // 初始化数据库连接

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
	desc := user.Desc

	fmt.Println(user,"user")
	log.Println(name,password,telephone)
	// 获取参数
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}

	// 如果没有写名称,就用随机字符串代替
	if len(name) == 0{
		name = util.RandomString(10)
	}

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
		Desc: desc,
	}
	db.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK,gin.H{"msg":"注册成功"})
}

// @Summary 更新用户信息
// @Description 用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.User
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v1/users/{id} [put]
func ChangePassword(ctx *gin.Context){
	id,_ := strconv.Atoi(ctx.Param("id"))

	data := &model.User{}
	hasedPassword,_ := bcrypt.GenerateFromPassword([]byte(data.Password),bcrypt.DefaultCost)
	fmt.Println(string(hasedPassword),"-***-")

	err := ctx.Bind(data)
	data.Password = string(hasedPassword)
	if err != nil{
		fmt.Println(err)
		return
	}

	db.Model(data).Where("id=?",id).Update(data)
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"更新成功",
		"data":data,
	})
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
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	code := ctx.PostForm("code")

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

	user.IP = ctx.ClientIP() // 给user的ip字段赋值
	db.Save(&user)	// 保存并更新数据

	// 判断密码是否正确
	fmt.Println(user.Password,"password")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest,gin.H{"msg":"密码错误"})
		return
	}

	// 判断验证码是否正确
	session := sessions.Default(ctx)
	captchaId := session.Get("captchaId").(string)

	verifyResult := util.VerfiyCaptcha(captchaId,code)
	if !verifyResult{
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"code":http.StatusUnauthorized,
			"msg":"验证码输入错误",
		})
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
		"data":gin.H{"token":token,"name":user.Name,"ip":user.IP},
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
// @Router /api/v1/auth/info [get]
func Info(ctx *gin.Context) {
	user,_ := ctx.Get("user")
	fmt.Println(user,"user")
	publicIP := util.ClientPublicIP() // 和ctx.ClientIP() 内置方法差不多
	clientIp := util.ClientIP(ctx.Request)
	ctx.JSON(http.StatusOK,gin.H{
		"data":gin.H{"user":model.ToUserDto(user.(model.User))},
		"clientIp":clientIp,
		"公网ip":publicIP,
		"客户ip":ctx.ClientIP(),
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
// @Router /api/v1/users [get]
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
	var count int
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Where("name LIKE?","%" + name + "%").Order("created_at desc").Find(&users).Count(&count)

	ip := util.GetClientIp()
	serverIp := util.GetServerIP()
	RemoteIP := util.RemoteIP(ctx.Request)
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"请求成功",
		"data":users,
		"ip":ip,
		"serverIp":serverIp,
		"RemoteIP":RemoteIP,
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
// @Router /api/v1/users/{id} [delete]
func UserDelete(ctx *gin.Context) {
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		panic(err)
	}

	var u model.User
	db.Where("id=?",id).Find(&u)	// 查找到对应id的整行数据
	fmt.Println(u)
	if u.Name == "coco"{	// 查找对应id的数据中判断是否有name为coco的数据
		ctx.JSON(http.StatusOK,gin.H{
			"msg":"coco是管理员不能删除",
		})
		return
	}

	fmt.Println(id,"--")
	db.Where("id=?",id).Delete(model.User{})
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"删除成功",
	})
}
