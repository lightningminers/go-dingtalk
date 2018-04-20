package dingtalk

import (
	"net/url"
)

type ChatCreateResponse struct {
	OpenAPIResponse
	ChatId string `json:"chatid"`
}

type ChatCreateRequest struct {
	Name            string   `json:"name"`
	Owner           string   `json:"owner"`
	UserIdList      []string `json:"useridlist"`
	ShowHistoryType int      `json:"showHistoryType,omitempty"`
}

type ChatUpdateResponse struct {
	OpenAPIResponse
}

type ChatUpdateRequest struct {
	ChatId        string   `json:"chatid"`
	Name          string   `json:"name,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	AddUserIdList []string `json:"add_useridlist,omitempty"`
	DelUserIdList []string `json:"del_useridlist,omitempty"`
}

type ChatGetResponse struct {
	OpenAPIResponse
	ChatInfo ChatGetInfo `json:"chat_info"`
}

type ChatGetInfo struct {
	Name       string
	Owner      string
	UserIdList []string `json:"useridlist"`
}

type ChatSendResponse struct {
	OpenAPIResponse
}

type ChatSendRequest struct {
	ChatId  string `json:"chatid"`
	MsgType string `json:"msgtype"`
}

type ChatSendActionCardRequest struct {
	ChatSendRequest
	ActionCard *ChatSendActionCard `json:"action_card"`
}

type ChatSendActionCard struct {
	Title          string      `json:"title"`
	Markdown       string      `json:"markdown"`
	SingleTitle    string      `json:"single_title,omitempty"`
	SingleUrl      string      `json:"single_url,omitempty"`
	BtnOrientation string      `json:"btn_orientation,omitempty"`
	BtnJSONList    interface{} `json:"btn_json_list,omitempty"`
	AgentId        string      `json:"agentid,omitempty"`
}

type ChatSendTextRequest struct {
	ChatSendRequest
	Text *ChatSendText `json:"text"`
}

type ChatSendText struct {
	Content string `json:"content"`
}

type ChatSendImageRequest struct {
	ChatSendRequest
	Image *ChatSendImage `json:"image"`
}

type ChatSendImage struct {
	MediaId string `json:"media_id"`
}

type ChatSendVoiceRequest struct {
	ChatSendRequest
	Voice *ChatSendVoice `json:"voice"`
}

type ChatSendVoice struct {
	MediaId  string `json:"media_id"`
	Duration string `json:"duration"`
}

type ChatSendFileRequest struct {
	ChatSendRequest
	File *ChatSendFile `json:"file"`
}

type ChatSendFile struct {
	MediaId string `json:"media_id"`
}

type ChatSendLinkRequest struct {
	ChatSendRequest
	Link *ChatSendLink `json:"link"`
}

type ChatSendLink struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicUrl     string `json:"pic_url"`
	MessageUrl string `json:"message_url"`
}

type ChatSendOARequest struct {
	ChatSendRequest
	Oa *ChatSendOA `json:"oa"`
}

type ChatSendOA struct {
	MessageUrl   string      `json:"message_url"`
	PcMessageUrl string      `json:"pc_message_url,omitempty"`
	Head         interface{} `json:"head"`
	Body         interface{} `json:"body"`
}

type ChatSendMarkdownRequest struct {
	ChatSendRequest
	Markdown *ChatSendMarkdown `json:"markdown"`
}

type ChatSendMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// 创建会话
func (dtc *DingTalkClient) ChatCreate(info *ChatCreateRequest) (ChatCreateResponse, error) {
	var data ChatCreateResponse
	err := dtc.httpRPC("chat/create", nil, info, &data)
	return data, err
}

// 修改会话
func (dtc *DingTalkClient) ChatUpdate(info *ChatUpdateRequest) (ChatUpdateResponse, error) {
	var data ChatUpdateResponse
	err := dtc.httpRPC("chat/update", nil, info, &data)
	return data, err
}

// 获取会话
func (dtc *DingTalkClient) ChatGet(chatId string) (ChatGetResponse, error) {
	var data ChatGetResponse
	params := url.Values{}
	params.Add("chatid", chatId)
	err := dtc.httpRPC("chat/get", params, nil, &data)
	return data, err
}

// 发送群消息
func (dtc *DingTalkClient) ChatSend(info interface{}) (ChatSendResponse, error) {
	var data ChatSendResponse
	err := dtc.httpRPC("chat/send", nil, info, &data)
	return data, err
}
