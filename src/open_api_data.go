package dingtalk

type DataRecordResponse struct {
	OpenAPIResponse
	ID string `json:"id"`
}

type DataRequest struct {
	ID          string      `json:"id,omitempty"`
	StartTimeMs string      `json:"startTimeMs"`
	EndTimeMs   string      `json:"endTimeMs"`
	Module      string      `json:"module,omitempty"`
	OriginID    string      `json:"originId,omitempty"`
	UserID      string      `json:"userid"`
	AgentID     string      `json:"agentId"`
	CallbackUrl string      `json:"callbackUrl"`
	Extension   interface{} `json:"extension,omitempty"`
}

type DataUpdateResponse struct {
	OpenAPIResponse
}

// 记录统计数据
func (dtc *DingTalkClient) DataRecord(info *DataRequest) (DataRecordResponse, error) {
	var data DataRecordResponse
	err := dtc.httpRPC("data/record", nil, info, &data)
	return data, err
}

// 更新统计数据
func (dtc *DingTalkClient) DataUpdate(info *DataRequest) (DataUpdateResponse, error) {
	var data DataUpdateResponse
	err := dtc.httpRPC("data/update", nil, info, &data)
	return data, err
}
