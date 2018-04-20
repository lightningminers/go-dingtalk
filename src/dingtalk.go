package dingtalk

import (
	"net/http"
	"time"
)

/*
*	date: 2018/05/20
*	version: 0.1
*	author: xiangwenwen(icepy)
*	description: DingTalk Golang SDK https://github.com/icepy
*
*	^_^ 想了很久，还是准备用中文写一些话，以后的日子打算做一个山野隐居的佛系程序员
*
* 平静的生活
* 平凡的人生
* 心中的自由
*
*	我们经营了一家很小的团队，五个人，曾经都来自大家常说的BAT
*
* 钉钉是我曾经工作过的地方，留下了很多回忆
*
*	企业服务市场对我而言，就像人每天要吃的饭
*
* 我们很乐意将我们的专业知识，服务于一些企业
*
* 如果你的公司有企业定制，技术咨询，技术培训等需求，不妨联系我们（钉钉搜索群号“21794502”）
 */

type DingTalkClient struct {
	CompanyConfig       *DTCompanyConfig
	MiniConfig          *DTMiniConfig
	IsvConfig           *DTIsvConfig
	TopConfig           *TopConfig
	HTTPClient          *http.Client
	AccessToken         string
	SSOAccessToken      string
	SNSAccessToken      string
	AccessTokenCache    Cache
	TicketCache         Cache
	SSOAccessTokenCache Cache
	SNSAccessTokenCache Cache
}

type TopConfig struct {
	TopFormat     string // json xml byte
	TopV          string
	TopSimplify   bool
	TopSecret     string
	TopSignMethod string
}

type DTMiniConfig struct {
	TopConfig
}

type DTIsvConfig struct {
	TopConfig
}

type DTCompanyConfig struct {
	TopConfig
	CorpID        string
	CorpSecret    string
	AgentID       string
	ChannelSecret string
	SSOSecret     string
	SNSAppID      string
	SNSSecret     string
}

func NewDingTalkCompanyClient(config *DTCompanyConfig) *DingTalkClient {
	c := &DingTalkClient{
		CompanyConfig: &DTCompanyConfig{},
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		TopConfig: &TopConfig{
			TopFormat:     topFormat,
			TopSecret:     topSecret,
			TopSignMethod: topSignMethod,
			TopSimplify:   topSimplify,
			TopV:          topV,
		},
		AccessTokenCache:    NewFileCache(".company_access_token_file"),
		TicketCache:         NewFileCache(".company_ticket_file"),
		SSOAccessTokenCache: NewFileCache(".company_sso_acess_token_file"),
		SNSAccessTokenCache: NewFileCache(".company_sns_access_token_file"),
	}
	if config != nil {
		if config.TopFormat != "" {
			c.TopConfig.TopFormat = config.TopFormat
		}
		if config.TopV != "" {
			c.TopConfig.TopV = config.TopV
		}
		if config.TopSecret != "" {
			c.TopConfig.TopSecret = config.TopSecret
		}
		if config.TopSignMethod != "" {
			c.TopConfig.TopSignMethod = config.TopSignMethod
		}
		if config.TopSimplify {
			c.TopConfig.TopSimplify = config.TopSimplify
		}
		c.CompanyConfig.CorpID = config.CorpID
		c.CompanyConfig.AgentID = config.AgentID
		c.CompanyConfig.CorpSecret = config.CorpSecret
		c.CompanyConfig.SSOSecret = config.SSOSecret
	}
	return c
}

func NewDingTalkISVClient() {

}

func NewDingTalkMiniClient() {

}
