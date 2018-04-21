package dingtalk

import (
	"encoding/json"
)

type TopMapRequest map[string]interface{}

func (t TopMapRequest) keys() []string {
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

// 查询用户是否开启了钉钉运动
func (dtc *DingTalkClient) TopCorpHealthStepinfoGetUserStatus(userId string) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpHealthStepinfoGetuserstatus,
		"userid": userId,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 批量查询多个用户的钉钉运动步数
func (dtc *DingTalkClient) TopCorpHealthStepinfoListByUserid(userIds []string, statDate string) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":    corpHealthStepinfoListByUserid,
		"userids":   userIds,
		"stat_date": statDate,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取角色的员工列表
func (dtc *DingTalkClient) TopCorpRoleSimpleList(roleId int, size int, offset int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":  corpRoleSimpleList,
		"role_id": roleId,
		"size":    size,
		"offset":  offset,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取企业角色列表
func (dtc *DingTalkClient) TopCorpRoleList(size int, offset int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpRoleList,
		"size":   size,
		"offset": offset,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 批量为员工增加角色信息
func (dtc *DingTalkClient) TopCorpRoleAddRolesForemps(rolelIdList []int, userIdList []string) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":       corpRoleAddRolesForemps,
		"rolelid_list": rolelIdList,
		"userid_list":  userIdList,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 批量删除员工角的色信息
func (dtc *DingTalkClient) TopCorpRoleRemoveRolesForemps(roleIdList []int, userIdList []string) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":      corpRoleRemoveRolesForemps,
		"roleid_list": roleIdList,
		"userid_list": userIdList,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 删除角色信息
func (dtc *DingTalkClient) TopCorpRoleDeleteRole(roleId int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":  corpRoleDeleteRole,
		"role_id": roleId,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取角色组信息
func (dtc *DingTalkClient) TopCorpRoleGetRoleGroup(groupId int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":   corpRoleGetRoleGroup,
		"group_id": groupId,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 企业会话消息异步发送
func (dtc *DingTalkClient) TopCorpMessageCorpconversationAsyncsend(info *TopCorpMessageCorpconversationAsyncsendRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":   corpMessageCorpconversationAsyncsend,
		"msgtype":  info.MsgType,
		"agent_id": info.AgentId,
	}
	if info.ToAllUser {
		general["to_all_user"] = info.ToAllUser
	}
	if content, err := json.Marshal(info.Msgcontent); err == nil {
		general["msgcontent"] = string(content)
	}
	if len(info.UserIdList) > 0 {
		general["userid_list"] = info.UserIdList
	}
	if len(info.DeptIdList) > 0 {
		general["dept_id_list"] = info.DeptIdList
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 通过用户授权码异步向企业会话发送消息
func (dtc *DingTalkClient) TopCorpMessageCorpconversationAsyncsendbycode(info *TopCorpMessageCorpconversationAsyncsendbycodeRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":   corpMessageCorpconversationAsyncsendbycode,
		"code":     info.Code,
		"msgtype":  info.MsgType,
		"agent_id": info.AgentId,
	}
	if info.ToAllUser {
		general["to_all_user"] = info.ToAllUser
	}
	if content, err := json.Marshal(info.Msgcontent); err == nil {
		general["msgcontent"] = string(content)
	}
	if len(info.UserIdList) > 0 {
		general["userid_list"] = info.UserIdList
	}
	if len(info.DeptIdList) > 0 {
		general["dept_id_list"] = info.DeptIdList
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取异步发送企业会话消息的发送进度
func (dtc *DingTalkClient) TopCorpMessageCorpconversationGetsendprogress(agentId int, taskId int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":   corpMessageCorpconversationGetsendprogress,
		"agent_id": agentId,
		"task_id":  taskId,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取异步向企业会话发送消息的结果
func (dtc *DingTalkClient) TopCorpMessageCorpconversationGetsendresult(agentId interface{}, taskId interface{}) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpMessageCorpconversationGetsendresult,
	}
	if agentId != nil {
		if v, ok := agentId.(int); ok {
			general["agent_id"] = v
		}
	}
	if taskId != nil {
		if v, ok := taskId.(int); ok {
			general["task_id"] = v
		}
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 考勤排班信息按天全量查询接口
func (dtc *DingTalkClient) TopSmartworkAttendsListschedule(workDate string, offset int, size int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":    smartworkAttendsListschedule,
		"work_date": workDate,
		"offset":    offset,
		"size":      size,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取考勤组列表详情
func (dtc *DingTalkClient) TopSmartworkAttendsGetsimplegroups(offset int, size int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": smartworkAttendsGetsimplegroups,
		"offset": offset,
		"size":   size,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取多个用户的签到记录
func (dtc *DingTalkClient) TopSmartworkCheckinRecordGet(info *SmartworkCheckinRecordGetRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":      smartworkCheckinRecordGet,
		"userid_list": info.UserIDList,
		"start_time":  info.StartTime,
		"end_time":    info.EndTime,
		"cursor":      info.Cursor,
		"size":        info.Size,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 复制审批流
func (dtc *DingTalkClient) TopSmartworkBpmsProcessCopy(info *SmartworkBpmsProcessCopyRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":       smartworkBpmsProcessCopy,
		"agent_id":     info.AgentID,
		"process_code": info.ProcessCode,
	}
	if info.BizCategoryID != "" {
		general["biz_category_id"] = info.BizCategoryID
	}
	if info.ProcessName != "" {
		general["process_name"] = info.ProcessName
	}
	if info.Description != "" {
		general["description"] = info.Description
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 更新审批流
func (dtc *DingTalkClient) TopSmartworkBpmsProcessSync(info *SmartworkBpmsProcessSyncRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":              smartworkBpmsProcessSync,
		"agent_id":            info.AgentID,
		"src_process_code":    info.SrcProcessCode,
		"target_process_code": info.TargetProcessCode,
	}
	if info.BizCategoryID != "" {
		general["biz_category_id"] = info.BizCategoryID
	}
	if info.ProcessName != "" {
		general["process_name"] = info.ProcessName
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 发起审批实例
func (dtc *DingTalkClient) TopSmartworkBpmsProcessinstanceCreate(info *SmartworkBpmsProcessinstanceCreateRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":             smartworkBpmsProcessinstanceCreate,
		"process_code":       info.ProcessCode,
		"originator_user_id": info.OriginatorUserID,
		"dept_id":            info.DeptID,
		"approvers":          info.Approvers,
	}
	if info.AgentID > 0 {
		general["agent_id"] = info.AgentID
	}
	if len(info.CCList) > 0 {
		general["cc_list"] = info.CCList
	}
	if len(info.CCPosition) > 0 {
		general["cc_position"] = info.CCPosition
	}
	b, e := json.Marshal(info.FormComponentValueVo)
	if e == nil {
		general["form_component_values"] = string(b)
	} else {
		panic(e)
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取审批实例列表
func (dtc *DingTalkClient) TopSmartworkBpmsProcessinstanceList(info *SmartworkBpmsProcessinstanceListRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":       smartworkBpmsProcessinstanceList,
		"process_code": info.ProcessCode,
		"start_time":   info.StartTime,
	}
	if info.EndTime > 0 {
		general["end_time"] = info.EndTime
	}
	if info.Size > 0 {
		general["size"] = info.Size
	}
	if len(info.UserIDList) > 0 {
		general["userid_list"] = info.UserIDList
	}
	if info.Cursor >= 0 {
		general["cursor"] = info.Cursor
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 添加企业外部联系人
func (dtc *DingTalkClient) TopCorpExtcontactCreate(info *CorpExtcontactRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpExtcontactCreate,
	}
	b, e := json.Marshal(info)
	if e == nil {
		general["contace"] = string(b)
	} else {
		panic(e)
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 更新外部联系人
func (dtc *DingTalkClient) TopCorpExtcontactUpdate(info *CorpExtcontactRequest) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpExtcontactUpdate,
	}
	b, e := json.Marshal(info)
	if e == nil {
		general["contace"] = string(b)
	} else {
		panic(e)
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 获取外部联系人列表
func (dtc *DingTalkClient) TopCorpExtcontactList(size int, offset int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpExtcontactList,
		"size":   size,
		"offset": offset,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 外部单个联系人详情
func (dtc *DingTalkClient) TopCorpExtcontactGet(userID string) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method":  corpExtcontactGet,
		"user_id": userID,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}

// 外部联系人标签列表
func (dtc *DingTalkClient) TopCorpExtcontactListlabelgroups(size int, offset int) ([]byte, error) {
	var data []byte
	general := TopMapRequest{
		"method": corpExtcontactListlabelgroups,
		"size":   size,
		"offset": offset,
	}
	err := dtc.httpTOP(general, &data)
	return data, err
}
