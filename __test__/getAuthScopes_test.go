package dingtalkTest

import (
	"testing"
)

func Test_GetAuthScopes(t *testing.T) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	data, err := c.GetAuthScopes()
	if err != nil {
		t.Error("测试获取Auth Scopes 未通过", err)
	} else {
		t.Log("测试获取Auth Scopes 通过", data)
	}
}
