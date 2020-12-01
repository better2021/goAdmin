package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/model"
	"net/http"
	"strconv"
)

// @Summary 获取音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param title query string false "title"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} model.Music
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/musics [get]
func MusicList(ctx *gin.Context){
	var musics []model.Music

	title := ctx.Query("title")
	userId,_ := strconv.Atoi(ctx.Query("userId"))
	pageNum,_ := strconv.Atoi(ctx.DefaultQuery("pageNum","1"))
	pageSize,_:= strconv.Atoi(ctx.DefaultQuery("pageSize","10"))
	fmt.Println(title,"--")

	var count int
	db.Model(&musics).Where("title LIKE ? AND user_id = ?","%" + title + "%",userId).Count(&count)
	db.Where("title LIKE ? AND user_id = ?","%"+ title + "%",userId).Offset((pageNum - 1)*pageSize).Limit(pageSize).Order("id desc").Find(&musics)

	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"请求成功",
		"data":musics,
		"attr":gin.H{
			"page":pageNum,
			"total":count,
		},
	})
}

// @Summary 创建音乐列表
// @Description 创建音乐
// @Tags 音乐
// @Accept json
// @Produce json
// @Param title query string false "title"
// @Param year query string false "year"
// @Param actor query string false "author"
// @Param desc query string false "desc"
// @Success 200 {object} model.Music
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/musics [post]
func MusicCreate(ctx *gin.Context){
	var data = &model.Music{}
	err := ctx.Bind(data)
	if err != nil{
		fmt.Println(err)
		return
	}

	db.Create(data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"创建成功",
		"data":data,
	})
}

// @Summary 更新音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Music
// @Failure 400 {string} string "{ "code": 400, "message": "id必传" }"
// @Router /api/musics/{id} [put]
func MusicUpdate(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"id")

	data := &model.Music{}
	err := ctx.Bind(data)
	if err != nil{
		fmt.Println(err)
		return
	}

	db.Model(data).Where("id = ?",id).Update(data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"更新成功",
		"data":data,
	})
}

// @Summary 删除音乐列表
// @Description 音乐列表
// @Tags 音乐
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Music
// @Failure 400 {string} string "{ "code": 400, "message": "id必传" }"
// @Router /api/musics/{id} [delete]
func MusicDelete(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"id")

	db.Where("id=?",id).Delete(&model.Music{})
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"删除成功",
		"id":id,
	})
}