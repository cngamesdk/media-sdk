package model

import "errors"

// ========== 更新客户人群 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/update

// CustomAudiencesUpdateReq 更新客户人群请求
type CustomAudiencesUpdateReq struct {
	GlobalReq
	AccountID   int64  `json:"account_id"`            // 推广帐号 id (必填)
	AudienceID  int64  `json:"audience_id"`           // 人群 id (必填)
	Name        string `json:"name,omitempty"`        // 人群名称，1-32字节；name/description/cooperated 至少填一个
	Description string `json:"description,omitempty"` // 人群描述，1-100字节；name/description/cooperated 至少填一个
	Cooperated  *bool  `json:"cooperated,omitempty"`  // 深度数据合作；name/description/cooperated 至少填一个
}

func (p *CustomAudiencesUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新客户人群请求参数
func (p *CustomAudiencesUpdateReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AudienceID == 0 {
		return errors.New("audience_id为必填")
	}
	if p.Name == "" && p.Description == "" && p.Cooperated == nil {
		return errors.New("name、description、cooperated至少填写一个")
	}
	if p.Name != "" && len(p.Name) > 32 {
		return errors.New("name长度不能超过32字节")
	}
	if p.Description != "" && len(p.Description) > 100 {
		return errors.New("description长度不能超过100字节")
	}
	return nil
}

// CustomAudiencesUpdateResp 更新客户人群响应（应答字段为空）
type CustomAudiencesUpdateResp struct{}
