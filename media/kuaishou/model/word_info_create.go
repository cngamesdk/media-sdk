package model

import "errors"

// WordInfoCreateReq 创建关键词请求
type WordInfoCreateReq struct {
	accessTokenReq
	AdvertiserId int64          `json:"advertiser_id"` // 广告主账号ID，必填
	CampaignId   int64          `json:"campaign_id"`   // 广告计划ID，必填
	UnitId       int64          `json:"unit_id"`       // 广告单元ID，必填；需保证是搜索广告Unit
	WordInfos    []WordInfoItem `json:"word_infos"`    // 关键词信息，必填
}

// WordInfoItem 关键词信息
type WordInfoItem struct {
	Word      string `json:"word"`       // 关键词内容，必填；新增关键词：最大长度20，不支持制表符、换行符、Emoji表情等特殊字符
	MatchType int    `json:"match_type"` // 关键词匹配类型，必填；1-精确匹配，2-短语匹配，3-广泛匹配
}

func (receiver *WordInfoCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CampaignId <= 0 {
		err = errors.New("campaign_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	if len(receiver.WordInfos) == 0 {
		err = errors.New("word_infos is empty")
		return
	}
	for _, item := range receiver.WordInfos {
		if len(item.Word) == 0 {
			err = errors.New("word is empty in word_infos")
			return
		}
		if item.MatchType < 1 || item.MatchType > 3 {
			err = errors.New("match_type must be 1, 2 or 3")
			return
		}
	}
	return
}

// WordInfoCreateError 添加失败关键词
type WordInfoCreateError struct {
	Word        string `json:"word"`         // 关键词内容
	MatchType   int    `json:"match_type"`   // 匹配类型：1-精确匹配，2-短语匹配，3-广泛匹配
	ErrorReason string `json:"error_reason"` // 失败原因
}

// WordInfoCreateSuccess 添加成功关键词
type WordInfoCreateSuccess struct {
	WordInfoId int64  `json:"word_info_id"` // 关键词ID
	Word       string `json:"word"`         // 关键词内容
	MatchType  int    `json:"match_type"`   // 匹配类型：1-精确匹配，2-短语匹配，3-广泛匹配
	PutStatus  int    `json:"put_status"`   // 投放状态：1-投放中，2-已暂停，3-已删除
}

// WordInfoCreateResp 创建关键词响应数据（仅data部分）
type WordInfoCreateResp struct {
	ErrorList   []WordInfoCreateError   `json:"error_list"`   // 添加失败关键词列表
	SuccessList []WordInfoCreateSuccess `json:"success_list"` // 添加成功关键词列表
}
