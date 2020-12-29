package home

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

// base64文件上传
func UploadBase64(c *gin.Context) {
	base64File := c.Query("file")
	ext := c.Query("ext")

	// 生成文件保存目录
	t := time.Now()
	var rootPath = "upload/"
	var savePath = t.Format("2006-01-02") + "/"
	// 判断该目录是否存在, 不存在则进行创建
	exists, _ := PathExists(rootPath + savePath)
	if !exists {
		err2 := os.Mkdir(rootPath+savePath, os.ModePerm)
		if err2 != nil {
			log.Error(err2.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  "目录创建失败, 原因：" + err2.Error(),
			})
			return
		}
	}
	// 生成文件名
	fileName := xid.New().String() + "." + ext

	// base64解码
	fileContent, err3 := base64.StdEncoding.DecodeString(base64File)
	if err3 != nil {
		log.Error(err3.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "目录创建失败, 原因：" + err3.Error(),
		})
		return
	}
	// 将解码的文件内容写入文件
	err4 := ioutil.WriteFile(rootPath+savePath+fileName, fileContent, 0666)
	if err4 != nil {
		log.Error(err4.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "文件写入失败：" + err4.Error(),
		})
		return
	}

	// 返回执行结果json(包含上传文件的具体信息)
	c.JSON(http.StatusOK, gin.H{
		"code":     "200",
		"url":      "/upload/" + savePath + fileName,
		"original": "",
		"size":     0,
		"state":    "SUCCESS",
		"title":    fileName,
		"type":     ext,
	})

}

// 单文件上传
func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "文件上传失败, 原因：" + err.Error(),
		})
		return
	}
	// 生成文件保存目录
	t := time.Now()
	var rootPath = "upload/"
	var savePath = t.Format("2006-01-02") + "/"
	// 判断该目录是否存在, 不存在则进行创建
	exists, _ := PathExists(rootPath + savePath)
	if !exists {
		err2 := os.Mkdir(rootPath+savePath, os.ModePerm)
		if err2 != nil {
			log.Error(err2.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  "目录创建失败, 原因：" + err2.Error(),
			})
			return
		}
	}

	// 生成文件名
	ext := path.Ext(file.Filename)
	fileName := xid.New().String() + ext

	// 保存文件
	err3 := c.SaveUploadedFile(file, rootPath+savePath+fileName)
	if err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "文件保存失败，原因：" + err3.Error(),
		})
		return
	}

	// 返回执行结果json(包含上传文件的具体信息)
	c.JSON(http.StatusOK, gin.H{
		"code":     "200",
		"url":      "/upload/" + savePath + fileName,
		"original": file.Filename,
		"size":     file.Size,
		"state":    "SUCCESS",
		"title":    fileName,
		"type":     ext,
	})
}

// 判断目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
