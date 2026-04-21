package model

import "errors"

// ========== 删除客户人群 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/delete

// CustomAudiencesDeleteReq 删除客户人群请求
type CustomAudiencesDeleteReq struct {
	GlobalReq
	AccountID  int64 `json:"account_id"`  // 推广帐号 id (必填)
	AudienceID int64 `json:"audience_id"` // 人群 id (必填)
}

func (p *CustomAudiencesDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证删除客户人群请求参数
func (p *CustomAudiencesDeleteReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AudienceID == 0 {
		return errors.New("audience_id为必填")
	}
	return nil
}

// CustomAudiencesDeleteResp 删除客户人群响应（应答字段为空）
type CustomAudiencesDeleteResp struct{}
