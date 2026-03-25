package model

import "errors"

// 常量定义 - 广告主类型
const (
	AdvertiserTypeDirect = "DIRECT_ADVERTISER" // 直接广告主
	AdvertiserTypeSub    = "SUB_ADVERTISER"    // 子广告主
)

type OrganizationAccountRelationGetReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id,omitempty"`       // 广告账户id
	AdvertiserType string `json:"advertiser_type,omitempty"`  // 广告主类型
	BusinessUnitID int64  `json:"business_unit_id,omitempty"` // 业务单元id
	CursorPageReq
}

func (p *OrganizationAccountRelationGetReq) Format() {
	p.GlobalReq.Format()
	p.CursorPageReq.Format()
}

// Validate 验证账户查询参数
func (p *OrganizationAccountRelationGetReq) Validate() error {

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if validateErr := p.CursorPageReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证广告主类型
	if p.AdvertiserType != "" && p.AdvertiserType != AdvertiserTypeDirect && p.AdvertiserType != AdvertiserTypeSub {
		return errors.New("advertiser_type值无效，允许值：DIRECT_ADVERTISER、SUB_ADVERTISER")
	}

	return nil
}

type OrganizationAccountRelationGetResp struct {
	List []*AccountInfo `json:"list"` // 返回信息列表
	PageInfoContainer
	CursorPageInfoContainer
}

// AccountInfo 账户信息
type AccountInfo struct {
	AccountID       int64             `json:"account_id"`             // 广告账户id
	CorporationName string            `json:"corporation_name"`       // 企业名称
	IsBid           bool              `json:"is_bid"`                 // 是否竞价广告广告账户即腾讯广告账户
	IsMp            bool              `json:"is_mp"`                  // 是否微信MP广告账户
	IsAdx           bool              `json:"is_adx,omitempty"`       // 废弃字段禁止使用
	CommentList     []*AccountComment `json:"comment_list,omitempty"` // 广告主备注
}

// AccountComment 广告主备注
type AccountComment struct {
	UserID  int64  `json:"user_id"` // 用户id
	Comment string `json:"comment"` // 广告主备注
}
