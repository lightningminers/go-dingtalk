package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type AccessTokenResponse struct {
	OpenAPIResponse
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
	Created     int64
}

type SSOAccessTokenResponse struct {
	OpenAPIResponse
	SSOAccessToken string `json:"access_token"`
	Expires        int    `json:"expires_in"`
	Created        int64
}

type SNSAccessTokenResponse struct {
	OpenAPIResponse
	SNSAccessToken string `json:"access_token"`
	Expires        int    `json:"expires_in"`
	Created        int64
}

type SuiteAccessTokenResponse struct {
	OpenAPIResponse
	SuiteAccessToken string `json:"suite_access_token"`
	Expires          int    `json:"expires_in"`
	Created          int64
}

type TicketResponse struct {
	OpenAPIResponse
	Ticket  string `json:"ticket"`
	Expires int    `json:"expires_in"`
	Created int64
}

type ScopesResponse struct {
	OpenAPIResponse
	AuthUserField  []string
	ConditionField []string
	AuthOrgScopes
}

type AuthOrgScopes struct {
	AuthedDept []int
	AuthedUser []string
}

func (e *AccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *AccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *TicketResponse) CreatedAt() int64 {
	return e.Created
}

func (e *TicketResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SSOAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SSOAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SNSAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SNSAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

func (e *SuiteAccessTokenResponse) CreatedAt() int64 {
	return e.Created
}

func (e *SuiteAccessTokenResponse) ExpiresIn() int {
	return e.Expires
}

// 刷新企业获取的access_token
func (dtc *DingTalkClient) RefreshCompanyAccessToken() error {
	dtc.Locker.Lock()
	var data AccessTokenResponse
	err := dtc.AccessTokenCache.Get(&data)
	if err == nil {
		dtc.AccessToken = data.AccessToken
		fmt.Printf("Get access_token To Local Cache=%s\n", dtc.AccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	params := url.Values{}
	params.Add("corpid", dtc.DTConfig.CorpID)
	params.Add("corpsecret", dtc.DTConfig.CorpSecret)
	err = dtc.httpRPC("gettoken", params, nil, &data)
	if err == nil {
		dtc.AccessToken = data.AccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.AccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 刷新企业获取的sso_access_token
func (dtc *DingTalkClient) RefreshSSOAccessToken() error {
	dtc.Locker.Lock()
	var data SSOAccessTokenResponse
	err := dtc.SSOAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SSOAccessToken = data.SSOAccessToken
		fmt.Printf("Get sso_access_token To Local Cache=%s\n", dtc.SSOAccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	params := url.Values{}
	params.Add("corpid", dtc.DTConfig.CorpID)
	params.Add("corpsecret", dtc.DTConfig.SSOSecret)
	err = dtc.httpSSO("sso/gettoken", params, nil, &data)
	if err == nil {
		dtc.SSOAccessToken = data.SSOAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SSOAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 刷新 SNS access_token
func (dtc *DingTalkClient) RefreshSNSAccessToken() error {
	dtc.Locker.Lock()
	var data SNSAccessTokenResponse
	err := dtc.SNSAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SNSAccessToken = data.SNSAccessToken
		fmt.Printf("Get sns_access_token To Local Cache=%s\n", dtc.SNSAccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	params := url.Values{}
	params.Add("appid", dtc.DTConfig.SNSAppID)
	params.Add("appsecret", dtc.DTConfig.SNSSecret)
	err = dtc.httpSNS("sns/gettoken", params, nil, &data)
	if err == nil {
		dtc.SNSAccessToken = data.SNSAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SNSAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 刷新 isv suite_access_token
func (dtc *DingTalkClient) RefreshSuiteAccessToken() error {
	dtc.Locker.Lock()
	var data SuiteAccessTokenResponse
	err := dtc.SuiteAccessTokenCache.Get(&data)
	if err == nil {
		dtc.SuiteAccessToken = data.SuiteAccessToken
		fmt.Printf("Get suite_access_token To Local Cache=%s\n", dtc.SuiteAccessToken)
		dtc.Locker.Unlock()
		return nil
	}
	info := map[string]string{
		"suite_key":    dtc.DTConfig.SuiteKey,
		"suite_secret": dtc.DTConfig.SuiteSecret,
		"suite_ticket": dtc.DTConfig.SuiteTicket,
	}
	err = dtc.httpSNS("service/get_suite_token", nil, info, &data)
	if err == nil {
		dtc.SuiteAccessToken = data.SuiteAccessToken
		data.Expires = data.Expires | 7200
		data.Created = time.Now().Unix()
		err = dtc.SuiteAccessTokenCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return err
}

// 获取Ticket
func (dtc *DingTalkClient) GetJSAPITicket() (ticket string, err error) {
	dtc.Locker.Lock()
	var data TicketResponse
	err = dtc.TicketCache.Get(&data)
	if err == nil {
		dtc.Locker.Unlock()
		return data.Ticket, err
	}
	err = dtc.httpRPC("get_jsapi_ticket", nil, nil, &data)
	if err == nil {
		ticket = data.Ticket
		dtc.TicketCache.Set(&data)
		dtc.Locker.Unlock()
	}
	return ticket, err
}

// 配置config信息
func (dtc *DingTalkClient) GetConfig(nonceStr string, timestamp string, url string) string {
	var config map[string]string
	ticket, _ := dtc.GetJSAPITicket()
	config = map[string]string{
		"url":       url,
		"nonceStr":  nonceStr,
		"agentId":   dtc.DTConfig.AgentID,
		"timeStamp": timestamp,
		"corpId":    dtc.DTConfig.CorpID,
		"ticket":    ticket,
		"signature": sign(ticket, nonceStr, timestamp, url),
	}
	bytes, _ := json.Marshal(&config)
	return string(bytes)
}

// 签名
func sign(ticket string, nonceStr string, timeStamp string, url string) string {
	s := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonceStr, timeStamp, url)
	return sha1Sign(s)
}

// 获取授权范围
func (dtc *DingTalkClient) GetAuthScopes() (ScopesResponse, error) {
	var data ScopesResponse
	err := dtc.httpRPC("auth/scopes", nil, nil, &data)
	return data, err
}
