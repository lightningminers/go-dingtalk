package dingtalkTest

import (
	"testing"
)

func Test_GetCompanyTicket(t *testing.T) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	ticket, err := c.GetJSAPITicket()
	if err != nil {
		t.Error("测试未能获取JSAPI Ticket")
	} else {
		t.Log("测试获取JSAPI Ticket通过", ticket)
	}
}
