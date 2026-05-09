package model

import "errors"

// DmpPopulationCreateReq 人群包创建请求(新)
// https://ad.e.kuaishou.com/rest/openapi/gw/dmp/v2/dmp/population/upload
type DmpPopulationCreateReq struct {
	accessTokenReq
	AdvertiserId    int64    `json:"advertiser_id"`       // 广告主ID，必填
	OrientationName string   `json:"orientation_name"`    // 人群包名称，必填，不能大于50个字符，不得重复
	Type            int      `json:"type"`                // 匹配类型，必填：1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 7=OAID 8=OAID_MD5 9=手机号_SHA256 16=CAID
	FilePaths       []string `json:"file_paths"`          // 文件路径，必填，来自文件上传接口，每次最多10个、总大小<3G
	GidMatch        *bool    `json:"gid_match,omitempty"` // gid匹配，默认false，需要开白申请
}

func (receiver *DmpPopulationCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.OrientationName == "" {
		err = errors.New("orientation_name is empty")
		return
	}
	if len([]rune(receiver.OrientationName)) > 50 {
		err = errors.New("orientation_name cannot exceed 50 characters")
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

// DmpPopulationCreateResp 人群包创建响应数据(新)（仅data部分）
type DmpPopulationCreateResp struct {
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
