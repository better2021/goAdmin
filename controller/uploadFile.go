package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

// @Summary 单个文件上传
// @Description 文件上传
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file query string true "file"
// @Success 200 {string} string "{ "code": 200, "message": "上传成功" }"
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/upload [post]
func UploadFile(ctx *gin.Context)  {
	file,header,err := ctx.Request.FormFile("file")
	filename := header.Filename
	fmt.Println(filename,"---")

	// 获取当前目录
	dir,err := os.Getwd()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dir,"**")
	// 创建新目录
	os.Mkdir(dir + "/uploadFiles",0777)

	out,err := os.Create(dir + "/uploadFiles/" + filename)
	if err != nil{
		fmt.Println(err)
	}
	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil{
		fmt.Println(err)
	}

	var protocol string
	if ctx.Request.Proto== "HTTP/1.1" {
		protocol = "http://"
	}else {
		protocol = "https://"
	}

	// time.Now().Format("2006-01-02 15:04:05") 格式化时间
	ctx.JSON(http.StatusOK, gin.H{
		"code":http.StatusOK,
		"msg":    "上传成功",
		"imgUrl":     protocol + ctx.Request.Host + "/static/" + filename,
		"createTime": time.Now().Format("2006-01-02 15:04:05"),
	})
}