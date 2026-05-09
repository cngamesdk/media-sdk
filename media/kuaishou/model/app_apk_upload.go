package model

import "errors"

// AppApkUploadReq 上传APK文件请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/upload/apk
// 此接口为新增&编辑安卓应用的前置步骤，上传APK生成blob_store_key后再调用「创建Android应用」接口
type AppApkUploadReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	File         []byte `json:"-"` // APK文件二进制内容，必填
	FileName     string `json:"-"` // 文件名，必填
	Type         *int   `json:"-"` // 上传分包用途，可选：不填=上传母包；1=上传母包并生成测试分包下载链接
}

func (receiver *AppApkUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppApkUploadReq) Validate() (err error) {
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

// AppApkUploadResp 上传APK文件响应数据（仅data部分）
type AppApkUploadResp struct {
	BlobStoreKey string `json:"blob_store_key"` // APK文件在快手的存储Key
	Url          string `json:"url"`            // APK文件地址
	TestApkUrl   string `json:"test_apk_url"`   // 测试分包下载地址
}
