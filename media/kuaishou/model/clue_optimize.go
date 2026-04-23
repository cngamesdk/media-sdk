package model

import "errors"

// ClueOptimizeSwitchReq 获取线索优选开关状态请求
type ClueOptimizeSwitchReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`
}

func (receiver *ClueOptimizeSwitchReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ClueOptimizeSwitchReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// ClueOptimizeSwitchItem 线索优选开关状态项
type ClueOptimizeSwitchItem struct {
	ClueOptimizeSwitch int    `json:"clue_optimize_switch"` // 线索优选开关类型：53=表单 786=私信留资 763=添加企业微信
	Name               string `json:"name"`                 // 线索优选开关名字
	Status             bool   `json:"status"`               // 开关状态：true=开启 false=关闭
}

// ClueOptimizeSwitchResp 获取线索优选开关状态响应数据（仅data部分）
type ClueOptimizeSwitchResp []ClueOptimizeSwitchItem
