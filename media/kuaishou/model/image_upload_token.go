package model

import "errors"

// ImageUploadTokenReq 获取图片上传token请求
// https://ad.e.kuaishou.com/rest/openapi/gw/ad/common/upload/token/generate
type ImageUploadTokenReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	FileType     string `json:"file_type"`     // 文件类型，视频文件传"mp4"，安卓应用文件传"apk"，必填
}

func (receiver *ImageUploadTokenReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImageUploadTokenReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.FileType == "" {
		err = errors.New("file_type is empty")
		return
	}
	return
}

// ImageUploadTokenResp 获取图片上传token响应数据（仅data部分）
type ImageUploadTokenResp struct {
	UploadToken string   `json:"upload_token"` // 文件上传token
	Endpoint    []string `json:"endpoint"`     // 文件上传的域名
}
