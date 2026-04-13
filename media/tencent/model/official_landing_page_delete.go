package model

import "errors"

// ========== 官方落地页-删除落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page/delete

// OfficialLandingPageDeleteReq 官方落地页删除请求（POST JSON）
type OfficialLandingPageDeleteReq struct {
	GlobalReq
	AccountId int64 `json:"account_id"` // 广告主帐号 id (必填)
	PageId    int64 `json:"page_id"`    // 落地页服务 id (必填)
}

func (r *OfficialLandingPageDeleteReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证删除落地页请求参数
func (r *OfficialLandingPageDeleteReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.PageId == 0 {
		return errors.New("page_id为必填")
	}
	return r.GlobalReq.Validate()
}

// OfficialLandingPageDeleteRespData 删除落地页响应数据
type OfficialLandingPageDeleteRespData struct {
	AccountId     int64 `json:"account_id"`      // 广告主帐号 id
	PageId        int64 `json:"page_id"`         // 落地页服务 id
	LandingPageId int   `json:"landing_page_id"` // 官方落地页 id
}

// OfficialLandingPageDeleteResp 删除落地页响应
type OfficialLandingPageDeleteResp struct {
	Code      int                                `json:"code"`
	Message   string                             `json:"message"`
	MessageCn string                             `json:"message_cn"`
	Data      *OfficialLandingPageDeleteRespData `json:"data"`
}
