package gocom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// SendWeixin ...
func SendWeixin(CORPID, CORPSECRET, Touser, Agentid, sz string) {
	go func() { recover() }()
	// request access_token
	_url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CORPID, CORPSECRET)
	req, err := http.Get(_url)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// decode access_token
	var acc accessTokenSt
	if err := json.NewDecoder(req.Body).Decode(&acc); err != nil {
		log.Println(err.Error())
		return
	}

	// post data
	_url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", acc.AccessToken)
	var msg = msgSt{
		Touser:  Touser,
		Msgtype: "text",
		Agentid: Agentid,
		Text:    contentSt{Content: sz},
		Safe:    0,
	}
	data, _ := json.Marshal(msg)
	req, err = http.Post(_url, "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err.Error())
		return
	}
	r, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(r))
}
