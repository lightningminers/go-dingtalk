package dingtalk

import (
	"io"
	"net/url"
)

type MediaUploadResponse struct {
	OpenAPIResponse
	Type      string
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

type MediaDownloadFileResponse struct {
	OpenAPIResponse
	MediaID string
	Writer  io.Writer
}

type uploadFile struct {
	FileName  string
	FieldName string
	Reader    io.Reader
}

// 上传媒体文件
func (dtc *DingTalkClient) MediaUpload(mediaType string, fileName string, reader io.Reader) (MediaUploadResponse, error) {
	var data MediaUploadResponse
	upload := &uploadFile{
		FieldName: "media",
		FileName:  fileName,
		Reader:    reader,
	}
	params := url.Values{}
	params.Add("type", mediaType)
	err := dtc.httpRPC("media/upload", params, upload, &data)
	return data, err
}

// 获取媒体文件
func (dtc *DingTalkClient) MediaDownloadFile(mediaID string, write io.Writer) error {
	var data MediaDownloadFileResponse
	data.MediaID = mediaID
	data.Writer = write
	params := url.Values{}
	params.Add("media_id", mediaID)
	err := dtc.httpRPC("media/downloadFile", params, nil, &data)
	return err
}
