package helper

import (
	"bytes"
	"fmt"
	_const "go-gin/common/const"
	"log"
	"net/http"
)

// 企业微信
func WeChatNotify(text string, mentionAll bool) {
	mentionAllStr := ""
	if mentionAll {
		mentionAllStr = "@all"
	}
	dataJsonStr := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s", "mentioned_list": [%s]}}`, text, mentionAllStr)
	resp, err := http.Post(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+_const.WX_ERROR_LOG_KEY,
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}
