package model

import "errors"

// NegativeWord 否定词内容
type NegativeWord struct {
	ExactWords  []string `json:"exact_words"`  // 精确否定词，最大数量200，单个词最大长度20
	PhraseWords []string `json:"phrase_words"` // 短语否定词，最大数量200，单个词最大长度20
}

// WordInfoNegativeWordUpdateReq 更新否定词请求
type WordInfoNegativeWordUpdateReq struct {
	accessTokenReq
	AdvertiserId int64        `json:"advertiser_id"` // 广告主id，必填
	NegativeWord NegativeWord `json:"negative_word"` // 否定词内容，必填
	UnitIds      []int64      `json:"unit_ids"`      // 广告单元ID，最大数量20
}

func (receiver *WordInfoNegativeWordUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoNegativeWordUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.UnitIds) == 0 {
		err = errors.New("unit_ids is empty")
		return
	}
	if len(receiver.UnitIds) > 20 {
		err = errors.New("unit_ids max length is 20")
		return
	}
	if len(receiver.NegativeWord.ExactWords) > 200 {
		err = errors.New("exact_words max length is 200")
		return
	}
	if len(receiver.NegativeWord.PhraseWords) > 200 {
		err = errors.New("phrase_words max length is 200")
		return
	}
	return
}

// ErrorMsg 错误信息
type ErrorMsg struct {
	ErrorMsg string `json:"error_msg"` // 错误信息
	Id       int64  `json:"id"`        // id
}

// WordInfoNegativeWordUpdateResp 更新否定词响应数据（仅data部分）
type WordInfoNegativeWordUpdateResp struct {
	Errors         []ErrorMsg `json:"errors"`           // 错误信息列表
	SuccessUnitIds []int64    `json:"success_unit_ids"` // 更新成功单元ID列表
	ErrorUnitIds   []int64    `json:"error_unit_ids"`   // 更新失败单元ID列表
}
