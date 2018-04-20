package dingtalk

type MessageSendToConversationResponse struct {
	OpenAPIResponse
	Receiver string
}

type MessageSendRequest struct {
	Sender  string `json:"sender"`
	Cid     string `json:"cid"`
	MsgType string `json:"msgtype"`
}

type MessageSendActionCardRequest struct {
	MessageSendRequest
	ActionCard *MessageSendActionCard `json:"action_card"`
}

type MessageSendActionCard struct {
	Title          string      `json:"title"`
	Markdown       string      `json:"markdown"`
	SingleTitle    string      `json:"single_title,omitempty"`
	SingleUrl      string      `json:"single_url,omitempty"`
	BtnOrientation string      `json:"btn_orientation,omitempty"`
	BtnJSONList    interface{} `json:"btn_json_list,omitempty"`
	AgentId        string      `json:"agentid,omitempty"`
}

type MessageSendTextRequest struct {
	MessageSendRequest
	Text *MessageSendText `json:"text"`
}

type MessageSendText struct {
	Content string `json:"content"`
}

type MessageSendImageRequest struct {
	MessageSendRequest
	Image *MessageSendImage `json:"image"`
}

type MessageSendImage struct {
	MediaId string `json:"media_id"`
}

type MessageSendVoiceRequest struct {
	MessageSendRequest
	Voice *MessageSendVoice `json:"voice"`
}

type MessageSendVoice struct {
	MediaId  string `json:"media_id"`
	Duration string `json:"duration"`
}

type MessageSendFileRequest struct {
	MessageSendRequest
	File *MessageSendFile `json:"file"`
}

type MessageSendFile struct {
	MediaId string `json:"media_id"`
}

type MessageSendLinkRequest struct {
	MessageSendRequest
	Link *MessageSendLink `json:"link"`
}

type MessageSendLink struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicUrl     string `json:"pic_url"`
	MessageUrl string `json:"message_url"`
}

type MessageSendOARequest struct {
	MessageSendRequest
	Oa *MessageSendOA `json:"oa"`
}

type MessageSendOA struct {
	MessageUrl   string      `json:"message_url"`
	PcMessageUrl string      `json:"pc_message_url,omitempty"`
	Head         interface{} `json:"head"`
	Body         interface{} `json:"body"`
}

type MessageSendMarkdownRequest struct {
	MessageSendRequest
	Markdown *MessageSendMarkdown `json:"markdown"`
}

type MessageSendMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// 发送普通消息
func (dtc *DingTalkClient) MessageSendToConversation(info interface{}) (MessageSendToConversationResponse, error) {
	var data MessageSendToConversationResponse
	err := dtc.httpRPC("message/send_to_conversation", nil, info, &data)
	return data, err
}
