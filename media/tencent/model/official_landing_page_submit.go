package model

import "errors"

// ========== 官方落地页-送审落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_submit/update

// OfficialLandingPageSubmitUpdateReq 官方落地页送审请求（POST JSON）
type OfficialLandingPageSubmitUpdateReq struct {
	GlobalReq
	AccountId int64 `json:"account_id"`        // 广告主帐号 id (必填)
	PageId    int64 `json:"page_id,omitempty"` // 落地页服务 id
}

func (r *OfficialLandingPageSubmitUpdateReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证送审落地页请求参数
func (r *OfficialLandingPageSubmitUpdateReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	return r.GlobalReq.Validate()
}

// OfficialLandingPageSubmitUpdateResp 送审落地页响应
type OfficialLandingPageSubmitUpdateResp struct {
	AccountId     int64 `json:"account_id"`      // 广告主帐号 id
	PageId        int64 `json:"page_id"`         // 落地页服务 id
	LandingPageId int   `json:"landing_page_id"` // 官方落地页 id
}
