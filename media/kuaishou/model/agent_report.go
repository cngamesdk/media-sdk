package model

import "errors"

// AgentReportReq 代理商数据请求（t-1 数据需要第二天 9 点以后获取）
type AgentReportReq struct {
	accessTokenReq
	AgentId   int64  `json:"agent_id"`            // 代理商ID（注：非账户快手ID），在获取accessToken时返回，必填
	StartDate string `json:"start_date"`          // 过滤筛选条件，格式 yyyy-MM-dd，必填
	EndDate   string `json:"end_date"`            // 过滤筛选条件，格式 yyyy-MM-dd，必填
	Page      int    `json:"page,omitempty"`      // 请求的页码，默认为1
	PageSize  int    `json:"page_size,omitempty"` // 每页行数，默认为20，最大支持1000
	TimeType  string `json:"time_type,omitempty"` // 时间粒度：ALL=汇总，DAY=天，MONTH=月
}

func (receiver *AgentReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.StartDate) == 0 {
		err = errors.New("start_date is empty")
		return
	}
	if len(receiver.EndDate) == 0 {
		err = errors.New("end_date is empty")
		return
	}
	return
}

// AgentReportDetail 代理商数据明细
type AgentReportDetail struct {
	DateTime                        string `json:"date_time"`                            // 数据日期，格式：YYYY-MM-DD
	AccountId                       int64  `json:"account_id"`                           // 广告主ID
	UserId                          int64  `json:"user_id"`                              // 快手id
	AccountName                     string `json:"account_name"`                         // 广告主名称
	TotalChargedInYuan              int64  `json:"total_charged_in_yuan"`                // 总消耗（元）
	TotalBalanceInYuan              int64  `json:"total_balance_in_yuan"`                // 总余额（元）
	RealChargedInYuan               int64  `json:"real_charged_in_yuan"`                 // 现金消耗（元）
	TotalRebateRealChargedInYuan    int64  `json:"total_rebate_real_charged_in_yuan"`    // 返点消耗（元）
	ContractRebateRealChargedInYuan int64  `json:"contract_rebate_real_charged_in_yuan"` // 框返消耗（元）
	DirectRebateRealChargedInYuan   int64  `json:"direct_rebate_real_charged_in_yuan"`   // 激励消耗（元）
	CreditRealChargedInYuan         int64  `json:"credit_real_charged_in_yuan"`          // 信用消耗（元）
	ChargeDayOnDayPercent           string `json:"charge_day_on_day_percent"`            // 消耗环比
	PlayedNum                       int64  `json:"played_num"`                           // 播放数
	AdPhotoImpression               int64  `json:"ad_photo_impression"`                  // 封面曝光数
	AdPhotoClick                    int64  `json:"ad_photo_click"`                       // 封面点击数
	AdItemImpression                int64  `json:"ad_item_impression"`                   // 素材曝光数
	AdItemClick                     int64  `json:"ad_item_click"`                        // 行为数
	PhotoClickRatio                 string `json:"photo_click_ratio"`                    // 封面点击率
	ItemClickRatio                  string `json:"item_click_ratio"`                     // 行为点击率
	ChargedCampaignCount            int64  `json:"charged_campaign_count"`               // 有消费计划数
	ProductName                     string `json:"product_name"`                         // 产品名
	CorporationName                 string `json:"corporation_name"`                     // 企业名称
	FirstCostDay                    int64  `json:"first_cost_day"`                       // 首次消耗日期
	Industry                        string `json:"industry"`                             // 一级行业
	SecondIndustry                  string `json:"second_industry"`                      // 二级行业
}

// AgentReportResp 代理商数据响应数据（仅data部分）
type AgentReportResp struct {
	TotalCount int                 `json:"total_count"` // 数据的总行数
	Details    []AgentReportDetail `json:"details"`     // 数据明细信息
}
