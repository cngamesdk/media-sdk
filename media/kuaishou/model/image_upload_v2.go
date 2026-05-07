package model

import "errors"

// ImageUploadV2Req 上传图片v2请求
type ImageUploadV2Req struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	Type         int    `json:"-"` // 上传图片类型，必填
	UploadType   int    `json:"-"` // 上传方式，必填
	Signature    string `json:"-"` // 图片MD5值，可选
	Name         string `json:"-"` // 图片名称，可选
	Url          string `json:"-"` // 图片URL，upload_type为URL上传时填写
	File         []byte `json:"-"` // 图片文件二进制内容，upload_type为文件上传时填写
	FileName     string `json:"-"` // 图片文件名，upload_type为文件上传时填写
}

func (receiver *ImageUploadV2Req) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImageUploadV2Req) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Type <= 0 {
		err = errors.New("type is empty")
		return
	}
	if receiver.UploadType <= 0 {
		err = errors.New("upload_type is empty")
		return
	}
	if len(receiver.File) == 0 && receiver.Url == "" {
		err = errors.New("file or url is required")
		return
	}
	return
}

// ImageUploadV2Resp 上传图片v2响应数据（仅data部分）
type ImageUploadV2Resp struct {
	Url        string `json:"url"`         // 图片预览地址
	Width      int64  `json:"width"`       // 图片宽度
	Height     int64  `json:"height"`      // 图片高度
	Size       int64  `json:"size"`        // 图片大小
	Format     string `json:"format"`      // 图片格式
	Signature  string `json:"signature"`   // 图片MD5值
	ImageToken string `json:"image_token"` // 图片token
	PicId      string `json:"pic_id"`      // 图片库图片ID
	PicType    string `json:"pic_type"`    // 图片类型
}
