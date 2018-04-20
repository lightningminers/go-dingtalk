package dingtalkTest

import (
	"testing"
)

func Test_GetCompanyAccessToken(t *testing.T) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	if c.AccessToken != "" {
		t.Log("测试获取access_token通过")
	} else {
		t.Error("测试获取access_token未通过")
	}
}
