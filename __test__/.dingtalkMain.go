package main

import (
	"os"
)

func main() {
	c := getCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
}

func getCompanyDingTalkClient() *dingtalk.DingTalkClient {
	CorpID := os.Getenv("CorpId")
	CorpSecret := os.Getenv("CorpSecret")
	config := &dingtalk.DTCompanyConfig{
		CorpID:     CorpID,
		CorpSecret: CorpSecret,
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}
