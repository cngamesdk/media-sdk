package model

import "errors"

// ========== 获取计量处罚明细 ==========
// https://developers.e.qq.com/v3.0/docs/api/punish_detail/get

// PunishDetailGetReq 获取计量处罚明细请求
type PunishDetailGetReq struct {
	GlobalReq
	PartitionTime           int64    `json:"partition_time,omitempty"`             // 统计周期，格式 YYYYMM，不传则默认所有周期，最小值 100000，最大值 999999
	StartPartitionTime      int64    `json:"start_partition_time,omitempty"`       // 开始月份，yyyyMM 格式，最小值 100000，最大值 999999
	EndPartitionTime        int64    `json:"end_partition_time,omitempty"`         // 结束月份，yyyyMM 格式，最小值 100000，最大值 999999
	AccountID               int64    `json:"account_id"`                           // 当前账户 id (必填)
	AgencyUidList           []string `json:"agency_uid_list,omitempty"`            // 代理商 id 列表，字段长度最小 1 字节，长度最大 255 字节
	OpsAdvertiserNameList   []string `json:"ops_advertiser_name_list,omitempty"`   // 广告主主体名称列表，字段长度最小 1 字节，长度最大 255 字节
	ZcAgentGroup            string   `json:"zc_agent_group,omitempty"`             // 代理商政策集团，字段长度最小 1 字节，长度最大 255 字节
	FirstLevelIndustryName  string   `json:"first_level_industry_name,omitempty"`  // 开户一级行业，字段长度最小 1 字节，长度最大 255 字节
	SecondLevelIndustryName string   `json:"second_level_industry_name,omitempty"` // 开户二级行业，字段长度最小 1 字节，长度最大 255 字节
	AdgroupIDList           []int64  `json:"adgroup_id_list,omitempty"`            // 广告 id 列表
	DynamicCreativeIDList   []int64  `json:"dynamic_creative_id_list,omitempty"`   // 动态创意 id 列表
	ComponentIDList         []int64  `json:"component_id_list,omitempty"`          // 组件 id 列表
	ElementType             string   `json:"element_type,omitempty"`               // 元素类型，字段长度最小 1 字节，长度最大 255 字节
	IsElementAppeal         *int     `json:"is_element_appeal,omitempty"`          // 是否有申诉记录，1 是，0 否
	AgencyNameList          []string `json:"agency_name_list,omitempty"`           // 代理商名称列表，字段长度最小 1 字节，长度最大 255 字节
	PhysicalFingerList      []string `json:"physical_finger_list,omitempty"`       // 元素指纹列表，字段长度最小 1 字节，长度最大 255 字节
	RejectReason            string   `json:"reject_reason,omitempty"`              // 拒绝原因，字段长度最小 1 字节，长度最大 255 字节
	ReviewBeginTime         string   `json:"review_begin_time,omitempty"`          // 审核开始时间，yyyy-MM-dd，字段长度为 10 字节
	ReviewEndTime           string   `json:"review_end_time,omitempty"`            // 最后审核时间，yyyy-MM-dd，字段长度为 10 字节
	QueryAccountIDList      []int64  `json:"query_account_id_list,omitempty"`      // 需要查询的账户 id 列表
	PageNum                 int      `json:"page_num,omitempty"`                   // 页码，最小值 1，最大值 999999，默认值 1
	PageSize                int      `json:"page_size,omitempty"`                  // 每页大小，最小值 1，最大值 1000，默认值 20
}

func (p *PunishDetailGetReq) Format() {
	p.GlobalReq.Format()
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
}

// Validate 验证获取计量处罚明细请求参数
func (p *PunishDetailGetReq) Validate() error {
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

// PunishDetailGetResp 获取计量处罚明细响应
type PunishDetailGetResp struct {
	DataList []*PunishDetailItem    `json:"data_list,omitempty"` // 计量处罚明细数据列表
	PageConf *PunishMetricsPageConf `json:"page_conf,omitempty"` // 分页内容
}

// PunishDetailItem 计量处罚明细数据
type PunishDetailItem struct {
	PartitionTime           int64  `json:"partition_time,omitempty"`             // 统计周期，格式 YYYYMM
	AccountID               int64  `json:"account_id,omitempty"`                 // 账户 id
	AgencyUid               string `json:"agency_uid,omitempty"`                 // 代理商 id
	LastReviewTime          string `json:"last_review_time,omitempty"`           // 最后审核时间，yyyy-MM-dd
	OpsAdvertiserName       string `json:"ops_advertiser_name,omitempty"`        // 广告主主体名称
	ZcAgentGroup            string `json:"zc_agent_group,omitempty"`             // 代理商政策集团
	FirstLevelIndustryName  string `json:"first_level_industry_name,omitempty"`  // 开户一级行业
	SecondLevelIndustryName string `json:"second_level_industry_name,omitempty"` // 开户二级行业
	AdgroupID               int64  `json:"adgroup_id,omitempty"`                 // 广告 id，AID
	DynamicCreativeID       int64  `json:"dynamic_creative_id,omitempty"`        // 动态创意 id，DCID
	ComponentID             int64  `json:"component_id,omitempty"`               // 组件 id
	ElementType             string `json:"element_type,omitempty"`               // 元素类型
	ElementValue            string `json:"element_value,omitempty"`              // 元素值
	PhysicalFinger          string `json:"physical_finger,omitempty"`            // 元素指纹
	RejectReason            string `json:"reject_reason,omitempty"`              // 拒绝原因
	IsElementAppeal         int    `json:"is_element_appeal,omitempty"`          // 是否有申诉记录，1 是，0 否
	AgencyName              string `json:"agency_name,omitempty"`                // 代理商名称
	AccountFrame            string `json:"account_frame,omitempty"`              // 账户 K 框
}
