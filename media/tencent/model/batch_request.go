package model

import "errors"

// ========== 创建批量请求 ==========
// https://developers.e.qq.com/v3.0/docs/api/batch_requests/add

// 字段限制常量
const (
	MinBatchRequestSpecCount = 1 // batch_request_spec 最小长度
)

// BatchRequestSpec 批量调用接口所需条件
type BatchRequestSpec struct {
	RelativePath string `json:"relative_path"` // 相对接口请求路径 (必填)
	// POST 请求不带参数，形如：v3.0/adgroups/update
	// GET  请求带参数，形如：v3.0/adgroups/get?account_id=12345
	Body string `json:"body,omitempty"` // POST 原始请求时的正文，JSON 编码的字符串
}

// Validate 验证单个批量请求条件
func (b *BatchRequestSpec) Validate() error {
	if b.RelativePath == "" {
		return errors.New("relative_path为必填")
	}
	return nil
}

// BatchRequestAddReq 创建批量请求
// https://developers.e.qq.com/v3.0/docs/api/batch_requests/add
type BatchRequestAddReq struct {
	GlobalReq
	BatchRequestSpec []*BatchRequestSpec `json:"batch_request_spec"` // 批量调用接口所需条件列表 (必填)
}

func (p *BatchRequestAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建批量请求参数
func (p *BatchRequestAddReq) Validate() error {
	if len(p.BatchRequestSpec) < MinBatchRequestSpecCount {
		return errors.New("batch_request_spec为必填，至少包含1个请求条件")
	}
	for i, spec := range p.BatchRequestSpec {
		if spec == nil {
			return errors.New("batch_request_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("batch_request_spec[" + itoa(i) + "]: " + err.Error())
		}
	}
	return p.GlobalReq.Validate()
}

// BatchResponseHeader 批量响应中的单个 HTTP header
type BatchResponseHeader struct {
	Name  string `json:"name"`  // 单个 header 名
	Value string `json:"value"` // 单个 header 值
}

// BatchResponseItem 批量响应列表项
type BatchResponseItem struct {
	HttpCode int                    `json:"http_code"` // HTTP 返回码
	Headers  []*BatchResponseHeader `json:"headers"`   // HTTP header 列表
	Body     string                 `json:"body"`      // 返回的正文数据
}

// BatchRequestAddResp 创建批量请求响应
// https://developers.e.qq.com/v3.0/docs/api/batch_requests/add
type BatchRequestAddResp struct {
	List []*BatchResponseItem `json:"list"` // 返回信息列表，顺序与 batch_request_spec 一致
}
