package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CurlGET 向服务端发送get请求
func CurlGET(url string) (bodystr string) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	// 接收服务端返回给客户端的信息
	res, _ := client.Do(request)
	if res.StatusCode == 200 {
		str, _ := io.ReadAll(res.Body)
		bodystr = string(str)
		// fmt.Println("CurlGET==", bodystr)
	}
	return bodystr

}

// CurlPOST 向服务端发送POST请求
func CurlPOST(url string, parm url.Values) (bodystr string) {
	client := &http.Client{}
	parm.Add("info", "")
	res, err := client.PostForm(url, parm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, _ := io.ReadAll(res.Body)
		bodystr = string(body)
		fmt.Println("CurlPOST==", bodystr)
	}
	return bodystr
}

// CurlJSON 向服务端发送json数据
func CurlJSON(url string, data string) (bodystr string) {
	client := &http.Client{}
	postdata := bytes.NewBuffer([]byte(data))
	res, err := client.Post(url, "application/json", postdata)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, _ := io.ReadAll(res.Body)
		bodystr = string(body)
		fmt.Println("CurlJSON==", bodystr)
	}
	return bodystr
}
