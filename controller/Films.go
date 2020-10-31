package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/common"
	"goAdmin/model"
	"net/http"
	"strconv"
)

// var db = common.InitDB() // 初始化数据库连接

// @Summary 获取电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} model.Film
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/film/ [get]
func FilmList(ctx *gin.Context)  {
	db := common.InitDB()
	var films []model.Film

	name := ctx.Query("name")
	pageNum,_ := strconv.Atoi(ctx.DefaultPostForm("pageNum","1"))
	pageSize,_ := strconv.Atoi(ctx.DefaultPostForm("pageSize","10"))
	fmt.Println(name,pageNum,pageSize,"--")

	var count int // 总数据条数
	db.Where("name LIKE ?","%name%").Offset((pageNum-1)*pageSize).Order("id desc").Find(&films).Count(count)

	ctx.JSON(http.StatusOK,gin.H{
		"msg":"请求成功",
		"data":films,
		"attr":gin.H{
			"page":pageNum,
			"total":count,
		},
	})
}

// @Summary 创建电影列表
// @Description 创建电影
// @Tags 电影
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param year query string false "year"
// @Param address query string false "address"
// @Param actor query string false "actor"
// @Param desc query string false "desc"
// @Success 200 {object} model.Film
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/v1/film/ [post]
func FilmCreate(ctx *gin.Context) {
	db := common.InitDB()
	var data = &model.Film{}
	err := ctx.Bind(data)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(data,"--")
	db.Create(data)
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"创建成功",
		"data":data,
	})
}

// @Summary 更新电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Film
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v1/film/{id} [put]
func FilmUodate(ctx *gin.Context) {
	db := common.InitDB()
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"--")

	data := &model.Film{}
	err := ctx.Bind(data)
	if err !=nil{
		fmt.Println(err)
		return
	}

	db.Model(data).Where("id=?",id).Update(data)
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"更新成功",
		"data":data,
	})
}

// @Summary 删除电影列表
// @Description 电影列表
// @Tags 电影
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Film
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/v1/film/{id} [delete]
func FilmDelete(ctx *gin.Context)  {
	db := common.InitDB()
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"--")

	db.Where("id=?",id).Delete(model.Film{})
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"删除成功",
		"id":id,
	})
}