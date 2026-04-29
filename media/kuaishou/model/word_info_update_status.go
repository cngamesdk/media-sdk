package model

import "errors"

// WordInfoUpdateStatusReq 修改关键词投放状态请求
type WordInfoUpdateStatusReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主账号ID，必填
	WordInfoIds  []int64 `json:"word_info_ids"` // 关键词ID，必填；ID不重复，最大数量20，关键词未删除
	PutStatus    int     `json:"put_status"`    // 投放状态，必填；1-投放，2-暂停，3-删除
}

func (receiver *WordInfoUpdateStatusReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoUpdateStatusReq) Validate() (err error) {
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
	if receiver.PutStatus < 1 || receiver.PutStatus > 3 {
		err = errors.New("put_status must be 1, 2 or 3")
		return
	}
	return
}

// WordInfoUpdateStatusResp 修改关键词投放状态响应数据（仅data部分）
type WordInfoUpdateStatusResp struct {
	WordInfoIds []int64 `json:"word_info_ids"` // 修改成功关键词ID列表
}
