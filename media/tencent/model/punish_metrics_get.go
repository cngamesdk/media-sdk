package model

import "errors"

// ========== 获取处罚指标数据 ==========
// https://developers.e.qq.com/v3.0/docs/api/punish_metrics/get

// PunishMetricsGetReq 获取处罚指标数据请求
type PunishMetricsGetReq struct {
	GlobalReq
	PartitionTime           int64    `json:"partition_time,omitempty"`             // 统计周期，格式 YYYYMM，不传则默认所有周期，最小值 100000，最大值 999999
	StartPartitionTime      int64    `json:"start_partition_time,omitempty"`       // 开始月份，yyyyMM 格式，最小值 100000，最大值 999999
	EndPartitionTime        int64    `json:"end_partition_time,omitempty"`         // 结束月份，yyyyMM 格式，最小值 100000，最大值 999999
	AccountID               int64    `json:"account_id"`                           // 当前账户 id (必填)
	OpsAdvertiserNameList   []string `json:"ops_advertiser_name_list,omitempty"`   // 客户名称列表，字段长度最小 1 字节，长度最大 255 字节
	ZcAgentGroup            string   `json:"zc_agent_group,omitempty"`             // 代理商政策集团，字段长度最小 1 字节，长度最大 255 字节
	FirstLevelIndustryName  string   `json:"first_level_industry_name,omitempty"`  // 开户一级行业，字段长度最小 1 字节，长度最大 255 字节
	SecondLevelIndustryName string   `json:"second_level_industry_name,omitempty"` // 开户二级行业，字段长度最小 1 字节，长度最大 255 字节
	PageNum                 int      `json:"page_num,omitempty"`                   // 页码，最小值 1，最大值 999999，默认值 1
	PageSize                int      `json:"page_size,omitempty"`                  // 每页大小，最小值 1，最大值 1000，默认值 20
}

func (p *PunishMetricsGetReq) Format() {
	p.GlobalReq.Format()
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
}

// Validate 验证获取处罚指标数据请求参数
func (p *PunishMetricsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PageNum < 1 {
		return errors.New("page_num最小值为1")
	}
	if p.PageSize < 1 || p.PageSize > 1000 {
		return errors.New("page_size必须在1-1000之间")
	}
	return nil
}

// PunishMetricsGetResp 获取处罚指标数据响应
type PunishMetricsGetResp struct {
	DataList []*PunishMetricsItem   `json:"data_list,omitempty"` // 处罚指标数据列表
	PageConf *PunishMetricsPageConf `json:"page_conf,omitempty"` // 分页内容
}

// PunishMetricsItem 处罚指标数据
type PunishMetricsItem struct {
	PartitionTime           int64  `json:"partition_time,omitempty"`             // 统计周期，格式 YYYYMM
	OpsAdvertiserName       string `json:"ops_advertiser_name,omitempty"`        // 客户名称，广告主主体名称
	ZcAgentGroup            string `json:"zc_agent_group,omitempty"`             // 代理商政策集团
	FirstLevelIndustryName  string `json:"first_level_industry_name,omitempty"`  // 开户一级行业
	SecondLevelIndustryName string `json:"second_level_industry_name,omitempty"` // 开户二级行业
	IsPenalized             int    `json:"is_penalized,omitempty"`               // 是否处罚（1=是，0=否）
	PenaltyReviewTotal      int64  `json:"penalty_review_total,omitempty"`       // 复审总量
	PenaltyReviewRejectCnt  int64  `json:"penalty_review_reject_cnt,omitempty"`  // 处罚拒绝量
	PenaltyRejectRate       string `json:"penalty_reject_rate,omitempty"`        // 处罚口径拒绝率，%
}

// PunishMetricsPageConf 处罚指标分页信息
type PunishMetricsPageConf struct {
	Page      int `json:"page,omitempty"`      // 页码，默认值 1
	PageSize  int `json:"pageSize,omitempty"`  // 每页大小，默认值 20
	TotalPage int `json:"totalPage,omitempty"` // 总页数
	TotalNum  int `json:"totalNum,omitempty"`  // 总条数
}
