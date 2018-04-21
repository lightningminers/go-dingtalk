package dingtalk

type TopErrorResponse struct {
	SubMsg  string `json:"sub_msg"`
	Code    int    `json:"code"`
	SubCode string `json:"sub_code"`
	Msg     string `json:"msg"`
}

// start --- 查询用户是否开启了钉钉运动

type TopCorpHealthStepinfoGetUserStatusResponse struct {
	ErrorResponse                                   TopErrorResponse  `json:"error_response"`
	DingTalkCorpHealthStepinfoGetuserstatusResponse TopCHSGUSResponse `json:"dingtalk_corp_health_stepinfo_getuserstatus_response"`
}

type TopCHSGUSResponse struct {
	Result    TopCHSGUSResult
	RequestId string `json:"request_id"`
}

type TopCHSGUSResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	Status          bool
}

// end --- 查询用户是否开启了钉钉运动

// start --- 批量查询多个用户的钉钉运动步数

type TopCorpHealthStepinfoListByUseridResponse struct {
	ErrorResponse                                  TopErrorResponse  `json:"error_response"`
	DingtalkCorpHealthStepinfoListByUseridResponse TopCHSLBUResponse `json:"dingtalk_corp_health_stepinfo_listbyuserid_response"`
}

type TopCHSLBUResponse struct {
	Result    TopCHSLBUResult
	RequestId string `json:"request_id"`
}

type TopCHSLBUResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	StepInfoList    interface{} `json:"stepinfo_list"`
}

// end --- 批量查询多个用户的钉钉运动步数

// start --- 获取角色的员工列表

type TopCorpRoleSimpleListResponse struct {
	ErrorResponse                      TopErrorResponse `json:"error_response"`
	DingtalkCorpRoleSimpleListResponse TopCRSLResponse  `json:"dingtalk_corp_role_simplelist_response"`
}

type TopCRSLResponse struct {
	Result    TopCRSLResult
	RequestId string `json:"request_id"`
}

type TopCRSLResult struct {
	HasMore bool `json:"has_more"`
	List    interface{}
}

// end --- 获取角色的员工列表

// start --- 企业会话消息异步发送

type TopCorpMessageCorpconversationAsyncsendRequest struct {
	MsgType    string
	AgentId    int
	UserIdList []string
	DeptIdList []int
	ToAllUser  bool
	Msgcontent interface{}
}

type TopCorpMessageCorpconversationAsyncsendResponse struct {
	ErrorResponse                                        TopErrorResponse `json:"error_response"`
	DingtalkCorpMessageCorpconversationAsyncsendResponse TopCMCAResponse  `json:"dingtalk_corp_message_corpconversation_asyncsend_response"`
}

type TopCMCAResponse struct {
	Result    TopCMCAResult
	RequestId string `json:"request_id"`
}

type TopCMCAResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	TaskId          int `json:"task_id"`
}

// end --- 企业会话消息异步发送

// start --- 通过用户授权码异步向企业会话发送消息

type TopCorpMessageCorpconversationAsyncsendbycodeRequest struct {
	MsgType    string
	AgentId    int
	UserIdList []string
	DeptIdList []int
	ToAllUser  bool
	Msgcontent interface{}
	Code       string
}

type TopCorpMessageCorpconversationAsyncsendbycodeResponse struct {
	ErrorResponse                                              TopErrorResponse    `json:"error_response"`
	DingtalkCorpMessageCorpconversationAsyncsendbycodeResponse TopCMCACodeResponse `json:"dingtalk_corp_message_corpconversation_asyncsendbycode_response"`
}

type TopCMCACodeResponse struct {
	Result TopCMCACodeResult
}

type TopCMCACodeResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	TaskId          int `json:"task_id"`
}

// end --- 通过用户授权码异步向企业会话发送消息

// start --- 获取异步发送企业会话消息的发送进度

type TopCorpMessageCorpconversationGetsendprogressResponse struct {
	ErrorResponse                                              TopErrorResponse `json:"error_response"`
	DingtalkCorpMessageCorpconversationGetsendprogressResponse TopCMCGRResponse `json:"dingtalk_corp_message_corpconversation_getsendprogress_response"`
}

type TopCMCGRResponse struct {
	Result TopCMCGRResult
}

type TopCMCGRResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	Progress        TopCMCGRProgress `json:"progress"`
}

type TopCMCGRProgress struct {
	ProgressInPercent int `json:"progress_in_percent"`
	Status            int `json:"status"`
}

// end --- 获取异步发送企业会话消息的发送进度

// start --- 获取异步向企业会话发送消息的结果

type TopCorpMessageCorpconversationGetsendresultResponse struct {
	ErrorResponse                                            TopErrorResponse `json:"error_response"`
	DingtalkCorpMessageCorpconversationGetsendresultResponse TopCMCGResponse  `json:"dingtalk_corp_message_corpconversation_getsendresult_response"`
}

type TopCMCGResponse struct {
	Result TopCMCGRResult
}

type TopCMCGResult struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool
	SendResult      interface{} `json:"send_result"`
}

// end --- 获取异步向企业会话发送消息的结果

// start --- 获取多个用户的签到记录

type SmartworkCheckinRecordGetRequest struct {
	UserIDList []string `json:"userid_list"`
	StartTime  int64    `json:"start_time"`
	EndTime    int64    `json:"end_time"`
	Cursor     int      `json:"cursor"`
	Size       int      `json:"size"`
}

// end --- 获取多个用户的签到记录

// start --- 复制审批流

type SmartworkBpmsProcessCopyRequest struct {
	AgentID       int64  `json:"agent_id"`
	ProcessCode   string `json:"process_code"`
	BizCategoryID string `json:"biz_category_id"`
	ProcessName   string `json:"process_name"`
	Description   string `json:"description"`
}

// end --- 复制审批流

// start --- 更新审批流

type SmartworkBpmsProcessSyncRequest struct {
	AgentID           int64  `json:"agent_id"`
	SrcProcessCode    string `json:"src_process_code"`
	TargetProcessCode string `json:"target_process_code"`
	BizCategoryID     string `json:"biz_category_id"`
	ProcessName       string `json:"process_name"`
}

// end ---

// start --- 发起审批实例

type SmartworkBpmsProcessinstanceCreateRequest struct {
	AgentID              int64               `json:"agent_id"`
	ProcessCode          string              `json:"process_code"`
	OriginatorUserID     string              `json:"originator_user_id"`
	DeptID               int                 `json:"dept_id"`
	Approvers            []string            `json:"approvers"`
	CCList               []string            `json:"cc_list"`
	CCPosition           []string            `json:"cc_position"`
	FormComponentValueVo []map[string]string `json:"form_component_values"`
}

// end ---

// start --- 获取审批实例列表

type SmartworkBpmsProcessinstanceListRequest struct {
	ProcessCode string   `json:"process_code"`
	StartTime   int64    `json:"start_time"`
	EndTime     int64    `json:"end_time"`
	Size        int      `json:"size"`
	Cursor      int      `json:"cursor"`
	UserIDList  []string `json:"userid_list"`
}

// end --- 获取审批实例列表

// start --- 添加外部联系人|更新外部联系人

type CorpExtcontactRequest struct {
	Title          string   `json:"title,omitempty"`
	LabelIDs       []int    `json:"label_ids"`
	ShareDeptIDs   []int    `json:"share_dept_ids,omitempty"`
	Address        string   `json:"address,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	FollowerUserID string   `json:"follower_user_id"`
	Name           string   `json:"name"`
	UserID         string   `json:"user_id,omitempty"`
	StateCode      string   `json:"state_code"`
	CompanyName    string   `json:"company_name,omitempty"`
	ShareUserIDs   []string `json:"share_user_ids,omitempty"`
	Mobile         string   `json:"mobile"`
}

// end --- 添加外部联系人|更新外部联系人
