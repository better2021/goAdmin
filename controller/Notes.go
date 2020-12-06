package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/model"
	"net/http"
	"strconv"
)

// @Summary 获取留言列表
// @Description 留言列表
// @Tags 留言
// @Accept json
// @Produce json
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} model.Note
// @Failure 400 {string} string  "{ "code": 400, "message": "请求失败" }"
// @Router /api/notes [get]
func NoteList(ctx *gin.Context)  {
	var notes []model.Note

	pageNum,_ := strconv.Atoi(ctx.DefaultQuery("pageNum","1"))
	pageSize,_ := strconv.Atoi(ctx.DefaultQuery("pageSize","10"))
	fmt.Println(pageNum,pageSize)

	var count int
	db.Offset((pageNum - 1)*pageSize).Limit(pageSize).Order("created_at desc").Find(&notes).Count(&count)

	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"请求成功",
		"data":notes,
		"attr":gin.H{
			"page":pageSize,
			"total":count,
		},
	})
}

// @Summary 创建留言列表
// @Description 创建留言
// @Tags 留言
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/notes [post]
func NoteCreate(ctx *gin.Context)  {
	var data = &model.Note{}
	err := ctx.ShouldBind(data)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(data,"--")
	db.Create(data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"data":data,
		"msg":"创建成功",
	})
}

// @Summary 删除留言列表
// @Description 留言列表
// @Tags 留言
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Note
// @Failure 400 {string} string "{ "code": 400, "message": "id必传" }"
// @Router /api/notes/{id} [delete]
func NoteDelete(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"id")

	db.Where("id=?",id).Delete(&model.Note{})
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"删除成功",
		"id":id,
	})
}