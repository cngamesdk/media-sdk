package model

import "errors"

// ImageUploadCompleteReq 图合并分片请求
// https://{endpoint}/api/upload/complete?fragment_count=分片总数&upload_token=xxxx
type ImageUploadCompleteReq struct {
	Endpoint      string `json:"-"` // 上传域名，来自获取上传token接口的endpoint字段，必填
	UploadToken   string `json:"-"` // 上传token，来自获取上传token接口，必填
	FragmentCount int    `json:"-"` // 分片总数，必须是分片上传接口所有完整+成功的分片数量，必填
}

func (receiver *ImageUploadCompleteReq) Format() {}

func (receiver *ImageUploadCompleteReq) Validate() (err error) {
	if receiver.Endpoint == "" {
		err = errors.New("endpoint is empty")
		return
	}
	if receiver.UploadToken == "" {
		err = errors.New("upload_token is empty")
		return
	}
	if receiver.FragmentCount <= 0 {
		err = errors.New("fragment_count must be > 0")
		return
	}
	return
}

// ImageUploadCompleteResp 图合并分片响应
type ImageUploadCompleteResp struct {
	Result int `json:"result"` // 返回码/错误码，1表示成功，其他都为失败
}
