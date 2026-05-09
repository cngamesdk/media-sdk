package model

import "errors"

// DmpDatasourceAddReq 数据源新建请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v1/datasource/add
type DmpDatasourceAddReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`    // 广告主ID，必填
	DataSourceName string   `json:"data_source_name"` // 数据源名称，必填，不能大于20个字符，不得重复
	MatchType      string   `json:"match_type"`       // 匹配类型，必填：IMEI/IDFA/IMEI_MD5/IDFA_MD5/OAID/OAID_MD5/MOBILE_MD5/MOBILE_SHA256
	SchemaType     string   `json:"schema_type"`      // 模版类型，必填：YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT
	FileKeys       []string `json:"file_keys"`        // 文件路径数组，必填，来自数据源文件上传接口，每次最多10个、总大小<3G
}

func (receiver *DmpDatasourceAddReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpDatasourceAddReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.DataSourceName == "" {
		err = errors.New("data_source_name is empty")
		return
	}
	if len([]rune(receiver.DataSourceName)) > 20 {
		err = errors.New("data_source_name cannot exceed 20 characters")
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
	if len(receiver.FileKeys) == 0 {
		err = errors.New("file_keys is empty")
		return
	}
	if len(receiver.FileKeys) > 10 {
		err = errors.New("file_keys cannot exceed 10 items")
		return
	}
	return
}

// DmpDatasourceAddResp 数据源新建响应数据（仅data部分）
type DmpDatasourceAddResp struct {
	AccountId      int64  `json:"account_id"`       // 广告主ID
	MatchType      int    `json:"match_type"`       // 匹配类型
	DataSourceName string `json:"data_source_name"` // 数据源名称
	SchemaType     string `json:"schema_type"`      // 模版类型：YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT
	AllFileSize    int64  `json:"all_file_size"`    // 文件大小
	CreateTime     int64  `json:"create_time"`      // 创建时间戳
	DataSourceId   int64  `json:"data_source_id"`   // 数据源ID
	CalcuStatus    int    `json:"calcu_status"`     // 数据源状态：0=未计算 1=计算中 2=计算成功 3=计算失败
	ErrorMessage   string `json:"error_message"`    // 错误信息，计算失败时有值
}
