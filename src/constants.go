package dingtalk

// 包外部可用的常量
const (
	VERSION               = "0.1"
	OAPIURL               = "https://oapi.dingtalk.com/"
	TOPAPIURL             = "https://eco.taobao.com/router/rest"
	MessageTypeText       = "text"
	MessageTypeActionCard = "action_card"
	MessageTypeImage      = "image"
	MessageTypeVoice      = "voice"
	MessageTypeFile       = "file"
	MessageTypeLink       = "link"
	MessageTypeOA         = "oa"
	MessageTypeMarkdown   = "markdown"
)

// 包内部用的常量
const (
	signMD5         = "MD5"
	signHMAC        = "HMAC"
	topFormat       = "json"
	topV            = "2.0"
	topSimplify     = false
	topSecret       = "github.com/icepy"
	topSignMethod   = signMD5
	typeJSON        = "application/json"
	typeForm        = "application/x-www-form-urlencoded"
	typeMultipart   = "multipart/form-data"
	aesEncodeKeyLen = 43
)

// Top接口
const (
	corpRoleSimpleList                         = "dingtalk.corp.role.simplelist"
	corpRoleList                               = "dingtalk.corp.role.list"
	corpHealthStepinfoGetuserstatus            = "dingtalk.corp.health.stepinfo.getuserstatus"
	corpHealthStepinfoListByUserid             = "dingtalk.corp.health.stepinfo.listbyuserid"
	corpRoleAddRolesForemps                    = "dingtalk.corp.role.addrolesforemps"
	corpRoleRemoveRolesForemps                 = "dingtalk.corp.role.removerolesforemps"
	corpRoleDeleteRole                         = "dingtalk.corp.role.deleterole"
	corpRoleGetRoleGroup                       = "dingtalk.corp.role.getrolegroup"
	corpMessageCorpconversationAsyncsend       = "dingtalk.corp.message.corpconversation.asyncsend"
	corpMessageCorpconversationAsyncsendbycode = "dingtalk.corp.message.corpconversation.asyncsendbycode"
	corpMessageCorpconversationGetsendprogress = "dingtalk.corp.message.corpconversation.getsendprogress"
	corpMessageCorpconversationGetsendresult   = "dingtalk.corp.message.corpconversation.getsendresult"
	smartworkAttendsListschedule               = "dingtalk.smartwork.attends.listschedule"
	smartworkAttendsGetsimplegroups            = "dingtalk.smartwork.attends.getsimplegroups"
	smartworkCheckinRecordGet                  = "dingtalk.smartwork.checkin.record.get"
	smartworkBpmsProcessCopy                   = "dingtalk.smartwork.bpms.process.copy"
	smartworkBpmsProcessSync                   = "dingtalk.smartwork.bpms.process.sync"
	smartworkBpmsProcessinstanceCreate         = "dingtalk.smartwork.bpms.processinstance.create"
	smartworkBpmsProcessinstanceList           = "dingtalk.smartwork.bpms.processinstance.list"
	corpExtcontactCreate                       = "dingtalk.corp.extcontact.create"
	corpExtcontactUpdate                       = "dingtalk.corp.extcontact.update"
	corpExtcontactList                         = "dingtalk.corp.extcontact.list"
	corpExtcontactGet                          = "dingtalk.corp.extcontact.get"
	corpExtcontactListlabelgroups              = "dingtalk.corp.extcontact.listlabelgroups"
)
