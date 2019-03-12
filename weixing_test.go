package com

import "testing"

func TestWeiXinMsg(t *testing.T) {
	CORPID := "wx3c261331b760809a"
	CORPSECRET := "fj00lPiYOztFjaz2kD_w2sAfWo1iwdTpnsEPHqwoLPs"
	Touser := "hk"       // 多个接收者用‘|’分隔
	Agentid := "1000002" // 企业应用的id
	msg := "hello"
	if err := WeiXinMsg(CORPID, CORPSECRET, Touser, Agentid, msg); err != nil {
		t.Error(err)
	}
}
