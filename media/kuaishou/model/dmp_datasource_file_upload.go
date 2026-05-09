package model

import "errors"

// DmpDatasourceFileUploadReq 数据源文件上传请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v1/datasource/file/upload
// 注意：同一广告主ID的文件不支持并行上传，单文件最大1G，单日上传总容量50G
type DmpDatasourceFileUploadReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"-"` // 广告主ID，必填
	MatchType    string `json:"-"` // 匹配类型，必填：IMEI/IDFA/IMEI_MD5/IDFA_MD5/OAID/OAID_MD5/MOBILE_MD5/MOBILE_SHA256
	SchemaType   string `json:"-"` // 模版类型，必填：YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT
	File         []byte `json:"-"` // 上传的文件，必填，支持.csv格式
	FileName     string `json:"-"` // 文件名，必填
}

var validMatchTypes = map[string]bool{
	"IMEI": true, "IDFA": true, "IMEI_MD5": true, "IDFA_MD5": true,
	"OAID": true, "OAID_MD5": true, "MOBILE_MD5": true, "MOBILE_SHA256": true,
}

var validSchemaTypes = map[string]bool{
	"YX": true, "JY": true, "DS": true, "JR": true, "BD": true,
	"JK": true, "QC": true, "XS": true, "SJ": true, "QT": true,
}

func (receiver *DmpDatasourceFileUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpDatasourceFileUploadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if !validMatchTypes[receiver.MatchType] {
		err = errors.New("match_type must be one of: IMEI/IDFA/IMEI_MD5/IDFA_MD5/OAID/OAID_MD5/MOBILE_MD5/MOBILE_SHA256")
		return
	}
	if !validSchemaTypes[receiver.SchemaType] {
		err = errors.New("schema_type must be one of: YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT")
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

// DmpDatasourceFileUploadResp 数据源文件上传响应数据（仅data部分）
type DmpDatasourceFileUploadResp struct {
	AccountId      int64  `json:"account_id"`       // 广告主ID
	MatchType      string `json:"match_type"`       // 匹配类型
	Md5            string `json:"md5"`              // 文件md5
	FileKey        string `json:"file_key"`         // 文件路径
	FileSize       int64  `json:"file_size"`        // 文件大小
	UploadFileType string `json:"upload_file_type"` // 文件类型
	RecordSize     int64  `json:"record_size"`      // 文件行数
	MatchSize      int64  `json:"match_size"`       // 匹配行数
	SchemaType     int    `json:"schema_type"`      // 模版类型
}
