package model

import "errors"

// DmpFileUploadReq 文件上传请求(新)
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v2/dmp/population/file/upload
// 注意：同一广告主ID的文件不支持并行上传，单文件最大1G，单日上传总容量50G
type DmpFileUploadReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	Type         int    `json:"-"` // 匹配类型，必填：1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 7=OAID 8=OAID_MD5 9=手机号_SHA256 16=CAID
	File         []byte `json:"-"` // 上传的文件，必填，支持.txt(UTF-8)或zip压缩包
	FileName     string `json:"-"` // 文件名，必填
}

func (receiver *DmpFileUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpFileUploadReq) Validate() (err error) {
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

// DmpFileUploadResp 文件上传响应数据(新)（仅data部分）
type DmpFileUploadResp struct {
	AdvertiserId   int64  `json:"advertiser_id"`    // 广告主ID
	MatchType      string `json:"match_type"`       // 匹配类型
	Md5            string `json:"md5"`              // md5
	FilePath       string `json:"file_path"`        // 文件路径，包含作为文件唯一标识的字符串
	FileSize       int64  `json:"file_size"`        // 文件大小
	UploadFileType string `json:"upload_file_type"` // 文件类型：TXT或ZIP
	RecordSize     int64  `json:"record_size"`      // 文件行数
	Type           int    `json:"type"`             // 匹配类型：1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 7=OAID 8=OAID_MD5 9=手机号_SHA256 16=CAID
}
