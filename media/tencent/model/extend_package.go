package model

import (
	"errors"
	"regexp"
)

// ========== 创建应用分包 ==========
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add

// 字段限制常量
const (
	MinExtendPackageChannelListLen           = 1   // channel_list 最小长度
	MaxExtendPackageChannelListLen           = 200 // channel_list 最大长度
	MinExtendPackageChannelIDBytes           = 1   // channel_id 最小长度
	MaxExtendPackageChannelIDBytes           = 200 // channel_id 最大长度
	MinExtendPackageChannelNameBytes         = 1   // channel_name 最小长度
	MaxExtendPackageChannelNameBytes         = 255 // channel_name 最大长度
	MinExtendPackageCustomizedChannelIDBytes = 1   // customized_channel_id 最小长度
	MaxExtendPackageCustomizedChannelIDBytes = 256 // customized_channel_id 最大长度
)

// channelIDRegex channel_id 合法字符正则：仅允许英文字母、数字和 _ . -
var channelIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_.\-]+$`)

// ExtendPackageChannelItem 渠道号信息
type ExtendPackageChannelItem struct {
	ChannelID           string `json:"channel_id"`                      // 渠道标识 (必填)，仅英文/数字/_.-，1-200 字节
	ChannelName         string `json:"channel_name,omitempty"`          // 渠道包名称，1-255 字节，默认 {package_id}_{channel_id}
	CustomizedChannelID string `json:"customized_channel_id,omitempty"` // 自定义渠道包 id，1-256 字节，须与 channel_id 一致
}

// Validate 验证单条渠道号信息
func (c *ExtendPackageChannelItem) Validate() error {
	if c.ChannelID == "" {
		return errors.New("channel_id为必填")
	}
	if len(c.ChannelID) < MinExtendPackageChannelIDBytes || len(c.ChannelID) > MaxExtendPackageChannelIDBytes {
		return errors.New("channel_id长度须在1-200字节之间")
	}
	if !channelIDRegex.MatchString(c.ChannelID) {
		return errors.New("channel_id只能由英文字母、数字和_.-组成")
	}
	if c.ChannelName != "" {
		if len(c.ChannelName) < MinExtendPackageChannelNameBytes || len(c.ChannelName) > MaxExtendPackageChannelNameBytes {
			return errors.New("channel_name长度须在1-255字节之间")
		}
	}
	if c.CustomizedChannelID != "" {
		if len(c.CustomizedChannelID) < MinExtendPackageCustomizedChannelIDBytes || len(c.CustomizedChannelID) > MaxExtendPackageCustomizedChannelIDBytes {
			return errors.New("customized_channel_id长度须在1-256字节之间")
		}
	}
	return nil
}

// ExtendPackageAddReq 创建应用分包请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add
type ExtendPackageAddReq struct {
	GlobalReq
	AccountID   int64                       `json:"account_id"`   // 广告主帐号 id (必填)
	PackageID   int64                       `json:"package_id"`   // Android 应用 id (必填)，≥0 且 <2^63
	ChannelList []*ExtendPackageChannelItem `json:"channel_list"` // 渠道号信息 (必填)，1-200 条
}

func (p *ExtendPackageAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建应用分包请求参数
func (p *ExtendPackageAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PackageID < 0 {
		return errors.New("package_id须大于等于0")
	}
	if len(p.ChannelList) < MinExtendPackageChannelListLen || len(p.ChannelList) > MaxExtendPackageChannelListLen {
		return errors.New("channel_list数组长度须在1-200之间")
	}
	for i, ch := range p.ChannelList {
		if ch == nil {
			return errors.New("channel_list[" + itoa(i) + "]不能为空")
		}
		if err := ch.Validate(); err != nil {
			return errors.New("channel_list[" + itoa(i) + "]: " + err.Error())
		}
	}
	return p.GlobalReq.Validate()
}

// ExtendPackageResultItem 渠道包操作结果（成功/失败列表共用）
type ExtendPackageResultItem struct {
	ChannelName      string `json:"channel_name"`       // 安卓应用渠道包名称
	ChannelPackageID string `json:"channel_package_id"` // 安卓应用渠道包 id
	Message          string `json:"message"`            // 渠道包操作结果描述
}

// ExtendPackageAddResp 创建应用分包响应
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add
type ExtendPackageAddResp struct {
	PackageID      int64                      `json:"package_id"`      // Android 应用 id
	SuccessResults []*ExtendPackageResultItem `json:"success_results"` // 渠道包操作成功信息列表
	FailedResults  []*ExtendPackageResultItem `json:"failed_results"`  // 应用分包操作失败列表
}
