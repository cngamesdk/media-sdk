package model

import "errors"

// WordBidWeightFilterParam 优词提量列表筛选参数
type WordBidWeightFilterParam struct {
	UnitId int64  `json:"unit_id"` // 广告单元ID
	Word   string `json:"word"`    // 关键词（模糊匹配）
}

// WordBidWeightPageInfo 分页信息
type WordBidWeightPageInfo struct {
	CurrentPage int   `json:"current_page"` // 当前页id，必填
	PageSize    int   `json:"page_size"`    // 页大小，必填
	TotalCount  int64 `json:"total_count"`  // 总条数
}

// WordBidWeightListReq 获取优词提量列表请求
type WordBidWeightListReq struct {
	accessTokenReq
	AdvertiserId int64                     `json:"advertiser_id"` // 广告主id，必填
	AllAvailable bool                      `json:"all_available"` // 是否获取全部
	FilterParam  *WordBidWeightFilterParam `json:"filter_param"`  // 筛选
	PageInfo     WordBidWeightPageInfo     `json:"page_info"`     // 分页，必填
}

func (receiver *WordBidWeightListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordBidWeightListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PageInfo.CurrentPage <= 0 {
		err = errors.New("page_info.current_page is empty")
		return
	}
	if receiver.PageInfo.PageSize <= 0 {
		err = errors.New("page_info.page_size is empty")
		return
	}
	return
}

// ScopeInfo 生效范围信息
type ScopeInfo struct {
	Scope    int      `json:"scope"`     // 生效范围
	UnitInfo []string `json:"unit_info"` // 广告组信息
}

// UnitBidWeightInfo 广告组优词提量信息
type UnitBidWeightInfo struct {
	BidWeight    float64   `json:"bid_weight"`     // 提量系数
	CampaignId   int64     `json:"campaign_id"`    // 计划id
	Id           int64     `json:"id"`             // id
	OldBidWeight float64   `json:"old_bid_weight"` // 旧提量系数
	Scope        int       `json:"scope"`          // 生效范围
	ScopeInfo    ScopeInfo `json:"scope_info"`     // 生效范围信息
	UnitId       int64     `json:"unit_id"`        // 广告组id
	UnitName     string    `json:"unit_name"`      // 广告组名称
	Word         string    `json:"word"`           // 关键词
}

// WordBidWeightInfo 优词提量信息
type WordBidWeightInfo struct {
	BidWeight float64             `json:"bid_weight"` // 提量系数，1.1-2区间，最多一位小数
	Scope     int                 `json:"scope"`      // 生效范围，1-账户维度，2-广告组维度
	UnitInfo  []UnitBidWeightInfo `json:"unit_info"`  // 广告组信息，scope=2时返回
	Word      string              `json:"word"`       // 关键词内容
}

// WordBidWeightListResp 获取优词提量列表响应数据（仅data部分）
type WordBidWeightListResp struct {
	CurrentPage int                 `json:"current_page"` // 当前页
	List        []WordBidWeightInfo `json:"list"`         // 优词信息
	PageSize    int                 `json:"page_size"`    // 页大小
	TotalCount  int64               `json:"total_count"`  // 优词数量
}
