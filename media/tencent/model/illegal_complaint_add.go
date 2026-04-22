package model

import "errors"

// ========== 新增广告主违规申述 ==========
// https://developers.e.qq.com/v3.0/docs/api/illegal_complaint/add

// IllegalComplaintAddReq 新增广告主违规申述请求（multipart/form-data）
type IllegalComplaintAddReq struct {
	GlobalReq
	AccountID       int64  // 广告主帐号 id，有操作权限的帐号 id，不支持代理商 id (必填)
	IllegalOrderID  string // 违规单 id (必填)
	ComplaintReason string // 申述理由 (必填)
	File            []byte // 上传申述证据文件，只支持上传 zip 文件，不超过 100MB (必填)
	FileName        string // 文件名（含扩展名），用于 multipart 表单 (必填)
}

func (p *IllegalComplaintAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证新增广告主违规申述请求参数
func (p *IllegalComplaintAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.IllegalOrderID == "" {
		return errors.New("illegal_order_id为必填")
	}
	if p.ComplaintReason == "" {
		return errors.New("complaint_reason为必填")
	}
	if len(p.File) == 0 {
		return errors.New("file为必填")
	}
	if p.FileName == "" {
		return errors.New("file_name为必填")
	}
	return nil
}

// IllegalComplaintAddResp 新增广告主违规申述响应（应答字段为无）
type IllegalComplaintAddResp struct{}
