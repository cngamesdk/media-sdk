package model

import "errors"

// OperationRecordListReq 账户操作记录信息查询请求
type OperationRecordListReq struct {
	accessTokenReq
	AdvertiserId    int64  `json:"advertiser_id"`              // 广告主ID，必填，在获取 access_token 的时候返回
	OperationType   int    `json:"operation_type,omitempty"`   // 操作类型：1=新增 2=修改/删除
	OperationTarget int    `json:"operation_target,omitempty"` // 操作对象类型：1=账户 2=计划 3=广告组 4=创意 5=视频 6=应用申请 7=人群包 9=程序化创意2.0 10=评论内容 11=托管项目 12=自动规则 13=评论屏蔽信息 15=极速推广项目 20=商品库 21=关键行为 22=商品库商品 24=品牌计划 25=品牌单元 26=品牌创意 28=账户智投 29=线索优化 32=智能创意 33=省心投 34=粉丝票持续投放
	RoleType        int    `json:"role_type,omitempty"`        // 操作人：1=广告主 2=代理商 3=系统 4=管理员 5=市场Api 10=智能交易 11=全自动投放
	Page            int    `json:"page,omitempty"`             // 页码，默认1
	PageSize        int    `json:"page_size,omitempty"`        // 每页数量，默认20，最大500
	StartDate       string `json:"start_date,omitempty"`       // 开始时间，最多可查6个月的操作记录
	EndDate         string `json:"end_date,omitempty"`         // 结束时间
}

func (receiver *OperationRecordListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *OperationRecordListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// OperationRecordContentLog 操作记录日志内容
type OperationRecordContentLog struct {
	FieldName    string `json:"field_name"`    // 字段名称
	OriginalData string `json:"original_data"` // 原始数据
	UpdateData   string `json:"update_data"`   // 更新数据
}

// OperationRecordDetail 操作记录详情
type OperationRecordDetail struct {
	ObjectId        string                      `json:"object_id"`        // 操作对象ID
	OperationType   int                         `json:"operation_type"`   // 操作类型：1=新增 2=修改
	OperationTarget int                         `json:"operation_target"` // 操作对象类型：1=账户 2=计划 3=广告组 4=创意 5=视频 6=应用申请 7=人群包
	RoleType        int                         `json:"role_type"`        // 操作人：1=广告主 2=代理商 3=系统 4=管理员 5=市场Api
	ObjectName      string                      `json:"object_name"`      // 操作对象名称
	OperationTime   string                      `json:"operation_time"`   // 操作时间
	ContentLog      []OperationRecordContentLog `json:"content_log"`      // 日志内容
}

// OperationRecordListResp 账户操作记录信息查询响应数据（仅data部分）
type OperationRecordListResp struct {
	TotalCount int64                   `json:"total_count"` // 总数量
	Details    []OperationRecordDetail `json:"details"`     // 详情列表
}
