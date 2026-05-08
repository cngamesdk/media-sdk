package model

import "errors"

// ImageUploadTokenVerifyReq 领用上传token请求
// https://ad.e.kuaishou.com/rest/openapi/gw/ad/common/upload/token/verify
type ImageUploadTokenVerifyReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	UploadToken  string `json:"upload_token"`  // 上传token，必填
}

func (receiver *ImageUploadTokenVerifyReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImageUploadTokenVerifyReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UploadToken == "" {
		err = errors.New("upload_token is empty")
		return
	}
	return
}

// ImageUploadTokenVerifyResp 领用上传token响应数据（仅data部分）
type ImageUploadTokenVerifyResp struct {
	BlobStoreKey string `json:"blob_store_key"` // 文件存储位置的标识
}
