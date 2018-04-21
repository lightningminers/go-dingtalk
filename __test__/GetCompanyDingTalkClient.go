package dingtalkTest

import (
	"os"

	"../src"
)

func GetCompanyDingTalkClient() *dingtalk.DingTalkClient {
	CorpID := os.Getenv("CorpId")
	CorpSecret := os.Getenv("CorpSecret")
	AgentID := os.Getenv("AgentID")
	SSOSecret := os.Getenv("SSOSecret")
	SNSAppID := os.Getenv("SNSAppID")
	SNSSecret := os.Getenv("SNSSecret")
	config := &dingtalk.DTConfig{
		CorpID:     CorpID,
		CorpSecret: CorpSecret,
		AgentID:    AgentID,
		SSOSecret:  SSOSecret,
		SNSAppID:   SNSAppID,
		SNSSecret:  SNSSecret,
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}
