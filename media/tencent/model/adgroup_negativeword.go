package model

import "errors"

// ========== 新增广告否定词 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/add

// 常量定义 - 操作执行状态
const (
	NegativewordOperSuccess = "OPER_SUCCESS" // 全部成功
	NegativewordOperFail    = "OPER_FAIL"    // 存在操作失败
)

// 字段限制常量
const (
	MaxNegativeWordsCount = 900 // 否定词数组最大长度
	MaxNegativeWordBytes  = 150 // 单个否定词最大字节数（字段长度最大150字节）
	MinNegativeWordBytes  = 1   // 单个否定词最小字节数
)

// NegativeWordGroup 否定词分组（请求/响应共用）
type NegativeWordGroup struct {
	PhraseNegativeWords []string `json:"phrase_negative_words"` // 短语否定词列表
	ExactNegativeWords  []string `json:"exact_negative_words"`  // 精确否定词列表
}

// AdgroupNegativewordAddReq 新增广告否定词请求
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/add
type AdgroupNegativewordAddReq struct {
	GlobalReq
	AccountID           int64    `json:"account_id"`            // 广告主帐号 id (必填)
	AdgroupID           int64    `json:"adgroup_id"`            // 广告 id (必填)
	PhraseNegativeWords []string `json:"phrase_negative_words"` // 短语否定词 (必填)，数组最大长度900，单词最大150字节
	ExactNegativeWords  []string `json:"exact_negative_words"`  // 精确否定词 (必填)，数组最大长度900，单词最大150字节
}

func (p *AdgroupNegativewordAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证新增广告否定词请求参数
func (p *AdgroupNegativewordAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if len(p.PhraseNegativeWords) == 0 && len(p.ExactNegativeWords) == 0 {
		return errors.New("phrase_negative_words和exact_negative_words不能同时为空")
	}
	if len(p.PhraseNegativeWords) > MaxNegativeWordsCount {
		return errors.New("phrase_negative_words数组长度不能超过900")
	}
	if len(p.ExactNegativeWords) > MaxNegativeWordsCount {
		return errors.New("exact_negative_words数组长度不能超过900")
	}
	for i, word := range p.PhraseNegativeWords {
		if len(word) < MinNegativeWordBytes || len(word) > MaxNegativeWordBytes {
			return errors.New("phrase_negative_words[" + itoa(i) + "]长度必须在1-150字节之间")
		}
	}
	for i, word := range p.ExactNegativeWords {
		if len(word) < MinNegativeWordBytes || len(word) > MaxNegativeWordBytes {
			return errors.New("exact_negative_words[" + itoa(i) + "]长度必须在1-150字节之间")
		}
	}
	return p.GlobalReq.Validate()
}

// AdgroupNegativewordAddResp 新增广告否定词响应
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/add
type AdgroupNegativewordAddResp struct {
	AdgroupID         int64              `json:"adgroup_id"`          // 广告 id
	Status            string             `json:"status"`              // 操作执行状态：OPER_FAIL / OPER_SUCCESS
	DuplicateWords    *NegativeWordGroup `json:"duplicate_words"`     // 因重复导致失败的否定词列表
	ExceedLengthWords *NegativeWordGroup `json:"exceed_length_words"` // 因单词长度超限导致失败的否定词列表
	ExceedLimitWords  *NegativeWordGroup `json:"exceed_limit_words"`  // 因超过个数限制导致失败的否定词列表
	HasSpecialWords   *NegativeWordGroup `json:"has_special_words"`   // 因含有特殊字符导致失败的否定词列表
	SuccessWords      *NegativeWordGroup `json:"success_words"`       // 操作成功的否定词列表
}

// ========== 更新广告否定词 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/update

// AdgroupNegativewordUpdateReq 更新广告否定词请求
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/update
type AdgroupNegativewordUpdateReq struct {
	GlobalReq
	AccountID           int64    `json:"account_id"`            // 广告主帐号 id (必填)
	AdgroupID           int64    `json:"adgroup_id"`            // 广告 id (必填)
	PhraseNegativeWords []string `json:"phrase_negative_words"` // 短语否定词 (必填)，数组最大长度900，单词最大150字节
	ExactNegativeWords  []string `json:"exact_negative_words"`  // 精确否定词 (必填)，数组最大长度900，单词最大150字节
}

func (p *AdgroupNegativewordUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新广告否定词请求参数
func (p *AdgroupNegativewordUpdateReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if len(p.PhraseNegativeWords) == 0 && len(p.ExactNegativeWords) == 0 {
		return errors.New("phrase_negative_words和exact_negative_words不能同时为空")
	}
	if len(p.PhraseNegativeWords) > MaxNegativeWordsCount {
		return errors.New("phrase_negative_words数组长度不能超过900")
	}
	if len(p.ExactNegativeWords) > MaxNegativeWordsCount {
		return errors.New("exact_negative_words数组长度不能超过900")
	}
	for i, word := range p.PhraseNegativeWords {
		if len(word) < MinNegativeWordBytes || len(word) > MaxNegativeWordBytes {
			return errors.New("phrase_negative_words[" + itoa(i) + "]长度必须在1-150字节之间")
		}
	}
	for i, word := range p.ExactNegativeWords {
		if len(word) < MinNegativeWordBytes || len(word) > MaxNegativeWordBytes {
			return errors.New("exact_negative_words[" + itoa(i) + "]长度必须在1-150字节之间")
		}
	}
	return p.GlobalReq.Validate()
}

// AdgroupNegativewordUpdateResp 更新广告否定词响应
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/update
type AdgroupNegativewordUpdateResp struct {
	AdgroupID         int64              `json:"adgroup_id"`          // 广告 id
	Status            string             `json:"status"`              // 操作执行状态：OPER_FAIL / OPER_SUCCESS
	DuplicateWords    *NegativeWordGroup `json:"duplicate_words"`     // 因重复导致失败的否定词列表
	ExceedLengthWords *NegativeWordGroup `json:"exceed_length_words"` // 因单词长度超限导致失败的否定词列表
	ExceedLimitWords  *NegativeWordGroup `json:"exceed_limit_words"`  // 因超过个数限制导致失败的否定词列表
	HasSpecialWords   *NegativeWordGroup `json:"has_special_words"`   // 因含有特殊字符导致失败的否定词列表
	SuccessWords      *NegativeWordGroup `json:"success_words"`       // 操作成功的否定词列表
}

// ========== 查询广告否定词 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/get

// 字段限制常量
const (
	MinAdgroupNegativewordGetIDsCount = 1   // adgroup_ids 最小长度
	MaxAdgroupNegativewordGetIDsCount = 100 // adgroup_ids 最大长度
)

// AdgroupNegativewordGetReq 查询广告否定词请求
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/get
type AdgroupNegativewordGetReq struct {
	GlobalReq
	AccountID  int64   `json:"account_id"`  // 广告主帐号 id (必填)
	AdgroupIDs []int64 `json:"adgroup_ids"` // 广告 id 列表 (必填)，1-100
}

func (p *AdgroupNegativewordGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证查询广告否定词请求参数
func (p *AdgroupNegativewordGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.AdgroupIDs) < MinAdgroupNegativewordGetIDsCount {
		return errors.New("adgroup_ids为必填，至少包含1个广告id")
	}
	if len(p.AdgroupIDs) > MaxAdgroupNegativewordGetIDsCount {
		return errors.New("adgroup_ids数组长度不能超过100")
	}
	return p.GlobalReq.Validate()
}

// AdgroupNegativewordGetListItem 广告否定词列表项
type AdgroupNegativewordGetListItem struct {
	AdgroupID           int64    `json:"adgroup_id"`            // 广告 id
	PhraseNegativeWords []string `json:"phrase_negative_words"` // 短语否定词列表
	ExactNegativeWords  []string `json:"exact_negative_words"`  // 精确否定词列表
}

// AdgroupNegativewordGetResp 查询广告否定词响应
// https://developers.e.qq.com/v3.0/docs/api/adgroup_negativewords/get
type AdgroupNegativewordGetResp struct {
	AdgroupErrorList []string                          `json:"adgroup_error_list"` // 请求失败的广告 id 列表
	AdgroupList      []*AdgroupNegativewordGetListItem `json:"adgroup_list"`       // 广告否定词列表
}
