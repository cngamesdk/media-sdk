package model

import "errors"

// DmpPopulationUpdateReq 人群包更新请求(新)
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v2/dmp/population/update
// 注意：每天每个人群包增量和缩量更新合计不超过10次；仅status=1,4,5,6时允许更新
type DmpPopulationUpdateReq struct {
	accessTokenReq
	AdvertiserId  int64    `json:"advertiser_id"`  // 广告主ID，必填
	OrientationId int64    `json:"orientation_id"` // 人群包ID，必填
	OperationType int      `json:"operation_type"` // 操作类型，必填：1=增量更新 3=缩量更新 4=全量更新
	Type          int      `json:"type"`           // 匹配类型，必填：1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 7=OAID 8=OAID_MD5 9=手机号_SHA256 16=CAID
	FilePaths     []string `json:"file_paths"`     // 文件路径，必填，来自文件上传接口，每次最多10个、总大小<3G
}

func (receiver *DmpPopulationUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.OrientationId <= 0 {
		err = errors.New("orientation_id is empty")
		return
	}
	validOpTypes := map[int]bool{1: true, 3: true, 4: true}
	if !validOpTypes[receiver.OperationType] {
		err = errors.New("operation_type must be 1(增量) 3(缩量) or 4(全量)")
		return
	}
	if receiver.Type <= 0 {
		err = errors.New("type is empty")
		return
	}
	if len(receiver.FilePaths) == 0 {
		err = errors.New("file_paths is empty")
		return
	}
	if len(receiver.FilePaths) > 10 {
		err = errors.New("file_paths cannot exceed 10 items")
		return
	}
	return
}

// DmpPopulationUpdateResp 人群包更新响应数据(新)（仅data部分）
type DmpPopulationUpdateResp struct {
	OrientationId   int64    `json:"orientation_id"`    // 人群包ID
	OrientationName string   `json:"orientation_name"`  // 人群包名称
	Type            int      `json:"type"`              // 匹配类型：1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 7=OAID 8=OAID_MD5 9=手机号_SHA256 16=CAID
	PopulationType  int      `json:"population_type"`   // 人群包类型
	Status          int      `json:"status"`            // 状态：0=计算中 1=已生效 2=已删除 3=上线中 4=已上线 5=计算失败 6=上线失败 7=已失效
	CreateTime      int64    `json:"create_time"`       // 创建时间戳
	RecordSize      int64    `json:"record_size"`       // 总行数
	FailedFilePaths []string `json:"failed_file_paths"` // 错误文件路径列表
	AdvertiserId    int64    `json:"advertiser_id"`     // 广告主ID
}
