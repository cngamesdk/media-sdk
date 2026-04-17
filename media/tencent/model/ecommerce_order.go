package model

import "errors"

// ========== 获取订单数据 ==========
// https://developers.e.qq.com/v3.0/docs/api/ecommerce_order/get

// 订单分页默认值
const (
	EcommerceOrderDefaultPage     = 1
	EcommerceOrderDefaultPageSize = 20
)

// 过滤字段常量
const (
	EcommerceOrderFilterFieldOrderId     = "ecommerce_order_id"     // 订单id
	EcommerceOrderFilterFieldOrderStatus = "ecommerce_order_status" // 订单状态
)

// 过滤条件数组长度限制
const (
	EcommerceOrderMinFilteringCount = 1
	EcommerceOrderMaxFilteringCount = 2
)

// 订单状态枚举
const (
	EcommerceOrderStatusAwaitingOrder = "AWAITING_ORDER" // 待下单
	EcommerceOrderStatusPendingPay    = "PENDING_PAY"    // 待付款
	EcommerceOrderStatusPaid          = "PAID"           // 已付款
	EcommerceOrderStatusShipped       = "SHIPPED"        // 已发货
	EcommerceOrderStatusDelivered     = "DELIVERED"      // 已签收
	EcommerceOrderStatusRefunding     = "REFUNDING"      // 退款中
	EcommerceOrderStatusRefunded      = "REFUNDED"       // 已退款
	EcommerceOrderStatusClosed        = "CLOSED"         // 已关闭
	EcommerceOrderStatusCompleted     = "COMPLETED"      // 已完成
	EcommerceOrderStatusReturning     = "RETURNING"      // 退货中
	EcommerceOrderStatusReturned      = "RETURNED"       // 已退货
	EcommerceOrderStatusPartRefunded  = "PART_REFUNDED"  // 部分退款
	EcommerceOrderStatusPartReturned  = "PART_RETURNED"  // 部分退货
	EcommerceOrderStatusExchanging    = "EXCHANGING"     // 换货中
	EcommerceOrderStatusExchanged     = "EXCHANGED"      // 已换货
	EcommerceOrderStatusPartExchanged = "PART_EXCHANGED" // 部分换货
	EcommerceOrderStatusPendingShip   = "PENDING_SHIP"   // 待发货
	EcommerceOrderStatusConfirmed     = "CONFIRMED"      // 已确认
	EcommerceOrderStatusCancelled     = "CANCELLED"      // 已取消
	EcommerceOrderStatusPendingReview = "PENDING_REVIEW" // 待审核
	EcommerceOrderStatusReviewFailed  = "REVIEW_FAILED"  // 审核失败
	EcommerceOrderStatusReviewPassed  = "REVIEW_PASSED"  // 审核通过
	EcommerceOrderStatusProcessing    = "PROCESSING"     // 处理中
	EcommerceOrderStatusAbnormal      = "ABNORMAL"       // 异常
	EcommerceOrderStatusUnknown       = "UNKNOWN"        // 未知
)

// 过滤操作符常量
const (
	EcommerceOrderOperatorEquals   = "EQUALS"
	EcommerceOrderOperatorIn       = "IN"
	EcommerceOrderOperatorContains = "CONTAINS"
)

// EcommerceOrderGetReq 获取订单数据请求
type EcommerceOrderGetReq struct {
	GlobalReq
	AccountID int64                    `json:"account_id"`          // 广告主帐号id (必填)
	Filtering []*EcommerceOrderFilter  `json:"filtering,omitempty"` // 过滤条件
	Date      string                   `json:"date,omitempty"`      // 日期，格式：YYYY-MM-DD
	DateRange *EcommerceOrderDateRange `json:"date_range"`          // 日期范围 (必填)
	Page      int                      `json:"page,omitempty"`      // 页数，默认值：1
	PageSize  int                      `json:"page_size,omitempty"` // 分页大小，默认值：20
}

// EcommerceOrderDateRange 日期范围
type EcommerceOrderDateRange struct {
	StartDate string `json:"start_date"` // 开始日期，格式：YYYY-MM-DD (必填)
	EndDate   string `json:"end_date"`   // 结束日期，格式：YYYY-MM-DD (必填)
}

// EcommerceOrderFilter 过滤条件
type EcommerceOrderFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)，可选值：ecommerce_order_id, ecommerce_order_status
	Operator string   `json:"operator"` // 操作符 (必填)，可选值：EQUALS, IN, CONTAINS
	Values   []string `json:"values"`   // 字段取值 (必填)
}

func (p *EcommerceOrderGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = EcommerceOrderDefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = EcommerceOrderDefaultPageSize
	}
}

// Validate 验证获取订单数据请求参数
func (p *EcommerceOrderGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证date_range
	if p.DateRange == nil {
		return errors.New("date_range为必填")
	}
	if p.DateRange.StartDate == "" {
		return errors.New("date_range.start_date为必填")
	}
	if len(p.DateRange.StartDate) != 10 {
		return errors.New("date_range.start_date长度必须为10字节")
	}
	if p.DateRange.EndDate == "" {
		return errors.New("date_range.end_date为必填")
	}
	if len(p.DateRange.EndDate) != 10 {
		return errors.New("date_range.end_date长度必须为10字节")
	}

	// 验证date
	if p.Date != "" && len(p.Date) != 10 {
		return errors.New("date长度必须为10字节")
	}

	// 验证filtering
	if len(p.Filtering) > 0 {
		if len(p.Filtering) < EcommerceOrderMinFilteringCount || len(p.Filtering) > EcommerceOrderMaxFilteringCount {
			return errors.New("filtering数组长度必须在1-2之间")
		}
		for _, filter := range p.Filtering {
			if filter.Field == "" {
				return errors.New("filtering.field为必填")
			}
			if filter.Field != EcommerceOrderFilterFieldOrderId && filter.Field != EcommerceOrderFilterFieldOrderStatus {
				return errors.New("filtering.field值无效，可选值：ecommerce_order_id、ecommerce_order_status")
			}
			if filter.Operator == "" {
				return errors.New("filtering.operator为必填")
			}
			if len(filter.Values) == 0 {
				return errors.New("filtering.values为必填")
			}
		}
	}

	return nil
}

// EcommerceOrderGetResp 获取订单数据响应
type EcommerceOrderGetResp struct {
	List     []*EcommerceOrderItem `json:"list,omitempty"`      // 订单列表
	PageInfo *PageInfo             `json:"page_info,omitempty"` // 分页配置信息
}

// EcommerceOrderItem 订单项
type EcommerceOrderItem struct {
	AccountID              int64                  `json:"account_id,omitempty"`               // 广告主帐号id
	EcommerceOrderId       string                 `json:"ecommerce_order_id,omitempty"`       // 订单id
	CustomizedPageName     string                 `json:"customized_page_name,omitempty"`     // 页面名称
	CommodityPackageDetail string                 `json:"commodity_package_detail,omitempty"` // 套餐明细
	Quantity               int64                  `json:"quantity,omitempty"`                 // 数量
	Price                  int64                  `json:"price,omitempty"`                    // 单价，单位为分
	TotalPrice             int64                  `json:"total_price,omitempty"`              // 总价，单位为分
	EcommerceOrderTime     string                 `json:"ecommerce_order_time,omitempty"`     // 下单时间
	EcommerceOrderStatus   string                 `json:"ecommerce_order_status,omitempty"`   // 订单状态
	UserName               string                 `json:"user_name,omitempty"`                // 姓名
	UserPhone              string                 `json:"user_phone,omitempty"`               // 手机号码
	UserProvince           string                 `json:"user_province,omitempty"`            // 下单省份
	UserCity               string                 `json:"user_city,omitempty"`                // 下单城市
	UserArea               string                 `json:"user_area,omitempty"`                // 下单地区
	UserAddress            string                 `json:"user_address,omitempty"`             // 收货地址
	UserIp                 string                 `json:"user_ip,omitempty"`                  // 下单IP
	UserMessage            string                 `json:"user_message,omitempty"`             // 留言
	DestinationUrl         string                 `json:"destination_url,omitempty"`          // 落地页url
	AdgroupId              int64                  `json:"adgroup_id,omitempty"`               // 广告组id
	AdgroupName            string                 `json:"adgroup_name,omitempty"`             // 广告组名称
	FromAccountId          int64                  `json:"from_account_id,omitempty"`          // 来源账户id
	DeliverySpec           *EcommerceDeliverySpec `json:"delivery_spec,omitempty"`            // 物流详情
}

// EcommerceDeliverySpec 物流详情
type EcommerceDeliverySpec struct {
	DeliveryTrackingNumber string                   `json:"delivery_tracking_number,omitempty"` // 快递单号
	ExpressCompany         string                   `json:"express_company,omitempty"`          // 快递公司
	DeliveryInfoList       []*EcommerceDeliveryInfo `json:"delivery_info_list,omitempty"`       // 快递详情
}

// EcommerceDeliveryInfo 快递详情
type EcommerceDeliveryInfo struct {
	DeliveryUpdateTime string `json:"delivery_update_time,omitempty"` // 更新时间
	DeliveryDetail     string `json:"delivery_detail,omitempty"`      // 快递详情
}
