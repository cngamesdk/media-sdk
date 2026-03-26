package model

import "errors"

type AdgroupsGetReq struct {
	GlobalReq
	CursorPageV2Req
	AccountID int64                 `json:"account_id"`           // 广告主帐号id (必填)
	IsDeleted bool                  `json:"is_deleted,omitempty"` // 是否已删除
	Fields    []string              `json:"fields,omitempty"`     // 指定返回的字段列表
	Filtering []*AdgroupQueryFilter `json:"filtering,omitempty"`  // 过滤条件
}

// AdgroupQueryFilter 广告组查询过滤条件
type AdgroupQueryFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// 常量定义 - 过滤字段
const (
	FieldAdgroupId                  = "adgroup_id"
	FieldAdgroupName                = "adgroup_name"
	FieldCreatedTime                = "created_time"
	FieldLastModifiedTime           = "last_modified_time"
	FieldMaterialPackageId          = "material_package_id"
	FieldJointBudgetRuleId          = "joint_budget_rule_id"
	FieldConfiguredStatus           = "configured_status"
	FieldAutoDerivedCreativeEnabled = "auto_derived_creative_enabled"
	FieldSmartDeliveryPlatform      = "smart_delivery_platform"
)

// 常量定义 - 操作符
const (
	OperatorEquals        = "EQUALS"
	OperatorIn            = "IN"
	OperatorContains      = "CONTAINS"
	OperatorLess          = "LESS"
	OperatorLessEquals    = "LESS_EQUALS"
	OperatorGreater       = "GREATER"
	OperatorGreaterEquals = "GREATER_EQUALS"
)

// 常量定义 - 配置状态
const (
	ConfiguredStatusNormal  = "AD_STATUS_NORMAL"  // 正常
	ConfiguredStatusSuspend = "AD_STATUS_SUSPEND" // 暂停
)

// 长度限制常量
const (
	MinFilteringCount  = 1
	MaxFilteringCount  = 255
	MinValuesCount     = 1
	MaxValuesCount     = 100
	CreatedTimeLength  = 10
	ModifiedTimeLength = 10
)

func (p *AdgroupsGetReq) Format() {
	p.GlobalReq.Format()
	p.CursorPageV2Req.Format()
}

// Validate 验证广告组查询参数
func (p *AdgroupsGetReq) Validate() error {
	// 如果过滤条件为空，视为无限制条件
	if len(p.Filtering) == 0 {
		return nil
	}

	// 验证过滤条件数量
	if len(p.Filtering) < MinFilteringCount || len(p.Filtering) > MaxFilteringCount {
		return errors.New("filtering数组长度必须在1-255之间")
	}

	// 验证每个过滤条件
	for i, filter := range p.Filtering {
		if err := filter.Validate(); err != nil {
			return errors.New("filtering[" + string(rune(i)) + "]验证失败: " + err.Error())
		}
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if validateErr := p.CursorPageV2Req.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证fields
	if err := p.validateFields(); err != nil {
		return err
	}

	return nil
}

// validateFields 验证字段列表
func (p *AdgroupsGetReq) validateFields() error {
	if len(p.Fields) == 0 {
		return nil
	}
	if len(p.Fields) < MinAdgroupFieldsCount || len(p.Fields) > MaxAdgroupFieldsCount {
		return errors.New("fields数组长度必须在1-1024之间")
	}
	for _, field := range p.Fields {
		if len(field) < MinAdgroupFieldLength || len(field) > MaxAdgroupFieldLength {
			return errors.New("fields中的字段长度必须在1-64字节之间")
		}
	}
	return nil
}

// Validate 验证单个过滤条件
func (f *AdgroupQueryFilter) Validate() error {
	// 1. 验证字段
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if !isValidField(f.Field) {
		return errors.New("field值无效，请参考文档中的允许值")
	}

	// 2. 验证操作符
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if !isValidOperatorForField(f.Field, f.Operator) {
		return errors.New("operator值无效，当前字段不支持该操作符")
	}

	// 3. 验证values
	if len(f.Values) == 0 {
		return errors.New("values为必填")
	}
	if err := validateValuesForField(f); err != nil {
		return err
	}

	return nil
}

// isValidField 验证字段是否有效
func isValidField(field string) bool {
	validFields := map[string]bool{
		FieldAdgroupId:                  true,
		FieldAdgroupName:                true,
		FieldCreatedTime:                true,
		FieldLastModifiedTime:           true,
		FieldMaterialPackageId:          true,
		FieldJointBudgetRuleId:          true,
		FieldConfiguredStatus:           true,
		FieldAutoDerivedCreativeEnabled: true,
		FieldSmartDeliveryPlatform:      true,
	}
	return validFields[field]
}

// isValidOperatorForField 验证字段支持的操作符
func isValidOperatorForField(field, operator string) bool {
	switch field {
	case FieldAdgroupId:
		return operator == OperatorEquals || operator == OperatorIn
	case FieldAdgroupName:
		return operator == OperatorEquals || operator == OperatorContains
	case FieldCreatedTime, FieldLastModifiedTime:
		return operator == OperatorEquals || operator == OperatorLessEquals ||
			operator == OperatorLess || operator == OperatorGreaterEquals ||
			operator == OperatorGreater
	case FieldMaterialPackageId:
		return operator == OperatorEquals || operator == OperatorLessEquals ||
			operator == OperatorLess || operator == OperatorGreaterEquals ||
			operator == OperatorGreater
	case FieldJointBudgetRuleId:
		return operator == OperatorEquals || operator == OperatorIn
	case FieldConfiguredStatus:
		return operator == OperatorEquals
	case FieldAutoDerivedCreativeEnabled:
		return operator == OperatorEquals
	case FieldSmartDeliveryPlatform:
		return operator == OperatorEquals || operator == OperatorIn || operator == OperatorGreaterEquals
	default:
		return false
	}
}

// validateValuesForField 验证字段取值
func validateValuesForField(f *AdgroupQueryFilter) error {
	switch f.Field {
	case FieldAdgroupId:
		return validateAdgroupIdValues(f)
	case FieldAdgroupName:
		return validateAdgroupNameValues(f)
	case FieldCreatedTime, FieldLastModifiedTime:
		return validateTimeValues(f)
	case FieldMaterialPackageId:
		return validateMaterialPackageIdValues(f)
	case FieldJointBudgetRuleId:
		return validateJointBudgetRuleIdValues(f)
	case FieldConfiguredStatus:
		return validateConfiguredStatusValues(f)
	case FieldAutoDerivedCreativeEnabled:
		return validateAutoDerivedCreativeEnabledValues(f)
	case FieldSmartDeliveryPlatform:
		return validateSmartDeliveryPlatformValues(f)
	}
	return nil
}

// validateAdgroupIdValues 验证广告组ID取值
func validateAdgroupIdValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
	case OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
	}
	return nil
}

// validateAdgroupNameValues 验证广告组名称取值
func validateAdgroupNameValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	for _, v := range f.Values {
		if len(v) < 1 || len(v) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	}
	return nil
}

// validateTimeValues 验证时间取值
func validateTimeValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	for _, v := range f.Values {
		if len(v) != CreatedTimeLength {
			return errors.New("时间字段长度必须为10字节")
		}
	}
	return nil
}

// validateMaterialPackageIdValues 验证素材包ID取值
func validateMaterialPackageIdValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	return nil
}

// validateJointBudgetRuleIdValues 验证联合预算规则ID取值
func validateJointBudgetRuleIdValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals, OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("values数组长度必须在1-100之间")
		}
	}
	return nil
}

// validateConfiguredStatusValues 验证配置状态取值
func validateConfiguredStatusValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	if f.Values[0] != ConfiguredStatusNormal && f.Values[0] != ConfiguredStatusSuspend {
		return errors.New("configured_status值无效，允许值：AD_STATUS_NORMAL、AD_STATUS_SUSPEND")
	}
	return nil
}

// validateAutoDerivedCreativeEnabledValues 验证自动衍生创意启用状态取值
func validateAutoDerivedCreativeEnabledValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
		return errors.New("字段长度必须在1-180字节之间")
	}
	return nil
}

// validateSmartDeliveryPlatformValues 验证智能投放平台取值
func validateSmartDeliveryPlatformValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
		if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	case OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
		for _, v := range f.Values {
			if len(v) < 1 || len(v) > MaxFieldLength {
				return errors.New("字段长度必须在1-180字节之间")
			}
		}
	case OperatorGreaterEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为GREATER_EQUALS时，values数组长度必须为1")
		}
		if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	}
	return nil
}

// 分页限制常量
const (
	DefaultPage     = 1
	DefaultPageSize = 10

	MinAdgroupFieldsCount = 1
	MaxAdgroupFieldsCount = 1024
	MinAdgroupFieldLength = 1
	MaxAdgroupFieldLength = 64

	MaxCursorLength = 10
)
