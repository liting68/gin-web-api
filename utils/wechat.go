package utils

import (
	"app/config"
	"app/db"
	"encoding/json"
	"fmt"
	"net/url"
)

//Wechat 微信
type Wechat struct {
	AppID             string
	AppSecret         string
	MsgToken          string
	MsgEncodingAesKey string
}

//WX 微信公众号
var WX *Wechat

//InitWechat 获取初始化微信实例
func init() {
	var w Wechat
	WX = w.SetConfig()
}

//AccessToken 获取AccessToken
func (w *Wechat) AccessToken() string {
	accessToken, _ := db.Redis.Do("get", "weixin_access_token_"+w.AppID).String()
	return accessToken
}

//SetConfig 设置患者微信公众号
func (w *Wechat) SetConfig() *Wechat {
	w.AppID = config.Info.Wechat.User.AppID
	w.AppSecret = config.Info.Wechat.User.AppSecret
	return w
}

//GetCodeRedirect 微信端授权链接
func (w *Wechat) GetCodeRedirect(redirect string, state string) string {
	u := url.Values{}
	u.Add("appid", w.AppID)
	u.Add("redirect_uri", redirect)
	u.Add("response_type", "code")
	u.Add("scope", "snsapi_userinfo")
	u.Add("state", state)
	uri := "https://open.weixin.qq.com/connect/oauth2/authorize?" + u.Encode() + "#wechat_redirect"
	return uri
}

//GetQRCodeRedirect WEB端授权二维码链接(需要网站应用)
func (w *Wechat) GetQRCodeRedirect(redirect string, state string) string {
	u := url.Values{}
	u.Add("appid", w.AppID)
	u.Add("redirect_uri", redirect)
	u.Add("response_type", "code")
	u.Add("scope", "snsapi_login")
	u.Add("state", state)
	uri := "https://open.weixin.qq.com/connect/qrconnect?" + u.Encode() + "#wechat_redirect"
	return uri
}

//WxTokenData 用户的tokenData
type WxTokenData struct {
	AccessToken  string `json:"access_token"`  //接口调用凭证
	ExpiresIn    int    `json:"expires_in"`    //接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //用户刷新access_token
	Openid       string `json:"openid"`        //授权用户唯一标识
	Unionid      string `json:"unionid"`       //当且仅当该网站应用已获得该用户的userinfo授权时，才会出现该字段。
}

//GetTokenData 根据用户授权获取tokenData
func (w *Wechat) GetTokenData(code string) WxTokenData {
	u := url.Values{}
	u.Add("appid", w.AppID)
	u.Add("secret", w.AppSecret)
	u.Add("code", code)
	u.Add("grant_type", "authorization_code")
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?" + u.Encode()

	data := CurlGET(url)
	var token WxTokenData
	err := json.Unmarshal([]byte(data), &token)
	if err != nil {
		panic(err)
	}
	return token
}

//WxUserInfo 用户信息
type WxUserInfo struct {
	Openid        string `json:"openid"`         //普通用户的标识，对当前开发者帐号唯一
	Nickname      string `json:"nickname"`       //普通用户昵称
	Sex           int    `json:"sex"`            //普通用户性别，1为男性，2为女性
	Province      string `json:"province"`       //普通用户个人资料填写的省份
	City          string `json:"city"`           //普通用户个人资料填写的城市
	Country       string `json:"country"`        //国家，如中国为CN
	Headimgurl    string `json:"headimgurl"`     //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
	SubscribeTime int    `json:"subscribe_time"` //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Subscribe     int    `json:"subscribe"`      //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Unionid       string `json:"unionid"`        //用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
}

//UserInfoFromToken 根据授权的tokenData获取用户信息
func (w *Wechat) UserInfoFromToken(token string, openid string) WxUserInfo {
	u := url.Values{}
	u.Add("access_token", token)
	u.Add("openid", openid)
	url := "https://api.weixin.qq.com/sns/userinfo?" + u.Encode()
	data := CurlGET(url)
	var user WxUserInfo
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserInfoFromToken==", user)
	}
	return user
}

//GetUserInfo 获取已授权的用户信息
func (w *Wechat) GetUserInfo(openid string) WxUserInfo {
	u := url.Values{}
	u.Add("access_token", w.AccessToken())
	u.Add("openid", openid)
	url := "https://api.weixin.qq.com/cgi-bin/user/info?" + u.Encode()
	data := CurlGET(url)
	var user WxUserInfo
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetUserInfo==", user)
	}
	return user
}

//CustomText 发送客服文字消息
func (w *Wechat) CustomText(openid string, msg string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","msgtype":"text","text":{"content":"%s"}}`, openid, msg)
	return CurlJSON(url, data)
}

//CustomArticle 发送客服图文消息
func (w *Wechat) CustomArticle(openid string, title string, desc string, link string, pic string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","msgtype":"news","news":{"articles":[{"title":"%s","description":"%s","url":"%s","picurl":"%s"}]}}`, openid, title, desc, link, pic)
	return CurlJSON(url, data)
}

//CustomNews 发送客服多图文消息
func (w *Wechat) CustomNews(openid string, news string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","msgtype":"news","news":{"articles":%s}}`, openid, news)
	return CurlJSON(url, data)
}

//CustomMpNews 发送图文素材
func (w *Wechat) CustomMpNews(openid string, mediaID string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","msgtype":"mpnews","mpnews":{"media_id":"%s"}}`, openid, mediaID)
	return CurlJSON(url, data)
}

//SendTmp 发送模板消息
func (w *Wechat) SendTmp(openid string, tmpID string, link string, tmpData string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","template_id":"%s","url":"%s","data":"%s"}`, openid, tmpID, link, tmpData)
	return CurlJSON(url, data)
}

//SendTmpWithMiniProgram 发送带小程序链接的模板消息
func (w *Wechat) SendTmpWithMiniProgram(openid string, tmpID string, link string, tmpData string, miniprogram string) string {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"touser":"%s","template_id":"%s","url":"%s","data":"%s","miniprogram":"%s"}`, openid, tmpID, link, tmpData, miniprogram)
	return CurlJSON(url, data)
}

//GetTempID 获取模板ID
func (w *Wechat) GetTempID(tempCode string) string {
	url := "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"template_id_short":"%s"}`, tempCode)
	return CurlJSON(url, data)
}

//DelTempByID 删除模板消息
func (w *Wechat) DelTempByID(tempID string) string {
	url := "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"template_id":"%s"}`, tempID)
	return CurlJSON(url, data)
}

//GetAllTemp 获取模板列表
func (w *Wechat) GetAllTemp() string {
	url := "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=" + w.AccessToken()
	return CurlGET(url)
}

//WxQrcode 微信二维码
type WxQrcode struct {
	Ticket        string `json:"ticket"`
	URL           string `json:"url"`
	ExpireSeconds string `json:"expire_seconds"`
}

//GetQrcode 临时二维码
func (w *Wechat) GetQrcode(str string) WxQrcode {
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"expire_seconds":604800,"action_name":"QR_STR_SCENE","action_info":{"scene": {"scene_str": "%s"}}`, str)
	qrcodeStr := CurlJSON(url, data)
	var qrcode WxQrcode
	err := json.Unmarshal([]byte(qrcodeStr), &qrcode)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetQrcode==", qrcode)
	}
	return qrcode
}

//GetLimitQrcode 永久二维码
func (w *Wechat) GetLimitQrcode(str string) WxQrcode {
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + w.AccessToken()
	data := fmt.Sprintf(`{"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "%s"}}}`, str)
	qrcodeStr := CurlJSON(url, data)
	var qrcode WxQrcode
	err := json.Unmarshal([]byte(qrcodeStr), &qrcode)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetLimitQrcode==", qrcode)
	}
	return qrcode
}

//WxToken 微信Token
type WxToken struct {
	AccessToken string `json:"access_token"` //接口调用凭证
	ExpiresIn   int    `json:"expires_in"`   //接口调用凭证超时时间，单位（秒）
}

//GetWxToken 获取全局唯一接口调用凭据
func (w *Wechat) GetWxToken() WxToken {
	url := " https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + w.AppID + "&secret=" + w.AppSecret
	tokenStr := CurlGET(url)
	var token WxToken
	err := json.Unmarshal([]byte(tokenStr), &token)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetWxToken==", token)
	}
	return token
}
