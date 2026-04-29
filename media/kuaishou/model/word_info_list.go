package model

import "errors"

// WordInfoListReq 获取关键词列表请求
type WordInfoListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
	UnitId       int64 `json:"unit_id"`       // 广告单元ID，必填
}

func (receiver *WordInfoListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	return
}

// WordInfoDetail 关键词详情
type WordInfoDetail struct {
	WordInfoId   int64  `json:"word_info_id"`  // 关键词ID
	Word         string `json:"word"`          // 关键词内容
	MatchType    int    `json:"match_type"`    // 匹配类型：1=精确匹配 2=短语匹配 3=广泛匹配
	ReviewStatus int    `json:"review_status"` // 审核状态：1=审核中 2=审核通过 3=审核未通过 7=待送审
	PutStatus    int    `json:"put_status"`    // 投放状态：1=投放中 2=已暂停 3=已删除
	Status       int    `json:"status"`        // 综合状态：101=已删除 102=审核失败 103=审核中 104=已暂停 105=投放中 106=待送审
}

// WordInfoListResp 获取关键词列表响应数据（仅data部分）
type WordInfoListResp struct {
	TotalCount int64            `json:"total_count"` // 关键词数量
	Details    []WordInfoDetail `json:"details"`     // 关键词信息列表
}
