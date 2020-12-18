package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/model"
	"net/http"
	"strconv"
)

// @Summary 白名单列表
// @Description 白名单
// @Tags 白名单
// @Accept json
// @Produce json
// @Success 200 {object} model.IpWhite
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/ipWhite [get]
func IpList(ctx *gin.Context)  {
	var ips []model.IpWhite
	db.Find(&ips)
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"ipList":ips,
	})
}

// @Summary 创建白名单列表
// @Description 创建白名单
// @Tags 白名单
// @Accept json
// @Produce json
// @Param title query string true "ip"
// @Success 200 {object} model.IpWhite
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/ipWhite [post]
func IpsCreate(ctx *gin.Context){
	var data = &model.IpWhite{}
	err := ctx.Bind(data)
	if err != nil{
		fmt.Println(err)
		return
	}

	db.Create(data)
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"新增成功",
		"data":data,
	})
}

// @Summary 删除白名单列表
// @Description 白名单列表
// @Tags 白名单
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.IpWhite
// @Failure 400 {string} string "{ "code": 400, "message": "id必传" }"
// @Router /api/ipWhite/{id} [delete]
func IpsDelete(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))

	db.Where("id = ?",id).Delete(&model.IpWhite{})
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"删除成功",
		"id":id,
	})
}