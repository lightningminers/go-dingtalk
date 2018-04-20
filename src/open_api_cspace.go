package dingtalk

import (
	"fmt"
	"net/url"
)

type CspaceAddToSingleChatRequest struct {
	AgentID  string `json:"agent_id"`
	UserID   string `json:"userid"`
	MediaID  string `json:"media_id"`
	FileName string `json:"file_name"`
}

type CspaceAddToSingleChatResponse struct {
	OpenAPIResponse
}

type CspaceAddRequest struct {
	AgentID   string `json:"agent_id,omitempty"`
	Code      string `json:"code"`
	MediaID   string `json:"media_id"`
	SpaceID   string `json:"space_id"`
	FolderID  string `json:"folder_id"`
	Name      string `json:"name"`
	Overwrite bool   `json:"overwrite,omitempty"`
}

type CspaceAddResponse struct {
	OpenAPIResponse
	Dentry string
}

type CspaceGetCustomSpaceRequest struct {
	Domain  string `json:"domain"`
	AgentID string `json:"agent_id"`
}

type CspaceGetCustomSpaceResponse struct {
	OpenAPIResponse
	SpaceID string
}

type CspaceGrantCustomSpaceRequest struct {
	AgentID  string `json:"agent_id"`
	Domain   string `json:"domain"`
	IType    string `json:"type"`
	UserID   string `json:"userid"`
	Path     string `json:"path"`
	Fileids  string `json:"fileids"`
	Duration int64  `json:"Duration"`
}

type CspaceGrantCustomSpaceResponse struct {
	OpenAPIResponse
}

// 发送文件给指定用户
func (dtc *DingTalkClient) CspaceAddToSingleChat(info *CspaceAddToSingleChatRequest) (CspaceAddToSingleChatResponse, error) {
	/*
		这一块要么是文档没写好，要么就是有bug，目前处理的方式是URL即拼接，也发送byte形式的POST
	*/
	var data CspaceAddToSingleChatResponse
	params := url.Values{}
	params.Add("agent_id", info.AgentID)
	params.Add("userid", info.UserID)
	params.Add("media_id", info.MediaID)
	params.Add("file_name", info.FileName)
	err := dtc.httpRPC("cspace/add_to_single_chat", params, info, &data)
	return data, err
}

// 新增文件到用户钉盘
func (dtc *DingTalkClient) CspaceAdd(info *CspaceAddRequest) (CspaceAddResponse, error) {
	var data CspaceAddResponse
	params := url.Values{}
	if info.AgentID != "" {
		params.Add("agent_id", info.AgentID)
	}
	params.Add("code", info.Code)
	params.Add("media_id", info.MediaID)
	params.Add("space_id", info.SpaceID)
	params.Add("folder_id", info.FolderID)
	params.Add("name", info.Name)
	if info.Overwrite {
		params.Add("overwrite", fmt.Sprintf("%s", info.Overwrite))
	}
	err := dtc.httpRPC("cspace/add", params, nil, &data)
	return data, err
}

// 获取企业下的自定义空间
func (dtc *DingTalkClient) CspaceGetCustomSpace(info *CspaceGetCustomSpaceRequest) (CspaceGetCustomSpaceResponse, error) {
	var data CspaceGetCustomSpaceResponse
	params := url.Values{}
	if info.AgentID != "" {
		params.Add("agent_id", info.AgentID)
	}
	if info.Domain != "" {
		params.Add("domain", info.Domain)
	}
	err := dtc.httpRPC("cspace/get_custom_space", params, nil, &data)
	return data, err
}

// 授权用户访问企业下的自定义空间
func (dtc *DingTalkClient) CspaceGrantCustomSpace(info *CspaceGrantCustomSpaceRequest) (CspaceGrantCustomSpaceResponse, error) {
	var data CspaceGrantCustomSpaceResponse
	params := url.Values{}
	if info.AgentID != "" {
		params.Add("agent_id", info.AgentID)
	}
	if info.Domain != "" {
		params.Add("domain", info.Domain)
	}
	params.Add("type", info.IType)
	params.Add("userid", info.UserID)
	if info.Path != "" {
		params.Add("path", info.Path)
	}
	if info.Fileids != "" {
		params.Add("fileids", info.Fileids)
	}
	params.Add("duration", fmt.Sprintf("%d", info.Duration))
	err := dtc.httpRPC("cspace/grant_custom_space", params, nil, &data)
	return data, err
}
