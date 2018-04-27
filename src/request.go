package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (dtc *DingTalkClient) httpRPC(path string, params url.Values, requestData interface{}, responseData Unmarshallable, isvGetCInfo ...interface{}) error {
	if dtc.DevType == "company" {
		if dtc.AccessToken != "" {
			if params == nil {
				params = url.Values{}
			}
			if params.Get("access_token") == "" {
				params.Set("access_token", dtc.AccessToken)
			}
		}
	}
	if dtc.DevType == "isv" {
		cur := isvGetCInfo[0]
		switch v := cur.(type) {
		case *DTIsvGetCompanyInfo:
			switch path {
			case "service/get_permanent_code", "service/activate_suite", "service/get_corp_token", "service/get_auth_info":
				if dtc.SuiteAccessToken != "" {
					if params == nil {
						params = url.Values{}
					}
					if params.Get("suite_access_token") == "" {
						params.Set("suite_access_token", dtc.SuiteAccessToken)
					}
				}
			default:
				if v.AuthAccessToken != "" {
					if params == nil {
						params = url.Values{}
					}
					if params.Get("access_token") == "" {
						params.Set("access_token", v.AuthAccessToken)
					}
				}
			}
		default:
			panic(errors.New("ERROR: *DTIsvGetCompanyInfo Error"))
		}
	}
	if dtc.DevType == "personalMini"{
		if dtc.SNSAccessToken != "" && path != "sns/getuserinfo" {
			if params == nil {
				params = url.Values{}
			}
			if params.Get("access_token") == "" {
				params.Set("access_token", dtc.SNSAccessToken)
			}
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

func (dtc *DingTalkClient) httpSNS(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	if dtc.SNSAccessToken != "" && path != "sns/getuserinfo" {
		if params == nil {
			params = url.Values{}
		}
		if params.Get("access_token") == "" {
			params.Set("access_token", dtc.SNSAccessToken)
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

func (dtc *DingTalkClient) httpSSO(path string, params url.Values, requestData interface{}, responseData Unmarshallable) error {
	if dtc.SSOAccessToken != "" {
		if params == nil {
			params = url.Values{}
		}
		if params.Get("access_token") == "" {
			params.Set("access_token", dtc.SSOAccessToken)
		}
	}
	return dtc.httpRequest("oapi", path, params, requestData, responseData)
}

func (dtc *DingTalkClient) httpTOP(requestData interface{}, responseData interface{}) error {
	var params []string
	var paramsJoin string
	var cipher []byte
	var cipherString string
	if body, ok := requestData.(TopMapRequest); ok {
		body["sign_method"] = dtc.TopConfig.TopSignMethod
		if dtc.DevType == "company" {
			body["session"] = dtc.AccessToken
		}
		body["format"] = dtc.TopConfig.TopFormat
		body["v"] = dtc.TopConfig.TopV
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)
		body["timestamp"] = tm.Format("2006-01-02 03:04:05 PM")
		if dtc.TopConfig.TopFormat == "json" {
			body["simplify"] = dtc.TopConfig.TopSimplify
		}
		params = sortParamsKey(body)
		paramsJoin = strings.Join(params, "")
		if dtc.TopConfig.TopSignMethod == signMD5 {
			paramsJoin = dtc.TopConfig.TopSecret + paramsJoin + dtc.TopConfig.TopSecret
			cipher = encryptMD5(paramsJoin)
		}
		if dtc.TopConfig.TopSignMethod == signHMAC {
			cipher = encryptHMAC(paramsJoin, dtc.TopConfig.TopSecret)
		}
		cipherString = byteToHex(cipher)
		body["sign"] = cipherString
		fmt.Printf("Top Params=%s\n", body)
		return dtc.httpRequest("tapi", nil, addPostBody(body), nil, responseData)
	}
	return errors.New("requestData Not TopMapRequest Type")
}

func addPostBody(topMap TopMapRequest) url.Values {
	body := url.Values{}
	for k, v := range topMap {
		switch v.(type) {
		case []string:
			for _, h := range v.([]string) {
				body.Add(k, h)
			}
		case []int:
			for _, h := range v.([]int) {
				body.Add(k, string(h))
			}
		default:
			body.Add(k, fmt.Sprintf("%s", v))
		}
	}
	return body
}

func encryptMD5(paramsJoin string) []byte {
	hMd5 := md5.New()
	hMd5.Write([]byte(paramsJoin))
	return hMd5.Sum(nil)
}

func encryptHMAC(paramsJoin string, secret string) []byte {
	hHmac := hmac.New(md5.New, []byte(secret))
	hHmac.Write([]byte(paramsJoin))
	return hHmac.Sum([]byte(""))
}

func byteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(strings.ToUpper(s))
	}
	return buffer.String()
}

func sortParamsKey(topParams TopMapRequest) []string {
	var t []string
	keys := topParams.keys()
	sort.Strings(keys)
	for _, k := range keys {
		t = append(t, k+fmt.Sprintf("%s", topParams[k]))
	}
	return t
}

func (dtc *DingTalkClient) httpRequest(tagType string, path interface{}, params url.Values, requestData interface{}, responseData interface{}) error {
	var request *http.Request
	var requestUrl string
	client := dtc.HTTPClient

	if tagType == "oapi" {
		requestUrl = OAPIURL + path.(string) + "?" + params.Encode()
		fmt.Printf("requestUrl=%s\n", requestUrl)
		if requestData != nil {
			switch v := requestData.(type) {
			case *uploadFile:
				var b bytes.Buffer
				if v.Reader == nil {
					return errors.New("upload file is empty")
				}
				w := multipart.NewWriter(&b)
				fw, err := w.CreateFormFile(v.FieldName, v.FileName)
				if err != nil {
					return err
				}
				if _, err = io.Copy(fw, v.Reader); err != nil {
					return err
				}
				if err = w.Close(); err != nil {
					return err
				}
				request, _ = http.NewRequest("POST", requestUrl, &b)
				request.Header.Set("Content-Type", w.FormDataContentType())
			default:
				d, _ := json.Marshal(requestData)
				request, _ = http.NewRequest("POST", requestUrl, bytes.NewReader(d))
				request.Header.Set("Content-Type", typeJSON+"; charset=UTF-8")
			}
		} else {
			request, _ = http.NewRequest("GET", requestUrl, nil)
		}
	}
	if tagType == "tapi" {
		requestUrl = TOPAPIURL
		request, _ = http.NewRequest("POST", requestUrl, strings.NewReader(params.Encode()))
		request.Header.Set("Content-Type", typeForm+"; charset=UTF-8")
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Server Error: " + resp.Status)
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	pos := len(typeJSON)

	if tagType == "oapi" {
		if len(contentType) >= pos && contentType[0:pos] == typeJSON {
			if content, err := ioutil.ReadAll(resp.Body); err == nil {
				json.Unmarshal(content, responseData)
				switch responseData.(type) {
				case Unmarshallable:
					resData := responseData.(Unmarshallable)
					return resData.checkError()
				}
			}
		} else {
			switch v := responseData.(type) {
			case *MediaDownloadFileResponse:
				io.Copy(v.Writer, resp.Body)
			}
		}
	}
	if tagType == "tapi" {
		if content, err := ioutil.ReadAll(resp.Body); err == nil {
			v := reflect.ValueOf(responseData)
			v = v.Elem()
			v.SetBytes(content)
		}
	}
	return err
}
