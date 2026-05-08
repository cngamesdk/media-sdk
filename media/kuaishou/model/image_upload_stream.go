package model

import "errors"

// ImageUploadStreamReq 图文件流式上传请求
// https://{endpoint}/api/upload/multipart?upload_token=xxx
type ImageUploadStreamReq struct {
	Endpoint    string `json:"-"` // 上传域名，来自获取上传token接口的endpoint字段，必填
	UploadToken string `json:"-"` // 上传token，来自获取上传token接口，必填
	File        []byte `json:"-"` // 文件二进制内容，通过Body(form-data)上传，必填
	FileName    string `json:"-"` // 文件名，必填
}

func (receiver *ImageUploadStreamReq) Format() {}

func (receiver *ImageUploadStreamReq) Validate() (err error) {
	if receiver.Endpoint == "" {
		err = errors.New("endpoint is empty")
		return
	}
	if receiver.UploadToken == "" {
		err = errors.New("upload_token is empty")
		return
	}
	if len(receiver.File) == 0 {
		err = errors.New("file is empty")
		return
	}
	if receiver.FileName == "" {
		err = errors.New("file_name is empty")
		return
	}
	return
}

// ImageUploadStreamResp 图文件流式上传响应
type ImageUploadStreamResp struct {
	Result int `json:"result"` // 返回码/错误码，1表示成功，其他都为失败
}
