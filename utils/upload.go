package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件
func Upload(c *gin.Context, field string, dir string, sufx string) (bool, string) {
	file, e := c.FormFile(field)
	if e != nil {
		log.Println(e.Error())
		return false, "文件上传失败：" + e.Error()
	}
	if file != nil {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e = os.MkdirAll(path, os.ModePerm)
		if e != nil {
			log.Println(e.Error())
			return false, "无法创建文件夹"
		}
		suffix := ""
		if strings.Contains(file.Filename, ".") {
			suffix = file.Filename[strings.LastIndex(file.Filename, "."):]
		} else {
			switch file.Header.Get("Content-Type") {
			case "image/png":
				suffix = ".png"
			case "image/jpeg":
				suffix = ".jpg"
			case "image/jpg":
				suffix = ".jpg"
			case "audio/wav":
				suffix = ".wav"
			case "audio/mp4":
				suffix = ".mp4"
			}
		}
		if suffix == "" {
			return false, "文件格式不正确"
		}
		fileName := strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + sufx + suffix
		e = c.SaveUploadedFile(file, path+fileName)
		if e != nil {
			log.Println(e.Error())
			return false, "无法保存文件"
		}
		return true, dir + fileName
	}
	return false, "文件上传失败"
}

// UploadFiles 上传多个文件
func UploadFiles(c *gin.Context, field string, dir string) ([]string, error) {
	form, e := c.MultipartForm()
	if e != nil {
		fmt.Println("文件上传失败：" + e.Error())
		return []string{""}, e
	}
	files := form.File[field]
	filepathArr := []string{}
	for k, file := range files {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e = os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return []string{"无法创建文件夹"}, e
		}
		suffix := ""
		if strings.Contains(file.Filename, ".") {
			suffix = file.Filename[strings.LastIndex(file.Filename, "."):]
		} else {
			switch file.Header.Get("Content-Type") {
			case "image/png":
				suffix = ".png"
			case "image/jpeg":
				suffix = ".jpg"
			case "image/jpg":
				suffix = ".jpg"
			case "audio/wav":
				suffix = ".wav"
			}
		}
		if suffix == "" {
			return []string{"文件格式不正确"}, e
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(k) + suffix
		e = c.SaveUploadedFile(file, path+fileName)
		if e != nil {
			fmt.Println("文件上传错误：", e.Error())
			return []string{"无法保存文件"}, e
		}
		filepathArr = append(filepathArr, dir+fileName)
	}
	return filepathArr, nil
}

// UploadFiles 上传多个文件
func UploadFilesBase64(strarr []string, dir string) ([]string, error) {
	filepathArr := []string{}
	for k, fileBase64 := range strarr {
		if fileBase64 != "" {
			basedir, _ := os.Getwd()
			path := basedir + dir
			e := os.MkdirAll(path, os.ModePerm)
			if e != nil {
				return filepathArr, fmt.Errorf("无法创建文件夹")
			}
			suffix := ""
			if len(fileBase64) > 200 {
				log.Printf("文件内容：%s", fileBase64[0:200])
			} else {
				log.Printf("文件内容：%s", fileBase64)
			}
			if !strings.Contains(fileBase64, ";base64") {
				return filepathArr, fmt.Errorf("文件内容需要BASE64格式编码：" + fileBase64)
			}
			if len(fileBase64) < 30 {
				return filepathArr, fmt.Errorf("文件解析内容不正确：" + fileBase64)
			}
			switch fileBase64[5:strings.Index(fileBase64, ";base64")] {
			case "image/png":
				suffix = ".png"
			case "image/jpeg":
				suffix = ".jpg"
			case "image/jpg":
				suffix = ".jpg"
			case "audio/wav":
				suffix = ".wav"
			}
			if suffix == "" {
				return filepathArr, fmt.Errorf("文件格式不正确")
			}
			data, _ := base64.StdEncoding.DecodeString(fileBase64[strings.LastIndex(fileBase64, "base64,")+7:])
			fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(k) + suffix
			// e = os.WriteFile(path+fileName, data, 0666)
			f, e := os.OpenFile(path+fileName, os.O_WRONLY|os.O_CREATE, 0666) //创建文件
			if e != nil {
				log.Println(e)
				return filepathArr, e
			}
			defer f.Close()
			e = os.WriteFile(path+fileName, data, 0666)
			if e != nil {
				return filepathArr, e
			}
			log.Println("upload success " + path + fileName)
			filepathArr = append(filepathArr, dir+fileName)
		}
	}
	return filepathArr, nil
}

// UploadBase64 上传Base64文件
func UploadBase64(fileBase64 string, dir string) (bool, string) {
	if len(fileBase64) < 200 {
		log.Printf("UploadBase64 ===%s\n", fileBase64)
	} else {
		log.Printf("UploadBase64 ===%s\n", fileBase64[0:200])
	}
	if fileBase64 != "" {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e := os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return false, "无法创建文件夹"
		}
		suffix := ""
		if len(fileBase64) > 200 {
			log.Printf("文件内容：%s", fileBase64[0:200])
		} else {
			log.Printf("文件内容：%s", fileBase64)
		}
		if !strings.Contains(fileBase64, ";base64") {
			return false, "文件内容需要BASE64格式编码：" + fileBase64
		}
		if len(fileBase64) < 30 {
			return false, "文件解析内容不正确：" + fileBase64
		}
		switch fileBase64[5:strings.Index(fileBase64, ";base64")] {
		case "image/png":
			suffix = ".png"
		case "image/jpeg":
			suffix = ".jpg"
		case "image/jpg":
			suffix = ".jpg"
		case "audio/wav":
			suffix = ".wav"
		}
		if suffix == "" {
			return false, "文件格式不正确"
		}
		data, _ := base64.StdEncoding.DecodeString(fileBase64[strings.LastIndex(fileBase64, "base64,")+7:])
		fileName := strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + suffix
		f, e := os.OpenFile(path+fileName, os.O_WRONLY|os.O_CREATE, 0666) //创建文件
		if e != nil {
			return false, e.Error()
		}
		defer f.Close()
		if suffix == ".png" {
			m, _ := png.Decode(bytes.NewBuffer(data))
			e = png.Encode(f, m) //写入文件
			if e != nil {
				log.Println(e)
			}
		} else if suffix == ".jpg" || suffix == ".jpeg" {
			m, _ := jpeg.Decode(bytes.NewBuffer(data))
			e = jpeg.Encode(f, m, &jpeg.Options{Quality: 75}) //写入文件
		} else {
			e = os.WriteFile(path+fileName, data, 0666)
		}
		if e != nil {
			return false, e.Error()
		}
		log.Println("upload success " + path + fileName)
		return true, dir + fileName
	}
	return true, ""
}

// Remove 删除文件
func Remove(path string) (bool, string) {
	basedir, _ := os.Getwd()
	file := basedir + path
	e := os.Remove(file)
	if e != nil {
		return false, "删除失败" + e.Error()
	}
	return true, "删除成功"
}

// UploadLog 上传日志
func UploadLog(c *gin.Context, field string, dir string) (bool, string) {
	file, e := c.FormFile(field)
	if e != nil {
		fmt.Println("文件上传失败：" + e.Error())
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
			case "text/plain":
				suffix = ".log"
			default:
				suffix = ""
			}
		}
		if suffix != ".log" && suffix != ".txt" {
			return false, "文件格式不正确"
		}
		fileName := strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + suffix
		e = c.SaveUploadedFile(file, path+fileName)
		if e != nil {
			return false, "无法保存文件"
		}
		return true, dir + fileName
	}
	return false, "文件上传失败"
}

// UploadBase64Fname 上传Base64文件
func UploadBase64Fname(fileBase64 string, dir string, fname string) string {
	if fileBase64 != "" {
		basedir, _ := os.Getwd()
		path := basedir + dir
		e := os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return ""
		}
		suffix := ""
		switch fileBase64[5:strings.Index(fileBase64, ";base64")] {
		case "image/png":
			suffix = ".png"
		case "image/jpeg":
			suffix = ".jpg"
		case "image/jpg":
			suffix = ".jpg"
		case "audio/wav":
			suffix = ".wav"
		}
		if suffix == "" {
			return ""
		}
		fname += suffix
		data, _ := base64.StdEncoding.DecodeString(fileBase64[strings.LastIndex(fileBase64, "base64,")+7:])
		// e = os.WriteFile(path+fileName, data, 0666)
		f, e := os.OpenFile(path+fname, os.O_WRONLY|os.O_CREATE, 0666) //创建文件
		if e != nil {
			return ""
		}
		defer f.Close()
		e = os.WriteFile(path+fname, data, 0666)
		if e != nil {
			return ""
		}
		return dir + fname
	}
	return ""
}
