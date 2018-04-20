package dingtalkTest

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Test_MediaDownloadFile(t *testing.T) {
	mediaID := "@lADPBY0V4zNROQfNBUbNCWA" //填写你刚刚上传的mediaID
	name := randStringBytesRmndr()
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	ce, ok := os.Create(name + ".jpg")
	if ok == nil {
		err := c.MediaDownloadFile(mediaID, ce)
		if err == nil {
			t.Log("测试下载图片通过")
		} else {
			t.Error("测试下载图片未通过", err)
		}
	} else {
		t.Error("创建图片未通过", ok)
	}
}

func randStringBytesRmndr() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "crazyof.me")
	io.WriteString(h, t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}
