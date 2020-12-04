package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// @Summary 多个文件上传
// @Description 多文件上传
// @Tags 多文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file query string true "upload"
// @Success 200 {string} string "{ "code": 200, "message": "上传成功" }"
// @Failure 400 {string} string "{ "code": 400, "message": "请求失败" }"
// @Router /api/uploads [post]
func UploadFiles(ctx *gin.Context)  {
	formdata := ctx.Request.MultipartForm
	files := formdata.File["upload"]
	num := len(files) // 文件数量

	// 获取当前目录
	dir,err := os.Getwd()
	if err != nil{
		fmt.Println(err)
		return
	}
	// 创建新目录
	os.Mkdir(dir + "/uploadFiles",0777)

	var protocol string
	if ctx.Request.Proto== "HTTP/1.1" {
		protocol = "http://"
	}else {
		protocol = "https://"
	}

	var images []string
	for i,_ := range files{
		file,err := files[i].Open()
		defer file.Close()
		if err != nil{
			panic(err)
		}

		out,err := os.Create(dir + "/uploadFiles" + files[i].Filename)
		defer out.Close()
		if err != nil{
			fmt.Println(err)
			return
		}
		_,err = io.Copy(out,file)
		if err != nil{
			fmt.Println(err)
			return
		}

		var imgUrl = protocol + ctx.Request.Host + "/static/" + files[i].Filename
		images = append(images,imgUrl)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"上传成功",
		"images":images,
		"len":num,
	})
}