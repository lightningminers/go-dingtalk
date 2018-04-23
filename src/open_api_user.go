package dingtalk

import (
	"fmt"
	"net/url"
)

type UserIdResponse struct {
	OpenAPIResponse
	UserID   string `json:"userid"`
	DeviceID string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}

type UserIdByUnionIdResponse struct {
	OpenAPIResponse
	UserID      string `json:"userid"`
	ContactType int    `json:"contactType"`
}

type UserInfoResponse struct {
	OpenAPIResponse
	UserID          string `json:"userid"`
	OpenID          string `json:"openid"`
	Name            string
	Tel             string
	WorkPlace       string
	Remark          string
	Mobile          string
	Email           string
	OrgEmail        string
	Active          bool
	IsAdmin         bool
	IsBoos          bool
	DingID          string
	UnionID         string
	IsHide          bool
	Department      []int
	Position        string
	Avatar          string
	Jobnumber       string
	IsSenior        bool
	StateCode       string
	OrderInDepts    string
	IsLeaderInDepts string
	Extattr         interface{}
	Roles           []Roles
}

type Roles struct {
	ID        int `json:"id"`
	Name      string
	GroupName string
}

type UserSimpleListResponse struct {
	OpenAPIResponse
	HasMore  bool
	UserList []USimpleList
}

type USimpleList struct {
	UserID string
	Name   string
}

type UserListResponse struct {
	OpenAPIResponse
	HasMore  bool
	UserList []UDetailedList
}

type UDetailedList struct {
	UserID     string `json:"userid"`
	Order      int
	DingID     string
	UnionID    string
	Mobile     string
	Tel        string
	WorkPlace  string
	Remark     string
	IsAdmin    bool
	IsBoss     bool
	IsHide     bool
	IsLeader   bool
	Name       string
	Active     bool
	Department []int
	Position   string
	Email      string
	Avatar     string
	Jobnumber  string
	Extattr    interface{}
}

type UserAdminListResponse struct {
	OpenAPIResponse
	AdminList []Admins
}

type Admins struct {
	UserID   string `json:"userid"`
	SysLevel int    `json:"sys_level"`
}

type UserCanAccessMicroappResponse struct {
	OpenAPIResponse
	CanAccess bool
}

type UserCreateResponse struct {
	OpenAPIResponse
	UserID string
}

type UserCreateRequest struct {
	UserID       string      `json:"userid,omitempty"`
	Name         string      `json:"name"`
	OrderInDepts string      `json:"orderInDepts,omitempty"`
	Department   []int       `json:"department"`
	Position     string      `json:"position,omitempty"`
	Mobile       string      `json:"mobile"`
	Tel          string      `json:"tel,omitempty"`
	WorkPlace    string      `json:"workPlace,omitempty"`
	Remark       string      `json:"remark,omitempty"`
	Email        string      `json:"email,omitempty"`
	OrgEmail     string      `json:"orgEmail,omitempty"`
	JobNumber    string      `json:"jobnumber,omitempty"`
	IsHide       bool        `json:"isHide,omitempty"`
	IsSenior     bool        `json:"isSenior,omitempty"`
	Extattr      interface{} `json:"extattr,omitempty"`
}

type UserUpdateResponse struct {
	OpenAPIResponse
}

type UserUpdateRequest struct {
	Lang         string      `json:"lang,omitempty"`
	UserID       string      `json:"userid"`
	Name         string      `json:"name"`
	OrderInDepts string      `json:"orderInDepts,omitempty"`
	Department   []int       `json:"department,omitempty"`
	Position     string      `json:"position,omitempty"`
	Mobile       string      `json:"mobile,omitempty"`
	Tel          string      `json:"tel,omitempty"`
	WorkPlace    string      `json:"workPlace,omitempty"`
	Remark       string      `json:"remark,omitempty"`
	Email        string      `json:"email,omitempty"`
	OrgEmail     string      `json:"orgEmail,omitempty"`
	JobNumber    string      `json:"jobnumber,omitempty"`
	IsHide       bool        `json:"isHide,omitempty"`
	IsSenior     bool        `json:"isSenior,omitempty"`
	Extattr      interface{} `json:"extattr,omitempty"`
}

type UserDeleteResponse struct {
	OpenAPIResponse
}

type UserBatchDeleteResponse struct {
	OpenAPIResponse
}

type UserGetOrgUserCountResponse struct {
	OpenAPIResponse
	Count int
}

// 通过Code换取userid
func (dtc *DingTalkClient) UserIdByCode(code string) (UserIdResponse, error) {
	var data UserIdResponse
	params := url.Values{}
	params.Add("code", code)
	err := dtc.httpRPC("user/getuserinfo", params, nil, &data)
	return data, err
}

// 通过UnionId获取UserId
func (dtc *DingTalkClient) UserIdByUnionId(unionID string) (UserIdByUnionIdResponse, error) {
	var data UserIdByUnionIdResponse
	params := url.Values{}
	params.Add("unionid", unionID)
	err := dtc.httpRPC("user/getUseridByUnionid", params, nil, &data)
	return data, err
}

// 通过userid 换取 用户详细信息
func (dtc *DingTalkClient) UserInfoByUserId(userID string, lang string, isvGetCompanyInfo *DTIsvGetCompanyInfo) (UserInfoResponse, error) {
	var data UserInfoResponse
	params := url.Values{}
	params.Add("lang", lang)
	params.Add("userid", userID)
	err := dtc.httpRPC("user/get", params, nil, &data, isvGetCompanyInfo)
	return data, err
}

// 获取部门成员（简化版）
func (dtc *DingTalkClient) UserSimpleList(departmentID int) (UserSimpleListResponse, error) {
	var data UserSimpleListResponse
	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	err := dtc.httpRPC("user/simplelist", params, nil, &data)
	return data, err
}

// 获取部门成员（详情版）
func (dtc *DingTalkClient) UserList(departmentID int) (UserListResponse, error) {
	var data UserListResponse
	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	err := dtc.httpRPC("user/list", params, nil, &data)
	return data, err
}

// 获取管理员列表
func (dtc *DingTalkClient) UserAdminList() (UserAdminListResponse, error) {
	var data UserAdminListResponse
	err := dtc.httpRPC("user/get_admin", nil, nil, &data)
	return data, err
}

// 获取管理员的微应用管理权限
func (dtc *DingTalkClient) UserCanAccessMicroapp(appID string, userID string) (UserCanAccessMicroappResponse, error) {
	var data UserCanAccessMicroappResponse
	params := url.Values{}
	params.Add("appId", appID)
	params.Add("userId", userID)
	err := dtc.httpRPC("user/can_access_microap", params, nil, &data)
	return data, err
}

// 创建成员
func (dtc *DingTalkClient) UserCreate(info *UserCreateRequest) (UserCreateResponse, error) {
	var data UserCreateResponse
	err := dtc.httpRPC("user/create", nil, info, &data)
	return data, err
}

// 更新成员
func (dtc *DingTalkClient) UserUpdate(info *UserUpdateRequest) (UserUpdateResponse, error) {
	var data UserUpdateResponse
	err := dtc.httpRPC("user/update", nil, info, &data)
	return data, err
}

// 删除成员
func (dtc *DingTalkClient) UserDelete(userID string) (UserDeleteResponse, error) {
	var data UserDeleteResponse
	params := url.Values{}
	params.Add("userid", userID)
	err := dtc.httpRPC("user/delete", params, nil, &data)
	return data, err
}

// 批量删除成员
func (dtc *DingTalkClient) UserBatchDelete(userIdList []string) (UserBatchDeleteResponse, error) {
	var data UserBatchDeleteResponse
	body := map[string][]string{
		"useridlist": userIdList,
	}
	err := dtc.httpRPC("user/batchdelete", nil, body, &data)
	return data, err
}

// 获取企业员工人数
func (dtc *DingTalkClient) UserGetOrgUserCount(onlyActive int) (UserGetOrgUserCountResponse, error) {
	var data UserGetOrgUserCountResponse
	params := url.Values{}
	params.Add("onlyActive", fmt.Sprintf("%d", onlyActive))
	err := dtc.httpRPC("user/get_org_user_count", params, nil, &data)
	return data, err
}
