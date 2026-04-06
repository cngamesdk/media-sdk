package model

import "errors"

// ========== 创建关键词 ==========
// https://developers.e.qq.com/v3.0/docs/api/bidword/add

// 常量定义 - 关键词匹配方式
const (
	BidwordMatchTypeExact  = "EXACT_MATCH"  // 精确匹配
	BidwordMatchTypeWide   = "WIDE_MATCH"   // 广泛匹配
	BidwordMatchTypeWord   = "WORD_MATCH"   // 词语匹配
	BidwordMatchTypePhrase = "PHRASE_MATCH" // 短语匹配
)

// 常量定义 - 关键词暂停状态
const (
	BidwordStatusNormal  = "KEYWORD_STATUS_NORMAL"  // 正常
	BidwordStatusSuspend = "KEYWORD_STATUS_SUSPEND" // 暂停
)

// 常量定义 - 是否使用组出价
const (
	BidwordUseGroupPrice    = "USE_GROUP_PRICE"     // 使用组出价
	BidwordNotUseGroupPrice = "NOT_USE_GROUP_PRICE" // 不使用组出价
)

// 字段限制常量
const (
	MaxBidwordListCount     = 1000  // list 最大长度
	MinBidwordListCount     = 1     // list 最小长度
	MaxBidwordBytes         = 60    // bidword 最大字节数
	MinBidwordBytes         = 1     // bidword 最小字节数
	MaxBidPrice             = 99999 // bid_price 最大值
	MinBidPrice             = 1     // bid_price 最小值
	MaxLandingPageListCount = 10    // landing_page_list 最大长度
)

// PcLandingPageItem 兜底落地页内容
type PcLandingPageItem struct {
	PageType string    `json:"page_type"` // 落地页类型 (必填)
	PageSpec *PageSpec `json:"page_spec"` // 落地页内容
}

// PcLandingPageInfo 关键词落地页信息
type PcLandingPageInfo struct {
	LandingPageList []*PcLandingPageItem `json:"landing_page_list,omitempty"` // 兜底落地页内容列表 (0-10)，仅在特定 page_type 下可用
	PageType        string               `json:"page_type"`                   // 落地页类型 (必填)
	PageSpec        *PageSpec            `json:"page_spec"`                   // 落地页内容
}

// BidwordListItem 关键词列表项（请求）
type BidwordListItem struct {
	AdgroupID         int64              `json:"adgroup_id"`                     // 广告 id (必填)
	Bidword           string             `json:"bidword"`                        // 关键词词面 (必填)，1-60字节
	BidPrice          int                `json:"bid_price,omitempty"`            // 关键词出价，单位分，1-99999
	UseGroupPrice     string             `json:"use_group_price,omitempty"`      // 是否使用组出价
	MatchType         string             `json:"match_type"`                     // 关键词匹配方式 (必填)
	ConfiguredStatus  string             `json:"configured_status,omitempty"`    // 暂停状态
	DynamicCreativeID int64              `json:"dynamic_creative_id,omitempty"`  // 广告创意 id
	PcLandingPageInfo *PcLandingPageInfo `json:"pc_landing_page_info,omitempty"` // 关键词落地页信息
}

// Validate 验证单个关键词列表项
func (b *BidwordListItem) Validate() error {
	if b.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if b.Bidword == "" {
		return errors.New("bidword为必填")
	}
	if len(b.Bidword) < MinBidwordBytes || len(b.Bidword) > MaxBidwordBytes {
		return errors.New("bidword字段长度最小1字节，长度最大60字节")
	}
	if b.BidPrice != 0 && (b.BidPrice < MinBidPrice || b.BidPrice > MaxBidPrice) {
		return errors.New("bid_price最小值1，最大值99999")
	}
	if b.UseGroupPrice != "" &&
		b.UseGroupPrice != BidwordUseGroupPrice &&
		b.UseGroupPrice != BidwordNotUseGroupPrice {
		return errors.New("use_group_price值无效，允许值：USE_GROUP_PRICE、NOT_USE_GROUP_PRICE")
	}
	if b.MatchType == "" {
		return errors.New("match_type为必填")
	}
	if b.MatchType != BidwordMatchTypeExact &&
		b.MatchType != BidwordMatchTypeWide &&
		b.MatchType != BidwordMatchTypeWord &&
		b.MatchType != BidwordMatchTypePhrase {
		return errors.New("match_type值无效，允许值：EXACT_MATCH、WIDE_MATCH、WORD_MATCH、PHRASE_MATCH")
	}
	if b.ConfiguredStatus != "" &&
		b.ConfiguredStatus != BidwordStatusNormal &&
		b.ConfiguredStatus != BidwordStatusSuspend {
		return errors.New("configured_status值无效，允许值：KEYWORD_STATUS_NORMAL、KEYWORD_STATUS_SUSPEND")
	}
	if b.PcLandingPageInfo != nil {
		if err := validatePcLandingPageInfo(b.PcLandingPageInfo); err != nil {
			return err
		}
	}
	return nil
}

// validatePcLandingPageInfo 验证落地页信息
func validatePcLandingPageInfo(info *PcLandingPageInfo) error {
	if info.PageType == "" {
		return errors.New("pc_landing_page_info.page_type为必填")
	}
	if len(info.LandingPageList) > MaxLandingPageListCount {
		return errors.New("pc_landing_page_info.landing_page_list数组长度不能超过10")
	}
	for _, item := range info.LandingPageList {
		if item.PageType == "" {
			return errors.New("pc_landing_page_info.landing_page_list中page_type为必填")
		}
	}
	return nil
}

// BidwordAddReq 创建关键词请求
// https://developers.e.qq.com/v3.0/docs/api/bidword/add
type BidwordAddReq struct {
	GlobalReq
	AccountID int64              `json:"account_id"` // 广告主帐号 id (必填)
	List      []*BidwordListItem `json:"list"`       // 关键词列表 (必填)，1-1000
}

func (p *BidwordAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建关键词请求参数
func (p *BidwordAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.List) < MinBidwordListCount {
		return errors.New("list为必填，至少包含1个关键词")
	}
	if len(p.List) > MaxBidwordListCount {
		return errors.New("list数组长度不能超过1000")
	}
	for i, item := range p.List {
		if item == nil {
			return errors.New("list中存在空的关键词项")
		}
		if err := item.Validate(); err != nil {
			return errors.New("list[" + itoa(i) + "]: " + err.Error())
		}
	}
	return p.GlobalReq.Validate()
}

// itoa 整数转字符串（避免引入strconv包）
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := false
	if i < 0 {
		neg = true
		i = -i
	}
	buf := [20]byte{}
	pos := len(buf)
	for i >= 10 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}
	pos--
	buf[pos] = byte('0' + i)
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

// BidwordResultItem 关键词结果项（成功/失败列表共用）
type BidwordResultItem struct {
	Index             int    `json:"index"`                         // 请求列表中的序号
	BidwordID         int64  `json:"bidword_id,omitempty"`          // 关键词 id
	Bidword           string `json:"bidword,omitempty"`             // 关键词词面
	BidPrice          int    `json:"bid_price,omitempty"`           // 关键词出价，单位分
	MatchType         string `json:"match_type,omitempty"`          // 关键词匹配方式
	ConfiguredStatus  string `json:"configured_status,omitempty"`   // 暂停状态
	ErrorMsg          string `json:"error_msg,omitempty"`           // 错误信息
	ApprovalStatus    int    `json:"approval_status,omitempty"`     // 审核状态
	DynamicCreativeID int64  `json:"dynamic_creative_id,omitempty"` // 广告创意 id
}

// BidwordAddResp 创建关键词响应
// https://developers.e.qq.com/v3.0/docs/api/bidword/add
type BidwordAddResp struct {
	SuccessList []*BidwordResultItem `json:"success_list"` // 返回成功的关键词列表
	ErrorList   []*BidwordResultItem `json:"error_list"`   // 返回失败的关键词列表
}
