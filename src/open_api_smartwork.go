package dingtalk

import (
	"fmt"
	"net/url"
)

type SmartworkAttendanceListRecordRequest struct {
	UserIds       []string `json:"userIds"`
	CheckDateFrom string   `json:"checkDateFrom"`
	CheckDateTo   string   `json:"checkDateTo"`
}

type SmartworkAttendanceListRecordResponse struct {
	OpenAPIResponse
	RecordResult []SmartworkALRRecordResult `json:"recordresult"`
}

type SmartworkALRRecordResult struct {
	GmtModified    int64   `json:"gmtModified"`
	IsLegal        string  `json:"isLegal"`
	BaseCheckTime  int64   `json:"baseCheckTime"`
	ID             int64   `json:"id"`
	UserAddress    string  `json:"userAddress"`
	UserID         string  `json:"userId"`
	CheckType      string  `json:"checkType"`
	TimeResult     string  `json:"timeResult"`
	DeviceID       string  `json:"deviceId"`
	CorpID         string  `json:"corpId"`
	SourceType     string  `json:"sourceType"`
	WorkDate       int64   `json:"workDate"`
	PlanCheckTime  int64   `json:"planCheckTime"`
	GmtCreate      int64   `json:"gmtCreate"`
	LocationMethod string  `json:"locationMethod"`
	LocationResult string  `json:"locationResult"`
	UserLongitude  float64 `json:"userLongitude"`
	PlanID         int64   `json:"planId"`
	GroupID        int64   `json:"groupId"`
	UserAccuracy   int     `json:"userAccuracy"`
	UserCheckTime  int64   `json:"userCheckTime"`
	UserLatitude   float64 `json:"userLatitude"`
	ProcInstId     string  `json:"procInstId"`
}

type SmartworkCheckinRecordRequest struct {
	DepartmentID string
	StartTime    int64
	EndTime      int64
	Offset       int
	Size         int
	Order        string
}

type SmartworkCheckinRecordResponse struct {
	OpenAPIResponse
	Data []SmartworkCheckinRecordData `json:"data"`
}

type SmartworkCheckinRecordData struct {
	Name        string   `json:"name"`
	UserID      string   `json:"userId"`
	Timestamp   int64    `json:"timestamp"`
	Avatar      string   `json:"avatar"`
	Place       string   `json:"place"`
	DetailPlace string   `json:"detailPlace"`
	Remark      string   `json:"remark"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	ImageList   []string `json:"imageList"`
}

// 考勤打卡记录开放
func (dtc *DingTalkClient) SmartworkAttendanceListRecord(info *SmartworkAttendanceListRecordRequest) (SmartworkAttendanceListRecordResponse, error) {
	var data SmartworkAttendanceListRecordResponse
	err := dtc.httpRPC("attendance/listRecord", nil, info, &data)
	return data, err
}

// 获得签到数据
func (dtc *DingTalkClient) SmartworkCheckinRecord(info *SmartworkCheckinRecordRequest) (SmartworkCheckinRecordResponse, error) {
	var data SmartworkCheckinRecordResponse
	params := url.Values{}
	params.Add("department_id", info.DepartmentID)
	params.Add("start_time", fmt.Sprintf("%d", info.StartTime))
	params.Add("end_time", fmt.Sprintf("%d", info.EndTime))
	if info.Offset >= 0 {
		params.Add("offset", fmt.Sprintf("%d", info.Offset))
	}
	if info.Size > 0 {
		params.Add("size", fmt.Sprintf("%d", info.Size))
	}
	if info.Order != "" {
		params.Add("order", info.Order)
	}
	err := dtc.httpRPC("checkin/record", params, nil, &data)
	return data, err
}
