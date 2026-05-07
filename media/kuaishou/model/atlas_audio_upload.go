package model

import "errors"

// AtlasAudioUploadReq 图文音频文件上传请求
type AtlasAudioUploadReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	File         []byte `json:"-"` // 音频文件二进制内容，必填
	FileName     string `json:"-"` // 音频文件名，必填
}

func (receiver *AtlasAudioUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AtlasAudioUploadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
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

// AtlasAudioUploadResp 图文音频文件上传响应数据（仅data部分）
type AtlasAudioUploadResp struct {
	BlobStoreKey string `json:"blob_store_key"` // 音频存储key，用于上传图文视频时传入
	BucketName   string `json:"bucket_name"`    // bucket名称
}
