package model

import "errors"

// ========== 获取违规处罚列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/punishment_query/get

// PunishmentQueryGetReq 获取违规处罚列表请求
type PunishmentQueryGetReq struct {
	GlobalReq
	OrderIDList      []string            `json:"order_id_list,omitempty"`      // 违规单 id 列表，字段长度最小 1 字节，长度最大 50 字节
	AccountIDList    []int64             `json:"account_id_list"`              // 账户 id 列表 (必填)
	CompanyName      string              `json:"company_name,omitempty"`       // 公司名称，字段长度最小 1 字节，长度最大 255 字节
	IndustryIDList   []int64             `json:"industry_id_list,omitempty"`   // 开户行业 id 列表
	IllegalStartTime int64               `json:"illegal_start_time,omitempty"` // 违规开始日期，毫秒时间戳
	IllegalEndTime   int64               `json:"illegal_end_time,omitempty"`   // 违规结束日期，毫秒时间戳
	IllegalNodeList  []string            `json:"illegal_node_list,omitempty"`  // 违规节点/对象列表，字段长度最小 1 字节，长度最大 512 字节
	ActionTypeList   []int64             `json:"action_type_list,omitempty"`   // 账户处罚动作列表，最小值 1，最大值 999999
	StartTime        int64               `json:"start_time,omitempty"`         // 违规开始日期，毫秒时间戳
	EndTime          int64               `json:"end_time,omitempty"`           // 处罚结束日期，毫秒时间戳
	LevelList        []int64             `json:"level_list,omitempty"`         // 处罚等级列表
	IllegalSceneList []string            `json:"illegal_scene_list,omitempty"` // 处罚场景列表，字段长度最小 1 字节，长度最大 512 字节
	PageConf         *PunishmentPageConf `json:"page_conf"`                    // 分页内容 (必填)
}

// PunishmentPageConf 处罚查询分页参数
type PunishmentPageConf struct {
	Page      int `json:"page,omitempty"`       // 第几页，最小值 1，最大值 999999，默认值 1
	PageSize  int `json:"page_size,omitempty"`  // 每页条数，最小值 1，最大值 1000，默认值 10
	TotalPage int `json:"total_page,omitempty"` // 总页数，最小值 1，最大值 99999999
	TotalNum  int `json:"total_num,omitempty"`  // 总条数
}

func (p *PunishmentQueryGetReq) Format() {
	p.GlobalReq.Format()
	if p.PageConf != nil {
		if p.PageConf.Page <= 0 {
			p.PageConf.Page = 1
		}
		if p.PageConf.PageSize <= 0 {
			p.PageConf.PageSize = 10
		}
	}
}

// Validate 验证获取违规处罚列表请求参数
func (p *PunishmentQueryGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if len(p.AccountIDList) == 0 {
		return errors.New("account_id_list为必填")
	}
	if p.PageConf == nil {
		return errors.New("page_conf为必填")
	}
	if p.PageConf.Page < 1 {
		return errors.New("page_conf.page最小值为1")
	}
	if p.PageConf.PageSize < 1 || p.PageConf.PageSize > 1000 {
		return errors.New("page_conf.page_size必须在1-1000之间")
	}
	return nil
}

// PunishmentQueryGetResp 获取违规处罚列表响应
type PunishmentQueryGetResp struct {
	List     []*PunishmentItem   `json:"list,omitempty"`      // 处罚场景列表
	PageConf *PunishmentPageConf `json:"page_conf,omitempty"` // 分页内容
}

// PunishmentItem 处罚信息
type PunishmentItem struct {
	OrderID             string             `json:"order_id,omitempty"`               // 违规单 id
	UID                 int64              `json:"uid,omitempty"`                    // 账户 id
	CompanyName         string             `json:"company_name,omitempty"`           // 公司名称
	AccountType         int                `json:"account_type,omitempty"`           // 账户类型，1=直客，3=子客
	AgencyID            int64              `json:"agency_id,omitempty"`              // 服务商 id
	FirstIndustryID     int64              `json:"first_industry_id,omitempty"`      // 开户一级行业 id
	FirstIndustry       string             `json:"first_industry,omitempty"`         // 开户一级行业
	SecondIndustryID    int64              `json:"second_industry_id,omitempty"`     // 开户二级行业 id
	SecondIndustry      string             `json:"second_industry,omitempty"`        // 开户二级行业
	KpiFirstIndustryID  int64              `json:"kpi_first_industry_id,omitempty"`  // KPI 一级行业 id
	KpiFirstIndustry    string             `json:"kpi_first_industry,omitempty"`     // KPI 一级行业
	KpiSecondIndustryID int64              `json:"kpi_second_industry_id,omitempty"` // KPI 二级行业 id
	KpiSecondIndustry   string             `json:"kpi_second_industry,omitempty"`    // KPI 二级行业
	IllegalAid          string             `json:"illegal_aid,omitempty"`            // 违规广告 AID
	IllegalTid          string             `json:"illegal_tid,omitempty"`            // 违规 TID
	IllegalDcID         string             `json:"illegal_dc_id,omitempty"`          // 违规 DCID
	IllegalComponentID  string             `json:"illegal_component_id,omitempty"`   // 违规组件 id
	IllegalTime         string             `json:"illegal_time,omitempty"`           // 违规日期，格式 YYYY-MM-DD
	NewIllegalObj       string             `json:"new_illegal_obj,omitempty"`        // 违规节点/对象
	ActionType          int                `json:"action_type,omitempty"`            // 账户处罚动作
	DepositData         *PunishmentDeposit `json:"deposit_data,omitempty"`           // 保证金处罚内容
	CreateTime          string             `json:"create_time,omitempty"`            // 处罚时间，如 2025-08-01 18:01:07
	Level               int                `json:"level,omitempty"`                  // 处罚等级
	Reason              string             `json:"reason,omitempty"`                 // 处罚原因
	CertificateList     []string           `json:"certificate_list,omitempty"`       // 处罚凭证列表
	IllegalScene        string             `json:"illegal_scene,omitempty"`          // 处罚场景
	MdmPunishData       []*MdmPunishData   `json:"mdm_punish_data,omitempty"`        // 主体处罚内容列表
	AppealStatus        bool               `json:"appeal_status,omitempty"`          // 是否可申诉
}

// PunishmentDeposit 保证金处罚内容
type PunishmentDeposit struct {
	PunishRatio int64 `json:"punish_ratio,omitempty"` // 处罚比例
	PunishMoney int64 `json:"punish_money,omitempty"` // 处罚金额
	DeductID    int64 `json:"deduct_id,omitempty"`    // 保证金扣罚单号
	DeductMoney int64 `json:"deduct_money,omitempty"` // 保证金扣罚金额
}

// MdmPunishData 主体处罚内容
type MdmPunishData struct {
	PunishLevel int `json:"punish_level,omitempty"` // 处罚等级 value
	PunishCnt   int `json:"punish_cnt,omitempty"`   // 处罚次数
}
