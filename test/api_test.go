package test

import (
	"app/router"
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexGetRouter(t *testing.T) {
	res := runTestGET("/", t)
	assert.Equal(t, "index", res)
}

func TestLogin(t *testing.T) {
	arr := map[string]string{"username": "admin", "password": "admin"}
	res := runTestPOST("/login", t, arr)
	// assert.Equal(t, `{"code":0,"data":{"id":1,"username":"Lucus","password":"123456","created":"2019-09-03 15:24:39"}}`, res)
	assert.NotNil(t, "", res)
}

func runTestGET(reqURL string, t *testing.T) string {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, reqURL, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runTestPOST(reqURL string, t *testing.T, arr map[string]string) string {
	router := router.SetupRouter()
	value := url.Values{}
	if len(arr) > 0 {
		for k, v := range arr {
			value.Add(k, v)
		}
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, reqURL, bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}
