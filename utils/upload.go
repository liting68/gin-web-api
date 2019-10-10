package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//Upload 上传文件
func Upload(c *gin.Context, field string, dir string) (bool, string) {
	file, e := c.FormFile(field)
	if e != nil {
		return true, ""
	}
	if file != nil {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e = os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return false, "无法创建文件夹"
		}
		suffix := ""
		if strings.Contains(file.Filename, ".") {
			suffix = file.Filename[strings.LastIndex(file.Filename, "."):]
		} else {
			switch file.Header.Get("Content-Type") {
			case "image/jpeg":
				suffix = ".jpg"
			case "image/png":
				suffix = ".png"
			default:
				suffix = ""
			}
		}
		if suffix == "" {
			return false, "文件格式不正确"
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + suffix
		e = c.SaveUploadedFile(file, path+fileName)
		if e != nil {
			return false, "无法保存文件"
		}
		return true, dir + fileName
	}
	return false, "文件上传失败"
}

//Upload 上传Base64文件
func UploadBase64(fileBase64 string, dir string) (bool, string) {
	if fileBase64 != "" {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e := os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return false, "无法创建文件夹"
		}
		suffix := ""
		fmt.Println(fileBase64[5:strings.Index(fileBase64, ";base64")])
		switch fileBase64[5:strings.Index(fileBase64, ";base64")] {
		case "image/png":
			suffix = ".png"
		case "image/jpeg", "image/jpg":
			suffix = ".jpg"
		}
		if suffix == "" {
			return false, "文件格式不正确"
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + suffix
		data, _ := base64.StdEncoding.DecodeString(fileBase64)
		e = ioutil.WriteFile(path+fileName, data, 0666)
		if e != nil {
			return false, "无法保存文件"
		}
		return true, dir + fileName
	}
	return true, ""
}

func Remove(path string) (bool, string) {
	basedir, _ := os.Getwd()
	file := basedir + path
	e := os.Remove(file)
	if e != nil {
		return false, "删除失败" + e.Error()
	}
	return true, "删除成功"
}
