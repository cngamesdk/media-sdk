package model

import "errors"

// ========== 获取创意组件 ==========
// https://developers.e.qq.com/v3.0/docs/api/components/get

// 常量定义 - 组件过滤字段
const (
	ComponentFieldComponentID            = "component_id"
	ComponentFieldComponentType          = "component_type"
	ComponentFieldComponentSubType       = "component_sub_type"
	ComponentFieldCreatedTime            = "created_time"
	ComponentFieldLastModifiedTime       = "last_modified_time"
	ComponentFieldGenerationType         = "generation_type"
	ComponentFieldPotentialStatus        = "potential_status"
	ComponentFieldVideoID                = "video_id"
	ComponentFieldImageID                = "image_id"
	ComponentFieldVideoSignature         = "video_signature"
	ComponentFieldImageSignature         = "image_signature"
	ComponentFieldFirstPublicationStatus = "first_publication_status"
	ComponentFieldSimilarityStatus       = "similarity_status"
	ComponentFieldScene                  = "scene"
)

// 常量定义 - 共享组件读取方式
const (
	ComponentIDFilteringModeSharingByAgencyBusinessUnit   = "SHARING_BY_AGENCY_BUSINESS_UNIT"   // 代理商业务单元共享
	ComponentIDFilteringModeSharingByCustomerBusinessUnit = "SHARING_BY_CUSTOMER_BUSINESS_UNIT" // 广告主业务单元共享
)

// 常量定义 - 组件生成类型
const (
	ComponentGenerationTypeManual = "GENERATION_TYPE_MANUAL" // 手动
	ComponentGenerationTypeAuto   = "GENERATION_TYPE_AUTO"   // 自动
)

// 常量定义 - 组件潜力状态
const (
	ComponentPotentialStatusHigh       = "POTENTIAL_STATUS_HIGH"        // 高潜力
	ComponentPotentialStatusLow        = "POTENTIAL_STATUS_LOW"         // 低潜力
	ComponentPotentialStatusNoJudgment = "POTENTIAL_STATUS_NO_JUDGMENT" // 无判断
)

// 常量定义 - 组件首次发布状态
const (
	ComponentFirstPublicationStatusYes = "FIRST_PUBLICATION_YES" // 首次发布
	ComponentFirstPublicationStatusNo  = "FIRST_PUBLICATION_NO"  // 非首次发布
)

// 常量定义 - 相似度检测状态
const (
	ComponentSimilarityStatusPass    = "SIMILARITY_STATUS_PASS"    // 通过
	ComponentSimilarityStatusReject  = "SIMILARITY_STATUS_REJECT"  // 拒绝
	ComponentSimilarityStatusPending = "SIMILARITY_STATUS_PENDING" // 待检测
)

// 字段限制常量
const (
	MaxComponentFilteringCount = 10
	DefaultComponentPage       = 1
	DefaultComponentPageSize   = 10
)

// ComponentsGetReq 获取创意组件请求
// https://developers.e.qq.com/v3.0/docs/api/components/get
type ComponentsGetReq struct {
	GlobalReq
	AccountID                int64                   `json:"account_id"`                            // 广告主帐号id (必填)
	OrganizationID           int64                   `json:"organization_id,omitempty"`             // 业务单元id
	Filtering                []*ComponentQueryFilter `json:"filtering,omitempty"`                   // 过滤条件，数组长度0-10
	Page                     int                     `json:"page,omitempty"`                        // 搜索页码，默认1，最大100
	PageSize                 int                     `json:"page_size,omitempty"`                   // 每页条数，默认10，最大100000
	IsDeleted                bool                    `json:"is_deleted,omitempty"`                  // 是否已删除
	Fields                   []string                `json:"fields,omitempty"`                      // 指定返回的字段列表
	ComponentIDFilteringMode string                  `json:"component_id_filtering_mode,omitempty"` // 共享组件读取方式
}

// ComponentQueryFilter 组件查询过滤条件
type ComponentQueryFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

func (p *ComponentsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DefaultComponentPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultComponentPageSize
	}
}

// Validate 验证获取创意组件请求参数
func (p *ComponentsGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if err := p.GlobalReq.Validate(); err != nil {
		return err
	}
	if p.Page < MinPage || p.Page > MaxPage {
		return errors.New("page必须在1-100之间")
	}
	if p.PageSize < MinPageSize || p.PageSize > 100000 {
		return errors.New("page_size必须在1-100000之间")
	}
	if p.ComponentIDFilteringMode != "" &&
		p.ComponentIDFilteringMode != ComponentIDFilteringModeSharingByAgencyBusinessUnit &&
		p.ComponentIDFilteringMode != ComponentIDFilteringModeSharingByCustomerBusinessUnit {
		return errors.New("component_id_filtering_mode值无效")
	}
	return validateComponentFiltering(p.Filtering)
}

// validateComponentFiltering 验证组件过滤条件
func validateComponentFiltering(filtering []*ComponentQueryFilter) error {
	if len(filtering) == 0 {
		return nil
	}
	if len(filtering) > MaxComponentFilteringCount {
		return errors.New("filtering数组长度不能超过10")
	}
	for _, f := range filtering {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate 验证单个组件过滤条件
func (f *ComponentQueryFilter) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if !isValidComponentField(f.Field) {
		return errors.New("field值无效，请参考文档中的允许值")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if !isValidComponentOperatorForField(f.Field, f.Operator) {
		return errors.New("operator值无效，当前字段不支持该操作符")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填")
	}
	return validateComponentValuesForField(f)
}

// isValidComponentField 验证组件过滤字段是否有效
func isValidComponentField(field string) bool {
	validFields := map[string]bool{
		ComponentFieldComponentID:            true,
		ComponentFieldComponentType:          true,
		ComponentFieldComponentSubType:       true,
		ComponentFieldCreatedTime:            true,
		ComponentFieldLastModifiedTime:       true,
		ComponentFieldGenerationType:         true,
		ComponentFieldPotentialStatus:        true,
		ComponentFieldVideoID:                true,
		ComponentFieldImageID:                true,
		ComponentFieldVideoSignature:         true,
		ComponentFieldImageSignature:         true,
		ComponentFieldFirstPublicationStatus: true,
		ComponentFieldSimilarityStatus:       true,
		ComponentFieldScene:                  true,
	}
	return validFields[field]
}

// isValidComponentOperatorForField 验证组件字段支持的操作符
func isValidComponentOperatorForField(field, operator string) bool {
	switch field {
	case ComponentFieldComponentID:
		return operator == OperatorEquals || operator == OperatorIn
	case ComponentFieldComponentType, ComponentFieldComponentSubType,
		ComponentFieldGenerationType, ComponentFieldPotentialStatus,
		ComponentFieldFirstPublicationStatus, ComponentFieldSimilarityStatus,
		ComponentFieldScene:
		return operator == OperatorEquals || operator == OperatorIn
	case ComponentFieldCreatedTime, ComponentFieldLastModifiedTime:
		return operator == OperatorEquals || operator == OperatorLess ||
			operator == OperatorLessEquals || operator == OperatorGreater ||
			operator == OperatorGreaterEquals
	case ComponentFieldVideoID, ComponentFieldImageID,
		ComponentFieldVideoSignature, ComponentFieldImageSignature:
		return operator == OperatorEquals || operator == OperatorIn
	default:
		return false
	}
}

// validateComponentValuesForField 验证组件字段取值
func validateComponentValuesForField(f *ComponentQueryFilter) error {
	switch f.Field {
	case ComponentFieldComponentID:
		if f.Operator == OperatorEquals && len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
		if f.Operator == OperatorIn && (len(f.Values) < MinValuesCount || len(f.Values) > 100) {
			return errors.New("component_id IN操作时，values数组长度必须在1-100之间")
		}
	case ComponentFieldCreatedTime, ComponentFieldLastModifiedTime:
		if len(f.Values) != 1 {
			return errors.New("时间字段values数组长度必须为1")
		}
		if len(f.Values[0]) != CreatedTimeLength {
			return errors.New("时间字段长度必须为10字节")
		}
	case ComponentFieldVideoID, ComponentFieldImageID,
		ComponentFieldVideoSignature, ComponentFieldImageSignature:
		if f.Operator == OperatorIn && (len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount) {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
		if f.Operator == OperatorEquals && len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
	}
	return nil
}

// ComponentsGetListItem 创意组件列表项
type ComponentsGetListItem struct {
	AccountID              int64               `json:"account_id"`                         // 广告主帐号id
	OrganizationID         int64               `json:"organization_id,omitempty"`          // 业务单元id
	ComponentID            int64               `json:"component_id"`                       // 创意组件id
	ComponentValue         *CreativeComponents `json:"component_value,omitempty"`          // 创意组件内容
	CreatedTime            int64               `json:"created_time,omitempty"`             // 创建时间，时间戳
	LastModifiedTime       int64               `json:"last_modified_time,omitempty"`       // 最后修改时间，时间戳
	ComponentSubType       string              `json:"component_sub_type,omitempty"`       // 创意组件子类型
	ComponentCustomName    string              `json:"component_custom_name,omitempty"`    // 创意组件自定义名称
	GenerationType         string              `json:"generation_type,omitempty"`          // 创意组件生成类型
	IsDeleted              bool                `json:"is_deleted,omitempty"`               // 是否已删除
	SimilarityStatus       string              `json:"similarity_status,omitempty"`        // 相似度检测状态
	PotentialStatus        string              `json:"potential_status,omitempty"`         // 组件潜力
	DisableMessage         string              `json:"disable_message,omitempty"`          // 不可用错误信息
	FirstPublicationStatus string              `json:"first_publication_status,omitempty"` // 组件首次发布状态
	Scene                  string              `json:"scene,omitempty"`                    // 创意组件适用场景
}

// ComponentsGetResp 获取创意组件响应
type ComponentsGetResp struct {
	List []*ComponentsGetListItem `json:"list,omitempty"` // 创意组件列表
	PageInfoContainer
}
