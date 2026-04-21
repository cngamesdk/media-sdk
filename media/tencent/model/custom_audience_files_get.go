package model

import "errors"

// ========== 获取客户人群数据文件 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_files/get

// CustomAudienceFilesGetReq 获取客户人群数据文件请求
type CustomAudienceFilesGetReq struct {
	GlobalReq
	AccountID            int64 `json:"account_id"`                        // 推广帐号 id (必填)
	AudienceID           int64 `json:"audience_id,omitempty"`             // 人群 id，只能是 CUSTOMER_FILE 类人群
	CustomAudienceFileID int64 `json:"custom_audience_file_id,omitempty"` // 数据文件 id
	Page                 int   `json:"page,omitempty"`                    // 当前页码，最小值 1，默认值 1
	PageSize             int   `json:"page_size,omitempty"`               // 分页大小，最小值 1，最大值 100，默认值 10
}

func (p *CustomAudienceFilesGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Validate 验证获取客户人群数据文件请求参数
func (p *CustomAudienceFilesGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Page < 1 {
		return errors.New("page最小值为1")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return errors.New("page_size必须在1-100之间")
	}
	return nil
}

// CustomAudienceFilesGetResp 获取客户人群数据文件响应
type CustomAudienceFilesGetResp struct {
	List     []*CustomAudienceFileItem `json:"list,omitempty"`      // 返回数组列表
	PageInfo *PageInfo                 `json:"page_info,omitempty"` // 分页信息
}

// CustomAudienceFileItem 数据文件信息
type CustomAudienceFileItem struct {
	AudienceID           int64  `json:"audience_id,omitempty"`             // 人群 id
	CustomAudienceFileID int64  `json:"custom_audience_file_id,omitempty"` // 数据文件 id
	Name                 string `json:"name,omitempty"`                    // 文件名称
	UserIDType           string `json:"user_id_type,omitempty"`            // 号码包用户 id 类型
	OperationType        string `json:"operation_type,omitempty"`          // 文件操作类型，APPEND 或 REDUCE
	OpenAppID            string `json:"open_app_id,omitempty"`             // 微信 appid，user_id_type 为 WECHAT_OPENID/WX_OPENID 时有效
	ProcessStatus        string `json:"process_status,omitempty"`          // 文件状态
	ProcessCode          int    `json:"process_code,omitempty"`            // 处理完成后的状态码，0 成功，非 0 失败
	ErrorMessage         string `json:"error_message,omitempty"`           // 错误具体信息
	LineCount            int64  `json:"line_count,omitempty"`              // 文件总行数
	ValidLineCount       int64  `json:"valid_line_count,omitempty"`        // 文件中没有格式错误的行数
	UserCount            int64  `json:"user_count,omitempty"`              // 文件包含的用户数
	Size                 int64  `json:"size,omitempty"`                    // 文件大小
	CreatedTime          string `json:"created_time,omitempty"`            // 创建时间，格式 yyyy-MM-dd HH:mm:ss
}
