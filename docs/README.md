---
home: true
actionText: Get Started →
actionLink: /guide/
features:
- title: 简单
  details: 导入SDK，立刻开发基于钉钉的应用
- title: 易用
  details: 化繁为简，将复杂的细节问题隐藏，给予你的是简单的输入输出式api
- title: 稳定
  details: 编写了完备的测试用例，保障SDK的稳定性
footer: MIT Licensed | Copyright © 2018-present icepy
---

### 起步就像数 1, 2, 3 一样容易

``` bash
# 安装
$ go get -u github.com/icepy/go-dingtalk
```

``` go
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

::: warning 注意事项
请正确的配置你的GOPATH
:::
