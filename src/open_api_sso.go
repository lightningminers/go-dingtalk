package dingtalk

import "net/url"

type SSOAdminInfoByCodeResponse struct {
	OpenAPIResponse
	CorpInfo SSOCorpInfo `json:"corp_info"`
	IsSys    bool        `json:"is_sys"`
}

type SSOCorpInfo struct {
	CorpName string `json:"corp_name"`
	CorpID   string `json:"corpid"`
}

type SSOUserInfo struct {
	Avatar string
	Email  string
	Name   string
	UserID string `json:"userid"`
}

// 通过CODE换取微应用管理员的身份信息
func (dtc *DingTalkClient) SSOAdminInfoByCode(code string) (SSOAdminInfoByCodeResponse, error) {
	var data SSOAdminInfoByCodeResponse
	params := url.Values{}
	params.Add("code", code)
	err := dtc.httpSSO("sso/getuserinfo", params, nil, &data)
	return data, err
}
