package dingtalk

import (
	"fmt"
	"io"
	"net/url"
)

type FileUploadStartTransactionResponse struct {
	OpenAPIResponse
	UploadID string `json:"upload_id"`
}

type FileUploadSingleResponse struct {
	OpenAPIResponse
	MediaID string `json:"media_id"`
}

type FileUploadEndTransactionRequest struct {
	AgentID      string
	FileSize     int64
	ChunkNumbers int64
	UploadDI     string
}

type FileUploadEndTransactionResponse struct {
	OpenAPIResponse
	MediaID string `json:"media_id"`
}

type FileUploadChunkRequest struct {
	AgentID       string
	UploadID      string
	ChunkSequence int64
	FileName      string
	Reader        io.Reader
}

type FileUploadChunkResponse struct {
	OpenAPIResponse
}

// 开启文件上传事务
func (dtc *DingTalkClient) FileUploadStartTransaction(agentID string, fileSize int, chunkNumbers int) (FileUploadStartTransactionResponse, error) {
	var data FileUploadStartTransactionResponse
	params := url.Values{}
	params.Add("agent_id", agentID)
	params.Add("file_size", fmt.Sprintf("%d", fileSize))
	params.Add("chunk_numbers", fmt.Sprintf("%d", chunkNumbers))
	err := dtc.httpRPC("/file/upload/transaction", params, nil, &data)
	return data, err
}

// 提交文件上传事务
func (dtc *DingTalkClient) FileUploadEndTransaction(info *FileUploadEndTransactionRequest) (FileUploadEndTransactionResponse, error) {
	var data FileUploadEndTransactionResponse
	params := url.Values{}
	params.Add("agent_id", info.AgentID)
	params.Add("file_size", fmt.Sprintf("%d", info.FileSize))
	params.Add("chunk_numbers", fmt.Sprintf("%d", info.ChunkNumbers))
	params.Add("upload_id", info.UploadDI)
	err := dtc.httpRPC("file/upload/transaction", params, nil, &data)
	return data, err
}

// 上传文件块
func (dtc *DingTalkClient) FileUploadChunk(info *FileUploadChunkRequest) (FileUploadChunkResponse, error) {
	var data FileUploadChunkResponse
	params := url.Values{}
	params.Add("agent_id", info.AgentID)
	params.Add("upload_id", info.UploadID)
	params.Add("chunk_sequence", fmt.Sprintf("%d", info.ChunkSequence))
	upload := &uploadFile{
		FieldName: "file",
		FileName:  info.FileName,
		Reader:    info.Reader,
	}
	err := dtc.httpRPC("file/upload/chunk", params, upload, &data)
	return data, err
}

// 上传单个文件
func (dtc *DingTalkClient) FileUploadSingle(agentID string, fileSize int64, fileName string, reader io.Reader) (FileUploadSingleResponse, error) {
	var data FileUploadSingleResponse
	upload := &uploadFile{
		FieldName: "file",
		FileName:  fileName,
		Reader:    reader,
	}
	params := url.Values{}
	params.Add("agent_id", agentID)
	params.Add("file_size", fmt.Sprintf("%d", fileSize))
	err := dtc.httpRPC("file/upload/single", params, upload, &data)
	return data, err
}
