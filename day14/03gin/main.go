package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// urlQuery URL参数
func urlQuery(c *gin.Context) {
	name := c.DefaultQuery("name", "无名氏") // 参数不存在时，返回指定的默认值
	city := c.Query("city")               // 参数不存在时，返回空字符串
	c.String(http.StatusOK, "name: %q, city: %q", name, city)
}

// login POST数据
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	keepLoginStr := c.DefaultPostForm("keep", "false")
	keepLogin, err := strconv.ParseBool(keepLoginStr)
	if err != nil {
		c.String(http.StatusOK, "输入有误：%v", err.Error())
		return
	}
	hobbies := c.PostFormArray("hobby")
	c.String(http.StatusOK, "username: %q, password: %q, keepLogin: %v, hobbies: %v", username, password, keepLogin, hobbies)
}

// upload 单文件上传
func upload(c *gin.Context) {
	attach, err := c.FormFile("attach")
	if err != nil {
		c.String(http.StatusOK, "上传失败：%v", err)
		return
	}
	if err := c.SaveUploadedFile(attach, attach.Filename); err != nil {
		c.String(http.StatusOK, "上传文件保存失败：%v", err)
		return
	}
	c.String(http.StatusOK, "文件上传成功: %v", attach)
}

func mutilUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get MultipartForm failed: %v", err)
		return
	}
	// 获取所有上传
	attachs := form.File["attachs"]
	// 遍历所有上传的图片
	var uploadedFiles = make([]string, 0, len(attachs))
	for _, attach := range attachs {
		// 逐个存储
		if err := c.SaveUploadedFile(attach, attach.Filename); err != nil {
			c.String(http.StatusBadRequest, "SaveUploadedFile failed: %v", err)
			return
		} else {
			uploadedFiles = append(uploadedFiles, attach.Filename)
		}
	}
	c.String(http.StatusOK, "upload success: %v", uploadedFiles)
}

func main() {
	r := gin.Default()
	// 限制上传大小为8m，gin默认是32m
	r.MaxMultipartMemory = 8 << 20
	r.GET("/url_query", urlQuery)
	r.POST("/login", login)
	r.POST("/upload", upload)
    r.POST("/mutilUpload", mutilUpload)
	r.Run(":8000")
}
