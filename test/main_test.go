package test

import (
	"app/controller"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type token struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

func getHeaders() map[string]string {
	dir, _ := os.Getwd()
	token, err := os.ReadFile(dir + "/config/token")
	hs := map[string]string{}
	if err == nil {
		hs["Authorization"] = string(token)
	}
	return hs
}

func runGET(reqURL string, t *testing.T) string {
	router := controller.RegisterServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, reqURL, nil)
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runPOST(reqURL string, t *testing.T, arr map[string]string) string {
	router := controller.RegisterServer()
	value := url.Values{}
	if len(arr) > 0 {
		for k, v := range arr {
			value.Add(k, v)
		}
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, reqURL, bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runDelete(reqURL string, t *testing.T) string {
	router := controller.RegisterServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, reqURL, bytes.NewBuffer([]byte("")))
	req.Header.Add("Content-Type", "application/json")
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runPostJSON(reqURL string, t *testing.T, jsonStr string) string {
	router := controller.RegisterServer()
	jsonData := []byte(jsonStr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func TestUserLoginErrorNotUsername(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "", "password": "test"}`)
	assert.Equal(t, `{"code":2001,"errMsg":"请输入账号"}`, res)
}

func TestUserLoginErrorNotPassword(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "test", "password": ""}`)
	assert.Equal(t, `{"code":2001,"errMsg":"请输入密码"}`, res)
}

func TestUserLoginErrorNotFound(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "test", "password": "test"}`)
	assert.Equal(t, `{"code":2001,"errMsg":"未找到此用户"}`, res)
}

func TestUserLoginErrorDisabled(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "user1", "password": "password"}`)
	assert.Equal(t, `{"code":2001,"errMsg":"此账号被禁用"}`, res)
}

func TestUserLoginErrorPass(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "user", "password": "test"}`)
	assert.Equal(t, `{"code":2002,"errMsg":"密码错误"}`, res)
}

func TestUserLoginSucc(t *testing.T) {
	res := runPostJSON("/login", t, `{"username": "user", "password": "password"}`)
	type Res struct {
		Code int
		Data string
	}
	var resJons Res
	json.Unmarshal([]byte(res), &resJons)
	assert.Equal(t, 0, resJons.Code)
}
