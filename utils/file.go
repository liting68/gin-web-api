package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FileRemove 删除文件
func FileRemove(filePath string) {
	basedir, _ := os.Getwd()
	path := basedir + filePath
	if FileExists(path) {
		err := os.Remove(path)
		if err != nil {
			log.Printf("文件删除失败：%s \n", err.Error())
		}
	}
}

// FileMd5 获取文件的md5码
func FileMd5(filePath string) string {
	basedir, _ := os.Getwd()
	path := basedir + filePath
	pFile, err := os.Open(path)
	if err != nil {
		log.Printf("打开文件失败，filename=%v, err=%v", filePath, err)
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil))
}
