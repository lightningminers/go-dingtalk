# DingTalk Golang SDK

DingTalk Golang SDK https://github.com/icepy

# Feature Overview

- 支持企业，SSO，SNS免登
- 支持对access_token自动续期过期管理
- 支持注册钉钉事件回调
- 支持对钉钉事件回调消息签名的加解密
- 支持全部 Open api
- 支持全部 Top api，并且自动处理生成加密签名

# Test

- Test get auth scopes
- Test get company acess_token
- Test get company ticket
- Test upload file
- Test download file

```bash
$ cd __test__
$ go test
```

# Install

```bash
$ go get -u github.com/icepy/go-dingtalk
```

# Guide

[Document]()

# Example

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

# Help



# Contribute

- For a small change, just send a PR.
- For bigger changes open an issue for discussion before sending a PR.
- PR should have:
  - Test case
  - Documentation
  - Example (If it makes sense)
- You can also contribute by:
  - Reporting issues
  - Suggesting new features or enhancements
  - Improve/fix documentation

# 打赏

<div align="left">
  <img width="100" heigth="100" src="docs/weixin.png">
</div>

# License

MIT License

Copyright (c) 2018 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
