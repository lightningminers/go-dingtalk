package dingtalkTest

import (
	"os"
	"testing"
)

func Test_MediaUpload(t *testing.T) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	o, ok := os.Open("wow.jpg")
	if ok == nil {
		data, err := c.MediaUpload("image", "wow.jpg", o)
		if err != nil {
			t.Error("测试图片上传未通过", err)
		} else {
			if data.MediaID != "" {
				t.Log("测试图片上传通过", data)
			} else {
				t.Error("测试图片上传未能获取到media_id")
			}
		}
	} else {
		t.Error("os.Open文件错误", ok)
	}

}
