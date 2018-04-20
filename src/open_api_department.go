package dingtalk

import (
	"fmt"
	"net/url"
)

type SubDepartmentListResponse struct {
	OpenAPIResponse
	SubDeptIdList []int
}

type DepartmentListResponse struct {
	OpenAPIResponse
	Department []Department
}

type Department struct {
	Id              int
	Name            string
	ParentId        int
	CreateDeptGroup bool
	AutoAddUser     bool
}

type DepartmentDetailResponse struct {
	OpenAPIResponse
	Id                    int
	Name                  string
	Order                 int
	ParentId              int
	CreateDeptGroup       bool
	AutoAddUser           bool
	DeptHiding            bool
	DeptPerimits          string
	UserPerimits          string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUserIdList string
	SourceIdentifier      string
}

type DepartmentCreateResponse struct {
	OpenAPIResponse
	Id int
}

type DepartmentCreateRequest struct {
	Name             string `json:"name"`
	ParentId         string `json:"parentid"`
	Order            string `json:"order,omitempty"`
	CreateDeptGroup  bool   `json:"createDeptGroup,omitempty"`
	DeptHiding       bool   `json:"deptHiding,omitempty"`
	DeptPerimits     string `json:"deptPerimits,omitempty"`
	UserPerimits     string `json:"userPerimits,omitempty"`
	OuterDept        string `json:"outerDept,omitempty"`
	OuterPermitDepts string `json:"outerPermitDepts,omitempty"`
	OuterPermitUsers string `json:"outerPermitUsers,omitempty"`
	SourceIdentifier string `json:"sourceIdentifier,omitempty"`
}

type DepartmentUpdateResponse struct {
	OpenAPIResponse
	Id int
}

type DepartmentUpdateRequest struct {
	Lang                  string `json:"lang,omitempty"`
	Name                  string `json:"name,omitempty"`
	ParentId              string `json:"parentid,omitempty"`
	Order                 string `json:"order,omitempty"`
	Id                    string `json:"id"`
	CreateDeptGroup       bool   `json:"createDeptGroup,omitempty"`
	AutoAddUser           bool   `json:"autoAddUser,omitempty"`
	DeptManagerUseridList string `json:"deptManagerUseridList,omitempty"`
	DeptHiding            bool   `json:"deptHiding,omitempty"`
	DeptPerimits          string `json:"deptPerimits,omitempty"`
	UserPerimits          string `json:"userPerimits,omitempty"`
	OuterDept             string `json:"outerDept,omitempty"`
	OuterPermitDepts      string `json:"outerPermitDepts,omitempty"`
	OuterPermitUsers      string `json:"outerPermitUsers,omitempty"`
	OrgDeptOwner          string `json:"orgDeptOwner,omitempty"`
	SourceIdentifier      string `json:"sourceIdentifier,omitempty"`
}

type DepartmentDeleteResponse struct {
	OpenAPIResponse
}

type DepartmentListParentDeptsByDeptResponse struct {
	OpenAPIResponse
	ParentIds []int `json:"parentIds"`
}

type DepartmentListParentDeptsResponse struct {
	OpenAPIResponse
	ParentIds interface{} `json:"dep"`
}

// 获取子部门Id列表
func (dtc *DingTalkClient) SubDepartmentList(id interface{}) (SubDepartmentListResponse, error) {
	var data SubDepartmentListResponse
	params := url.Values{}
	if id != nil {
		if v, ok := id.(string); ok {
			params.Add("id", v)
		}
	}
	err := dtc.httpRPC("department/list_ids", params, nil, &data)
	return data, err
}

// 获取部门id列表
func (dtc *DingTalkClient) DepartmentList(id interface{}, lang interface{}) (DepartmentListResponse, error) {
	var data DepartmentListResponse
	params := url.Values{}
	if id != nil {
		if v, ok := id.(string); ok {
			params.Add("id", v)
		}
	}
	if lang != nil {
		if v, ok := lang.(string); ok {
			params.Add("lang", v)
		}
	}
	err := dtc.httpRPC("department/list", params, nil, &data)
	return data, err
}

// 获取部门详情
func (dtc *DingTalkClient) DepartmentDetail(id interface{}, lang interface{}) (DepartmentDetailResponse, error) {
	var data DepartmentDetailResponse
	params := url.Values{}
	if id != nil {
		if v, ok := id.(string); ok {
			params.Add("id", v)
		}
	}
	if lang != nil {
		if v, ok := lang.(string); ok {
			params.Add("lang", v)
		}
	}
	err := dtc.httpRPC("department/get", params, nil, &data)
	return data, err
}

// 创建部门
func (dtc *DingTalkClient) DepartmentCreate(info *DepartmentCreateRequest) (DepartmentCreateResponse, error) {
	var data DepartmentCreateResponse
	err := dtc.httpRPC("department/create", nil, info, &data)
	return data, err
}

// 更新部门
func (dtc *DingTalkClient) DepartmentUpdate(info *DepartmentUpdateRequest) (DepartmentUpdateResponse, error) {
	var data DepartmentUpdateResponse
	err := dtc.httpRPC("department/update", nil, info, &data)
	return data, err
}

// 删除部门
func (dtc *DingTalkClient) DepartmentDelete(id int) (DepartmentDeleteResponse, error) {
	var data DepartmentDeleteResponse
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := dtc.httpRPC("department/delete", params, nil, &data)
	return data, err
}

// 查询部门的所有上级父部门路径
func (dtc *DingTalkClient) DepartmentListParentDeptsByDept(id int) (DepartmentListParentDeptsByDeptResponse, error) {
	var data DepartmentListParentDeptsByDeptResponse
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := dtc.httpRPC("department/list_parent_depts_by_dept", params, nil, &data)
	return data, err
}

// 查询指定用户的所有上级父部门路径
func (dtc *DingTalkClient) DepartmentListParentDepts(userId string) (DepartmentListParentDeptsResponse, error) {
	var data DepartmentListParentDeptsResponse
	params := url.Values{}
	params.Add("userId", userId)
	err := dtc.httpRPC("department/list_parent_depts", params, nil, &data)
	return data, err
}
