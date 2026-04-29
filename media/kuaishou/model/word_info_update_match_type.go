package model

import "errors"

// WordInfoUpdateMatchTypeReq 修改关键词匹配方式请求
type WordInfoUpdateMatchTypeReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主账号ID，必填
	WordInfoIds  []int64 `json:"word_info_ids"` // 关键词ID，必填；ID不重复，最大数量20，关键词未删除
	MatchType    int     `json:"match_type"`    // 匹配类型，必填；1-精确匹配，2-短语匹配，3-广泛匹配
}

func (receiver *WordInfoUpdateMatchTypeReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoUpdateMatchTypeReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.WordInfoIds) == 0 {
		err = errors.New("word_info_ids is empty")
		return
	}
	if len(receiver.WordInfoIds) > 20 {
		err = errors.New("word_info_ids max length is 20")
		return
	}
	if receiver.MatchType < 1 || receiver.MatchType > 3 {
		err = errors.New("match_type must be 1, 2 or 3")
		return
	}
	return
}

// WordInfoUpdateMatchTypeResp 修改关键词匹配方式响应数据（仅data部分）
type WordInfoUpdateMatchTypeResp struct {
	WordInfoIds []int64 `json:"word_info_ids"` // 修改成功关键词ID列表
}
