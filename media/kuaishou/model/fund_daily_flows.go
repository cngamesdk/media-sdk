package model

import "errors"

// AdvertiserFundDailyFlowsReq 获取广告账户流水信息请求
type AdvertiserFundDailyFlowsReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"`
	StartDate    string `json:"start_date,omitempty"` // 开始日期，格式：yyyy-MM-dd，默认当天
	EndDate      string `json:"end_date,omitempty"`   // 结束日期，格式：yyyy-MM-dd，默认当天
	Page         int    `json:"page,omitempty"`       // 查询页码，默认1
	PageSize     int    `json:"page_size,omitempty"`  // 单页行数，默认20，不超过500
}

func (receiver *AdvertiserFundDailyFlowsReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserFundDailyFlowsReq) Validate() (err error) {
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

// AdvertiserFundDailyFlowItem 广告账户流水明细
type AdvertiserFundDailyFlowItem struct {
	Date                        string  `json:"date"`                           // 日期
	AdvertiserId                int64   `json:"advertiser_id"`                  // 广告主ID
	Balance                     float64 `json:"balance"`                        // 日终结余，单位：元
	DailyCharge                 float64 `json:"daily_charge"`                   // 实耗总花费，单位：元
	RealCharged                 float64 `json:"real_charged"`                   // 实耗充值花费，单位：元
	ContractRebateRealCharged   float64 `json:"contract_rebate_real_charged"`   // 实耗框返花费，单位：元
	DirectRebateRealCharged     float64 `json:"direct_rebate_real_charged"`     // 实耗激励花费，单位：元
	DailyTransferIn             float64 `json:"daily_transfer_in"`              // 转入，单位：元
	DailyTransferOut            float64 `json:"daily_transfer_out"`             // 转出，单位：元
	RealRecharged               float64 `json:"real_recharged"`                 // 充值转入，单位：元
	ContractRebateRealRecharged float64 `json:"contract_rebate_real_recharged"` // 框返转入，单位：元
	DirectRebateRealRecharged   float64 `json:"direct_rebate_real_recharged"`   // 激励转入，单位：元
	DailyShareCharge            float64 `json:"daily_share_charge"`             // 每日共享总花费，单位：元
	ShareRealCharge             float64 `json:"share_real_charge"`              // 共享实耗花费，单位：元
	ShareCreditCharge           float64 `json:"share_credit_charge"`            // 共享信用花费，单位：元
	OrderTotalCharged           float64 `json:"order_total_charged"`            // 订单总花费，单位：元
	OrderRealCharged            float64 `json:"order_real_charged"`             // 订单充值花费，单位：元
	OrderContractCharged        float64 `json:"order_contract_charged"`         // 订单框返花费，单位：元
	OrderDirectCharged          float64 `json:"order_direct_charged"`           // 订单激励花费，单位：元
	FrameShareWalletCharge      float64 `json:"frame_share_wallet_charge"`      // 年框共享钱包花费，单位：元
}

// AdvertiserFundDailyFlowsResp 获取广告账户流水信息响应数据（仅data部分）
type AdvertiserFundDailyFlowsResp struct {
	TotalCount int                           `json:"total_count"` // 总记录数
	Details    []AdvertiserFundDailyFlowItem `json:"details"`     // 流水明细列表
}
