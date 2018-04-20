# 起步

## 安装

```bash
$ go get -u github.com/icepy/go-dingtalk
```

## 企业

使用`NewDingTalkCompanyClient`，并且传入一个`DTCompanyConfig`类型的指针，`DTCompanyConfig`用于用户自己组装配置，如：`CorpID`，`CorpSecret`等，`NewDingTalkCompanyClient`会返回一个`DingTalkClient`类型指针给用户使用。

**示例**

```go
package main

import (
	"os"
	"github.com/icepy/go-dingtalk/src"
)

func main() {
	c := getCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
}

func getCompanyDingTalkClient() *dingtalk.DingTalkClient {
	CorpID := os.Getenv("CorpId")
	CorpSecret := os.Getenv("CorpSecret")
	config := &dingtalk.DTCompanyConfig{
		CorpID:     CorpID,
		CorpSecret: CorpSecret,
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}

```

## ISV

待...

## 小程序

待...