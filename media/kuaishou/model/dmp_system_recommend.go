package model

import "errors"

// DmpSystemRecommendReq 系统推荐定向/排除人群包请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/v1/tool/population/system/recommend
type DmpSystemRecommendReq struct {
	accessTokenReq
	AdvertiserId     int64 `json:"advertiser_id"`               // 广告主ID，必填
	Type             int   `json:"type"`                        // 系统人群包类型，必填：1=定向 2=排除
	PopulationSource *int  `json:"population_source,omitempty"` // 人群包覆盖数来源：0=主站人群覆盖数 1=联盟人群覆盖数
}

func (receiver *DmpSystemRecommendReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpSystemRecommendReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Type != 1 && receiver.Type != 2 {
		err = errors.New("type must be 1(定向) or 2(排除)")
		return
	}
	return
}

// DmpSystemRecommendPopulation 系统推荐人群包详情（AdDmpPopulationView2）
type DmpSystemRecommendPopulation struct {
	OrientationId       int64  `json:"orientation_id"`        // 人群包ID
	OrientationName     string `json:"orientation_name"`      // 人群包名称
	AccountId           int64  `json:"account_id"`            // 账户ID
	Type                int    `json:"type"`                  // 人群数据类型
	PopulationType      int    `json:"population_type"`       // 人群包类型
	PopulationTypeName  string `json:"population_type_name"`  // 人群包类型名称
	RecordSize          int64  `json:"record_size"`           // 上传数量
	MatchSize           int64  `json:"match_size"`            // 匹配数量
	CoverNum            int64  `json:"cover_num"`             // 覆盖人数
	Status              int    `json:"status"`                // 状态
	SrcId               int64  `json:"src_id"`                // 来源ID
	CreateTime          int64  `json:"create_time"`           // 创建时间，13位毫秒级时间戳
	VerifyTime          int64  `json:"verify_time"`           // 更新时间，13位毫秒级时间戳
	ProfileTags         string `json:"profile_tags"`          // 画像标签
	Unbind              string `json:"unbind"`                // 解绑信息
	UnbindType          string `json:"unbind_type"`           // 解绑类型
	SuccessUnbind       string `json:"success_unbind"`        // 成功解绑
	FailUnbind          string `json:"fail_unbind"`           // 失败解绑
	TpCode              int    `json:"tp_code"`               // 第三方平台code
	IsExcludePopulation bool   `json:"is_exclude_population"` // 是否为排除人群包
	UpdateTime          int64  `json:"update_time"`           // 更新时间，13位毫秒级时间戳
	CategoryType        int    `json:"category_type"`         // 类目类型
	CanExclude          int    `json:"can_exclude"`           // 是否可排除
	CanTarget           int    `json:"can_target"`            // 是否可定向
	SrcType             int    `json:"src_type"`              // 人群包来源：0=DMP平台 1=MAPI平台 2=CDP平台 3=CDP投放共建类型
}

// DmpSystemRecommendResp 系统推荐定向/排除人群包响应数据（仅data部分，AdDmpSystemRecoView1）
type DmpSystemRecommendResp struct {
	IndustryTarget  []DmpSystemRecommendPopulation `json:"industry_target"`  // 行业人群
	ConsumeTarget   []DmpSystemRecommendPopulation `json:"consume_target"`   // 店铺消费者人群
	ReactTarget     []DmpSystemRecommendPopulation `json:"react_target"`     // 店铺内互动人群
	InterestTarget  []DmpSystemRecommendPopulation `json:"interest_target"`  // 店铺外兴趣人群
	NegativeExclude []DmpSystemRecommendPopulation `json:"negative_exclude"` // LA负反馈人群
	CommonExclude   []DmpSystemRecommendPopulation `json:"common_exclude"`   // 通用排除人群
}
