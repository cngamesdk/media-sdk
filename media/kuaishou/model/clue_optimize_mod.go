package model

import "errors"

// ClueOptimizeSwitchModItem 线索优选开关修改项
type ClueOptimizeSwitchModItem struct {
	ClueOptimizeType int  `json:"clue_optimize_type"` // 目标类型：表单转化-53 线索提交数-786 添加企业微信-763
	Status           bool `json:"status"`             // 开关状态：true=开启 false=关闭
}

// ClueOptimizeSwitchModReq 修改线索优选开关状态请求
type ClueOptimizeSwitchModReq struct {
	accessTokenReq
	AdvertiserId            int64                       `json:"advertiser_id"`
	ClueOptimizeSwitchTypes []ClueOptimizeSwitchModItem `json:"clue_optimize_switch_types"`
}

func (receiver *ClueOptimizeSwitchModReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ClueOptimizeSwitchModReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.ClueOptimizeSwitchTypes) <= 0 {
		err = errors.New("clue_optimize_switch_types is empty")
		return
	}
	return
}

// ClueOptimizeSwitchModResp 修改线索优选开关状态响应数据（仅data部分）
type ClueOptimizeSwitchModResp []ClueOptimizeSwitchModItem
