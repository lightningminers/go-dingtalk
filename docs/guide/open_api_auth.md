# 授权

授权是钉钉来验证合法性的一种方式，也是为了保护用户的信息数据安全，当你需要使用免登来拿用户的信息时，授权是非常重要的一个步骤。

`说明c是调用NewDingTalkCompanyClient返回的一个指针`

## RefreshCompanyAccessToken

用于获取企业`access_token`，使用者不需要关心续期过期管理。

**示例**

```go
c.RefreshCompanyAccessToken()
c.AccessToken
```

## RefreshCompanySSOAccessToken

用于获取企业`sso_access_token`，使用者不需要关心续期过期管理。

**示例**

```go
c.RefreshCompanySSOAccessToken()
c.SSOAccessToken
```

## RefreshSNSAccessToken

用于获取企业`sns_access_token`，使用者不需要关心续期过期管理。

**示例**

```go
c.RefreshSNSAccessToken()
c.SNSAccessToken
```

## GetJSAPITicket

用于获取企业`Ticket`，使用者不需要关心续期过期管理

**示例**

```go
ticket, err := c.GetJSAPITicket()
if err == nil{
  // ticket
}
```

## GetConfig

配置企业的config，使用者不需要关心签名，返回结果为 `字符串JSON`

**示例**

```go
c.GetConfig(nonceStr, timestamp, url)
```

## GetAuthScopes

用于获取企业的授权范围，返回结果为 `ScopesResponse`

```go
type ScopesResponse struct {
	OpenAPIResponse
	AuthUserField  []string
	ConditionField []string
	AuthOrgScopes
}

type AuthOrgScopes struct {
	AuthedDept []int
	AuthedUser []string
}
```

**示例**

```go
scopes, err := c.GetAuthScopes()
if err == nil{
  // scopes
}
```