package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/model"
	"net/http"
	"strconv"
)

// @Summary 获取书籍列表
// @Description 书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param name query string false "title"
// @Param pageNum query string true "pageNum"
// @Param pageSize query string true "pageSize"
// @Success 200 {object} model.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/books [get]
func BookList(ctx *gin.Context)  {
	var books []model.Book

	title := ctx.Query("title")
	pageNum,_ := strconv.Atoi(ctx.DefaultPostForm("pageNum","1"))
	pageSize,_ := strconv.Atoi(ctx.DefaultPostForm("pageSize","10"))
	fmt.Println(title,"--")

	var count int
	db.Where("title LIKE ?","%title%").Offset((pageNum-1)*pageSize).Order("id desc").Find(&books).Count(&count)

	ctx.JSON(http.StatusOK,gin.H{
		"msg":"请求成功",
		"data":books,
		"attr":gin.H{
			"page":pageNum,
			"total":count,
		},
	})
}

// @Summary 创建书籍列表
// @Description 创建书籍
// @Tags 书籍
// @Accept json
// @Produce json
// @Param name query string false "title"
// @Param year query string false "year"
// @Param actor query string false "author"
// @Param desc query string false "desc"
// @Success 200 {object} model.Book
// @Failure 400 {string} json "{ "code": 400, "message": "请求失败" }"
// @Router /api/books [post]
func BookCreate(ctx *gin.Context){
	var data = &model.Book{}
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

// @Summary 更新书籍列表
// @Description 书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Book
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/books/{id} [put]
func BookUpdate(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"--")

	data := &model.Book{}
	err := ctx.Bind(data)
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

// @Summary 删除书籍列表
// @Description 书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Book
// @Failure 400 {string} json "{ "code": 400, "message": "id必传" }"
// @Router /api/books/{id} [delete]
func BookDelete(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id,"--")

	db.Where("id=?",id).Delete(&model.Book{})
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"删除成功",
		"id":id,
	})
}