package model

import "errors"

// ========== 上传客户人群数据文件 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_files/add

// 号码包用户 id 类型枚举
const (
	CustomAudienceFileUserIDTypeHashIDFA          = "HASH_IDFA"
	CustomAudienceFileUserIDTypeHashIMEI          = "HASH_IMEI"
	CustomAudienceFileUserIDTypeHashMobilePhone   = "HASH_MOBILE_PHONE"
	CustomAudienceFileUserIDTypeIDFA              = "IDFA"
	CustomAudienceFileUserIDTypeIMEI              = "IMEI"
	CustomAudienceFileUserIDTypeWxOpenID          = "WX_OPENID"
	CustomAudienceFileUserIDTypeWxUnionID         = "WX_UNIONID"
	CustomAudienceFileUserIDTypeWechatOpenID      = "WECHAT_OPENID"
	CustomAudienceFileUserIDTypeSaltedHashIMEI    = "SALTED_HASH_IMEI"
	CustomAudienceFileUserIDTypeSaltedHashIDFA    = "SALTED_HASH_IDFA"
	CustomAudienceFileUserIDTypeOAID              = "OAID"
	CustomAudienceFileUserIDTypeHashOAID          = "HASH_OAID"
	CustomAudienceFileUserIDTypeSHA256MobilePhone = "SHA256_MOBILE_PHONE"
	CustomAudienceFileUserIDTypeMD5SHA256IMEI     = "MD5_SHA256_IMEI"
	CustomAudienceFileUserIDTypeMD5SHA256IDFA     = "MD5_SHA256_IDFA"
	CustomAudienceFileUserIDTypeMD5SHA256OAID     = "MD5_SHA256_OAID"
	CustomAudienceFileUserIDTypeCAID              = "CAID"
)

// 文件操作类型枚举
const (
	CustomAudienceFileOperationTypeAppend = "APPEND" // 追加（默认）
	CustomAudienceFileOperationTypeReduce = "REDUCE" // 删减
)

// CustomAudienceFilesAddReq 上传客户人群数据文件请求（multipart/form-data）
type CustomAudienceFilesAddReq struct {
	GlobalReq
	AccountID     int64  // 推广帐号 id (必填)
	AudienceID    int64  // 人群 id，只能是 CUSTOMER_FILE 类人群 (必填)
	UserIDType    string // 号码包用户 id 类型 (必填)
	File          []byte // 上传人群文件二进制内容，txt/csv 每行一号码，或压缩后的 zip 包，zip 不超过 100M (必填)
	FileName      string // 文件名（含扩展名），用于 multipart 表单 (必填)
	OperationType string // 文件操作类型，不填默认 APPEND，可选值：APPEND, REDUCE
	OpenAppID     string // 微信 appid，user_id_type 为 WECHAT_OPENID 或 WX_OPENID 时有效，1-128字节
}

func (p *CustomAudienceFilesAddReq) Format() {
	p.GlobalReq.Format()
	if p.OperationType == "" {
		p.OperationType = CustomAudienceFileOperationTypeAppend
	}
}

// Validate 验证上传客户人群数据文件请求参数
func (p *CustomAudienceFilesAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AudienceID == 0 {
		return errors.New("audience_id为必填")
	}
	if p.UserIDType == "" {
		return errors.New("user_id_type为必填")
	}
	if !isValidCustomAudienceFileUserIDType(p.UserIDType) {
		return errors.New("user_id_type值无效")
	}
	if len(p.File) == 0 {
		return errors.New("file为必填")
	}
	if p.FileName == "" {
		return errors.New("file_name为必填")
	}
	if p.OperationType != CustomAudienceFileOperationTypeAppend && p.OperationType != CustomAudienceFileOperationTypeReduce {
		return errors.New("operation_type值无效，可选值：APPEND, REDUCE")
	}
	if p.OpenAppID != "" && len(p.OpenAppID) > 128 {
		return errors.New("open_app_id长度不能超过128字节")
	}
	return nil
}

func isValidCustomAudienceFileUserIDType(t string) bool {
	validTypes := map[string]bool{
		CustomAudienceFileUserIDTypeHashIDFA:          true,
		CustomAudienceFileUserIDTypeHashIMEI:          true,
		CustomAudienceFileUserIDTypeHashMobilePhone:   true,
		CustomAudienceFileUserIDTypeIDFA:              true,
		CustomAudienceFileUserIDTypeIMEI:              true,
		CustomAudienceFileUserIDTypeWxOpenID:          true,
		CustomAudienceFileUserIDTypeWxUnionID:         true,
		CustomAudienceFileUserIDTypeWechatOpenID:      true,
		CustomAudienceFileUserIDTypeSaltedHashIMEI:    true,
		CustomAudienceFileUserIDTypeSaltedHashIDFA:    true,
		CustomAudienceFileUserIDTypeOAID:              true,
		CustomAudienceFileUserIDTypeHashOAID:          true,
		CustomAudienceFileUserIDTypeSHA256MobilePhone: true,
		CustomAudienceFileUserIDTypeMD5SHA256IMEI:     true,
		CustomAudienceFileUserIDTypeMD5SHA256IDFA:     true,
		CustomAudienceFileUserIDTypeMD5SHA256OAID:     true,
		CustomAudienceFileUserIDTypeCAID:              true,
	}
	return validTypes[t]
}

// CustomAudienceFilesAddResp 上传客户人群数据文件响应
type CustomAudienceFilesAddResp struct {
	CustomAudienceFileID int64 `json:"custom_audience_file_id"` // 数据文件 id
}
