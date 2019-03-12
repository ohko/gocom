package gocom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type accessTokenSt struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type msgSt struct {
	Touser  string    `json:"touser"`
	Msgtype string    `json:"msgtype"`
	Agentid string    `json:"agentid"`
	Text    contentSt `json:"text"`
	Safe    int       `json:"safe"`
}
type contentSt struct {
	Content string `json:"content"`
}

// {"errcode":0,"errmsg":"ok","invaliduser":""}
type msgResult struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	InvalidUser string `json:"invaliduser"`
}

// WeiXinMsg ...
func WeiXinMsg(CORPID, CORPSECRET, Touser, Agentid, msg string) (err error) {
	go func() { recover() }()

	var (
		_url string
		res1 *http.Response
		res2 *http.Response
		acc  accessTokenSt
		ret  msgResult
	)

	// request access_token
	_url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CORPID, CORPSECRET)
	res1, err = http.Get(_url)
	if err != nil {
		return
	}
	defer res1.Body.Close()

	// decode access_token
	if err = json.NewDecoder(res1.Body).Decode(&acc); err != nil {
		return
	}

	// post data
	_url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", acc.AccessToken)
	data, err := json.Marshal(msgSt{
		Touser:  Touser,
		Msgtype: "text",
		Agentid: Agentid,
		Text:    contentSt{Content: msg},
		Safe:    0,
	})
	if err != nil {
		return
	}

	res2, err = http.Post(_url, "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer res2.Body.Close()

	// decode result message
	if err = json.NewDecoder(res2.Body).Decode(&ret); err != nil {
		return
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf(`{"errcode":%d,"errmsg":"%s","invaliduser":"%s"}`,
			ret.ErrCode, ret.ErrMsg, ret.InvalidUser)
	}

	return
}
