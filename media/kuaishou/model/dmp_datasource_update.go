package model

import "errors"

// DmpDatasourceUpdateReq 数据源更新请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v1/datasource/update
type DmpDatasourceUpdateReq struct {
	accessTokenReq
	AdvertiserId  int64    `json:"advertiser_id"`  // 广告主ID，必填
	OperationType int      `json:"operation_type"` // 操作类型，必填：0=全量更新 1=增量更新
	DataSourceId  int64    `json:"data_source_id"` // 数据源ID，必填
	FileKeys      []string `json:"file_keys"`      // 文件路径数组，必填，来自数据源文件上传接口，每次最多10个、总大小<3G
	MatchType     string   `json:"match_type"`     // 匹配类型，必填：IMEI/IDFA/IMEI_MD5/IDFA_MD5/OAID/OAID_MD5/MOBILE_MD5/MOBILE_SHA256
	SchemaType    string   `json:"schema_type"`    // 模版类型，必填：YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT
}

func (receiver *DmpDatasourceUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpDatasourceUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.OperationType != 0 && receiver.OperationType != 1 {
		err = errors.New("operation_type must be 0(全量更新) or 1(增量更新)")
		return
	}
	if receiver.DataSourceId <= 0 {
		err = errors.New("data_source_id is empty")
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
	if !validMatchTypes[receiver.MatchType] {
		err = errors.New("match_type must be one of: IMEI/IDFA/IMEI_MD5/IDFA_MD5/OAID/OAID_MD5/MOBILE_MD5/MOBILE_SHA256")
		return
	}
	if !validSchemaTypes[receiver.SchemaType] {
		err = errors.New("schema_type must be one of: YX/JY/DS/JR/BD/JK/QC/XS/SJ/QT")
		return
	}
	return
}

// DmpDatasourceUpdateResp 数据源更新响应数据（仅data部分）
type DmpDatasourceUpdateResp struct {
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
