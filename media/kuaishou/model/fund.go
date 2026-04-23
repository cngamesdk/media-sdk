package model

import "errors"

// AdvertiserFundGetReq 获取广告账户余额信息请求
type AdvertiserFundGetReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`
}

func (receiver *AdvertiserFundGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserFundGetReq) Validate() (err error) {
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

// AdvertiserFundGetResp 获取广告账户余额信息响应数据（仅data部分）
type AdvertiserFundGetResp struct {
	Balance                       float64 `json:"balance"`                           // 账户总余额，单位：元
	RechargeBalance               float64 `json:"recharge_balance"`                  // 充值余额，单位：元
	ContractRebate                float64 `json:"contract_rebate"`                   // 框返余额，单位：元
	DirectRebate                  float64 `json:"direct_rebate"`                     // 激励余额，单位：元
	ExtendedBalance               float64 `json:"extended_balance"`                  // 平台激励余额，单位：元
	TotalSharedWalletBalance      float64 `json:"total_shared_wallet_balance"`       // 共享子钱包总余额，单位：元
	TotalFrameSharedWalletBalance float64 `json:"total_frame_shared_wallet_balance"` // 年框共享钱包总余额，单位：元
}
