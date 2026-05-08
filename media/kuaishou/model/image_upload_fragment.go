package model

import "errors"

// ImageUploadFragmentReq 图文件分片上传请求
// https://{endpoint}/api/upload/fragment?fragment_id=分片序号&upload_token=xxx
type ImageUploadFragmentReq struct {
	Endpoint    string `json:"-"` // 上传域名，来自获取上传token接口的endpoint字段，必填
	UploadToken string `json:"-"` // 上传token，来自获取上传token接口，必填
	FragmentId  int    `json:"-"` // 分片序号id，从0开始，按序递增，必填
	Fragment    []byte `json:"-"` // 文件分片，二进制流切分，通过Body(data-binary)上传，必填
}

func (receiver *ImageUploadFragmentReq) Format() {}

func (receiver *ImageUploadFragmentReq) Validate() (err error) {
	if receiver.Endpoint == "" {
		err = errors.New("endpoint is empty")
		return
	}
	if receiver.UploadToken == "" {
		err = errors.New("upload_token is empty")
		return
	}
	if receiver.FragmentId < 0 {
		err = errors.New("fragment_id must be >= 0")
		return
	}
	if len(receiver.Fragment) == 0 {
		err = errors.New("fragment is empty")
		return
	}
	return
}

// ImageUploadFragmentResp 图文件分片上传响应
type ImageUploadFragmentResp struct {
	Result   int    `json:"result"`   // 返回码/错误码，1表示成功，其他都为失败
	Checksum string `json:"checksum"` // 分片文件校验和
	Size     int64  `json:"size"`     // 文件大小，字节Byte
}
