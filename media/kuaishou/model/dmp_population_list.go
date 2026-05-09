package model

import "errors"

// DmpPopulationListReq 人群列表查询请求
// https://ad.e.kuaishou.com/rest/openapi/v2/dmp/population/list
type DmpPopulationListReq struct {
	accessTokenReq
	AdvertiserId   int64   `json:"advertiser_id"`             // 广告主ID，必填
	Status         *int    `json:"status,omitempty"`          // 人群包状态：0=计算中 1=已生效 2=已删除 3=上线中 4=已上线 5=计算失败 6=上线失败 7=已失效
	Page           int     `json:"page,omitempty"`            // 页码，默认1
	PageSize       int     `json:"page_size,omitempty"`       // 每页数量，默认20，最多500
	OrientationIds []int64 `json:"orientation_ids,omitempty"` // 人群包ID列表，最多500个，数量必须小于page_size
}

func (receiver *DmpPopulationListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PageSize > 500 {
		err = errors.New("page_size cannot exceed 500")
		return
	}
	if len(receiver.OrientationIds) > 500 {
		err = errors.New("orientation_ids cannot exceed 500 items")
		return
	}
	if len(receiver.OrientationIds) > 0 && receiver.PageSize > 0 && len(receiver.OrientationIds) >= receiver.PageSize {
		err = errors.New("orientation_ids count must be less than page_size")
		return
	}
	return
}

// DmpPopulationListResp 人群列表查询响应数据（仅data部分）
type DmpPopulationListResp struct {
	TotalCount int                       `json:"total_count"` // 总数量
	Details    []DmpPopulationListDetail `json:"details"`     // 人群包列表
}

// DmpPopulationListDetail 人群包详情
type DmpPopulationListDetail struct {
	OrientationId     int64  `json:"orientation_id"`      // 人群包ID
	OrientationName   string `json:"orientation_name"`    // 人群包名称
	Type              int    `json:"type"`                // 人群数据类型：0=非上传人群包 1=IMEI 2=IDFA 3=IMEI_MD5 4=IDFA_MD5 5=手机号-MD5 9=手机号_SHA256 16=CAID
	PopulationType    int    `json:"population_type"`     // 人群包类型
	SrcType           int    `json:"src_type"`            // 人群包来源：0=DMP平台 1=MAPI平台 2=CDP平台 3=CDP投放共建类型
	RecordSize        int64  `json:"record_size"`         // 上传数量
	MatchSize         int64  `json:"match_size"`          // 匹配数量
	CoverNum          int64  `json:"cover_num"`           // 覆盖人数
	Status            int    `json:"status"`              // 人群包状态：0=计算中 1=已生效 2=已删除 3=上线中 4=已上线 5=计算失败 6=上线失败 7=已失效
	CreateTime        int64  `json:"create_time"`         // 创建时间，13位毫秒级时间戳
	ThirdPlatformCode int    `json:"third_platform_code"` // 付费人群包第三方平台code
	ThirdPlatformName string `json:"third_platform_name"` // 付费人群包第三方供应商名称
	VerifyTime        int64  `json:"verify_time"`         // 更新时间，13位毫秒级时间戳
}
