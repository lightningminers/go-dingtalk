package dingtalk

import (
	"fmt"
	"net/url"
)

type MicroAppCreateResponse struct {
	OpenAPIResponse
	AgentId int `json:"agentId"`
}

type MicroAppCreateRequest struct {
	AppIcon       string `json:"appIcon"`
	AppName       string `json:"appName"`
	AppDesc       string `json:"appDesc"`
	HomePageUrl   string `json:"homepageUrl"`
	PcHomePageUrl string `json:"pcHomepageUrl,omitempty"`
	OmpLink       string `json:"ompLink,omitempty"`
}

type MicroAppUpdateResponse struct {
	OpenAPIResponse
	AgentId int `json:"agentId"`
}

type MicroAppUpdateRequest struct {
	AppIcon       string `json:"appIcon,omitempty"`
	AppName       string `json:"appName,omitempty"`
	AppDesc       string `json:"appDesc,omitempty"`
	HomePageUrl   string `json:"homepageUrl,omitempty"`
	PcHomePageUrl string `json:"pcHomepageUrl,omitempty"`
	OmpLink       string `json:"ompLink,omitempty"`
	AgentId       int    `json:"agentId"`
}

type MicroAppDeleteResponse struct {
	OpenAPIResponse
}

type MicroAppListResponse struct {
	OpenAPIResponse
	AppList []MALBUIAppList
}

type MicroAppListByUserIdResponse struct {
	OpenAPIResponse
	AppList []MALBUIAppList
}

type MALBUIAppList struct {
	AppIcon       string `json:"appIcon"`
	AgentId       int    `json:"agentId"`
	AppDesc       string `json:"appDesc"`
	IsSelf        bool   `json:"isSelf"`
	Name          string `json:"name"`
	HomePageUrl   string `json:"homepageUrl"`
	PcHomePageUrl string `json:"pcHomepageUrl"`
	AppStatus     int    `json:"appStatus"`
	OmpLink       string `json:"ompLink"`
}

type MicroAppVisibleScopesResponse struct {
	OpenAPIResponse
	IsHidden          bool     `json:"isHidden"`
	DeptVisibleScopes []int    `json:"deptVisibleScopes"`
	UserVisibleScopes []string `json:"userVisibleScopes"`
}

type MicroAppSetVisibleScopesResponse struct {
	OpenAPIResponse
}

type MicroAppSetVisibleScopesRequest struct {
	AgentId           int      `json:"agentId"`
	IsHiddent         bool     `json:"isHiddent,omitempty"`
	DeptVisibleScopes []int    `json:"deptVisibleScopes,omitempty"`
	UserVisibleScopes []string `json:"userVisibleScopes,omitempty"`
}

type MicroAppRuleGetRuleListResponse struct {
	OpenAPIResponse
	RuleIdList []int `json:"ruleIdList"`
}

type MicroAppRuleGetUserTotaResponse struct {
	OpenAPIResponse
	Result []MicroAppRGUTResult
}

type MicroAppRGUTResult struct {
	RuleId    int
	UserTotal int
}

type MicroAppRuleDeleteResponse struct {
	OpenAPIResponse
}

// 创建微应用
func (dtc *DingTalkClient) MicroAppCreate(info *MicroAppCreateRequest) (MicroAppCreateResponse, error) {
	var data MicroAppCreateResponse
	err := dtc.httpRPC("microapp/create", nil, info, &data)
	return data, err
}

// 更新微应用
func (dtc *DingTalkClient) MicroAppUpdate(info *MicroAppUpdateRequest) (MicroAppUpdateResponse, error) {
	var data MicroAppUpdateResponse
	err := dtc.httpRPC("microapp/update", nil, info, &data)
	return data, err
}

// 删除微应用
func (dtc *DingTalkClient) MicroAppDelete(agentId int) (MicroAppDeleteResponse, error) {
	var data MicroAppDeleteResponse
	params := url.Values{}
	params.Add("agentId", fmt.Sprintf("%s", agentId))
	err := dtc.httpRPC("microapp/delete", params, nil, &data)
	return data, err
}

// 列出微应用
func (dtc *DingTalkClient) MicroAppList() (MicroAppListResponse, error) {
	var data MicroAppListResponse
	err := dtc.httpRPC("microapp/list", nil, map[string]string{}, &data)
	return data, err
}

// 列出员工可见的微应用
func (dtc *DingTalkClient) MicroAppListByUserId(userId string) (MicroAppListByUserIdResponse, error) {
	var data MicroAppListByUserIdResponse
	params := url.Values{}
	params.Add("userid", userId)
	err := dtc.httpRPC("microapp/list_by_userid", params, nil, &data)
	return data, err
}

// 获取企业设置的微应用可见范围
func (dtc *DingTalkClient) MicroAppVisibleScopes(agentId int) (MicroAppVisibleScopesResponse, error) {
	var data MicroAppVisibleScopesResponse
	params := url.Values{}
	params.Add("agentId", fmt.Sprintf("%s", agentId))
	err := dtc.httpRPC("microapp/visible_scopes", params, nil, &data)
	return data, err
}

// 设置微应用的可见范围
func (dtc *DingTalkClient) MicroAppSetVisibleScopes(info *MicroAppSetVisibleScopesRequest) (MicroAppSetVisibleScopesResponse, error) {
	var data MicroAppSetVisibleScopesResponse
	err := dtc.httpRPC("microapp/set_visible_scopes", nil, info, &data)
	return data, err
}

// 获取指定微应用下指定用户绑定的全部规则（定向开放接口）
func (dtc *DingTalkClient) MicroAppRuleGetRuleList(userId string, agentId string) (MicroAppRuleGetRuleListResponse, error) {
	var data MicroAppRuleGetRuleListResponse
	err := dtc.httpRPC("microapp/rule/get_rule_list", nil, map[string]string{
		"userid":  userId,
		"agentId": agentId,
	}, &data)
	return data, err
}

// 获取规则绑定的用户数（定向开放接口）
func (dtc *DingTalkClient) MicroAppRuleGetUserTota(agentId int, ruleIdList []int) (MicroAppRuleGetUserTotaResponse, error) {
	var data MicroAppRuleGetUserTotaResponse
	err := dtc.httpRPC("microapp/rule/get_user_total", nil, map[string]interface{}{
		"agentId":    agentId,
		"ruleIdList": ruleIdList,
	}, &data)
	return data, err
}

// 删除规则（定向开放接口）
func (dtc *DingTalkClient) MicroAppRuleDelete(agentId int, ruleId int) (MicroAppRuleDeleteResponse, error) {
	var data MicroAppRuleDeleteResponse
	err := dtc.httpRPC("microapp/rule/delete", nil, map[string]int{
		"agentId": agentId,
		"ruleId":  ruleId,
	}, &data)
	return data, err
}
