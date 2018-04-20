package dingtalk

import "fmt"

type OpenAPIResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Unmarshallable interface {
	checkError() error
}

func (oar *OpenAPIResponse) checkError() error {
	var err error
	if oar.ErrCode != 0 {
		err = fmt.Errorf("errcode: %d\nerrmsg: %s", oar.ErrCode, oar.ErrMsg)
	}
	return err
}
