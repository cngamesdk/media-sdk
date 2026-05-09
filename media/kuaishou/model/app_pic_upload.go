package model

import "errors"

// AppPicUploadReq 上传图片请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/upload/pic
// 适用于将广告主原始图片上传至快手侧，生成CDN链接用于应用中心创编
type AppPicUploadReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	File         []byte `json:"-"` // 图片文件二进制内容，必填
	FileName     string `json:"-"` // 文件名，必填
	Type         int    `json:"-"` // 图片类型，必填：1=应用图标(450x450,<1MB,PNG/JPG/JPEG) 2=应用图片(9:20,≥720x1280,<2MB,PNG/JPG/JPEG) 5=单机承诺函(<10MB,PNG/JPG/JPEG) 6=APP备案截图 7=备案主体营业执照图片
}

var validAppPicTypes = map[int]bool{
	1: true, 2: true, 5: true, 6: true, 7: true,
}

func (receiver *AppPicUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppPicUploadReq) Validate() (err error) {
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
	if !validAppPicTypes[receiver.Type] {
		err = errors.New("type must be 1(应用图标) 2(应用图片) 5(单机承诺函) 6(APP备案截图) 7(备案主体营业执照)")
		return
	}
	return
}

// AppPicUploadResp 上传图片响应数据（仅data部分）
type AppPicUploadResp struct {
	Url string `json:"url"` // 所上传图片的快手CDN链接
}
