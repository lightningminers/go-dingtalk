package dingtalk

type CBCallBackRequest struct {
	CallbackTag []string `json:"call_back_tag"`
	Token       string   `json:"token"`
	AesKey      string   `json:"aes_key"`
	URL         string   `json:"url"`
}

type CBCallBackResponse struct {
	OpenAPIResponse
}

type CBQueryCallbackResponse struct {
	OpenAPIResponse
	CallbackTag []string `json:"call_back_tag"`
	Token       string   `json:"token"`
	AesKey      string   `json:"aes_key"`
	URL         string   `json:"url"`
}

type CBGetFailedCallbackResponse struct {
	OpenAPIResponse
	HasMore    bool              `json:"has_more"`
	FailedList []FailedCallbacks `json:"failed_list"`
}

type FailedCallbacks struct {
	EventTime   int      `json:"event_time"`
	CallbackTag string   `json:"call_back_tag"`
	UserID      []string `json:"userid"`
	CorpID      string   `json:"corpid"`
}

// 注册事件回调接口
func (dtc *DingTalkClient) CBRegisterCallback(info *CBCallBackRequest) (CBCallBackResponse, error) {
	var data CBCallBackResponse
	err := dtc.httpRPC("call_back/register_call_back", nil, info, &data)
	return data, err
}

// 查询事件回调接口
func (dtc *DingTalkClient) CBQueryCallback() (CBQueryCallbackResponse, error) {
	var data CBQueryCallbackResponse
	err := dtc.httpRPC("call_back/get_call_back", nil, nil, &data)
	return data, err
}

// 更新事件回调接口
func (dtc *DingTalkClient) CBUpdateCallback(info *CBCallBackRequest) (CBCallBackResponse, error) {
	var data CBCallBackResponse
	err := dtc.httpRPC("call_back/update_call_back", nil, info, &data)
	return data, err
}

// 删除事件回调接口
func (dtc *DingTalkClient) CBDeleteCallback() (CBCallBackResponse, error) {
	var data CBCallBackResponse
	err := dtc.httpRPC("call_back/delete_call_back", nil, nil, &data)
	return data, err
}

// 获取回调失败的结果
func (dtc *DingTalkClient) CBGetFailedCallbacks() (CBGetFailedCallbackResponse, error) {
	var data CBGetFailedCallbackResponse
	err := dtc.httpRPC("call_back/get_call_back_failed_result", nil, nil, &data)
	return data, err
}
