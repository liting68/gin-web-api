/*
 * @Author: hiwein.lucus
 * @Date: 2019-10-08 15:41:54
 * @Last Modified by:   hiwein.lucus
 * @Last Modified time: 2019-10-08 15:41:54
 */
package test

import (
	"app/lib"
	"bytes"
	"io/ioutil"
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
	token, err := ioutil.ReadFile(dir + "/config/token")
	hs := map[string]string{}
	if err == nil {
		hs["Authorization"] = string(token)
	}
	return hs
}

func runGET(reqURL string, t *testing.T) string {
	router := lib.InitGin()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, reqURL, nil)
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runPOST(reqURL string, t *testing.T, arr map[string]string) string {
	router := lib.InitGin()
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
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runPostJSON(reqURL string, t *testing.T, jsonStr string) string {
	router := lib.InitGin()
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
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}

func runDeleteJSON(reqURL string, t *testing.T, jsonStr string) string {
	router := lib.InitGin()
	jsonData := []byte(jsonStr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, reqURL, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	headers := getHeaders()
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	return w.Body.String()
}
