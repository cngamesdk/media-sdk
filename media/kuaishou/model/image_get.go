package model

import "errors"

// ImageGetReq 查询图片信息请求
type ImageGetReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"`         // 广告主ID，必填
	ImageToken   string `json:"image_token,omitempty"` // 图片token，可选
}

func (receiver *ImageGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImageGetReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// ImageGetResp 查询图片信息响应数据（仅data部分）
type ImageGetResp struct {
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
