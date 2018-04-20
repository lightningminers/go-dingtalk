# 用户相关

## UserIdByCode

通过Code换取userid，返回结果为 `UserIdResponse`

```go
type UserIdResponse struct {
	OpenAPIResponse
	UserId   string `json:"userid"`
	DeviceId string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}
```

**示例**

```go
c.UserIdByCode(code)
```

## UserIdByUnionId

通过UnionId获取UserId，返回结果为 `UserIdByUnionIdResponse`

```go
type UserIdByUnionIdResponse struct {
	OpenAPIResponse
	UserId      string `json:"userid"`
	ContactType int    `json:"contactType"`
}
```

**示例**

```go
c.UserIdByUnionId(unionId)
```

## UserInfoByUserId

通过userid换取用户详细信息，返回结果为 `UserInfoResponse`

```go
type UserInfoResponse struct {
	OpenAPIResponse
	UserId          string `json:"userid"`
	OpenId          string `json:"openid"`
	Name            string
	Tel             string
	WorkPlace       string
	Remark          string
	Mobile          string
	Email           string
	OrgEmail        string
	Active          bool
	IsAdmin         bool
	IsBoos          bool
	DingId          string
	UnionId         string
	IsHide          bool
	Department      []int
	Position        string
	Avatar          string
	Jobnumber       string
	IsSenior        bool
	StateCode       string
	OrderInDepts    string
	IsLeaderInDepts string
	Extattr         interface{}
	Roles           []Roles
}

type Roles struct {
	Id        int
	Name      string
	GroupName string
}
```

**示例**

```go
c.UserInfoByUserId(userId, lang)
```