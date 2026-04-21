package model

import "errors"

// ========== 添加人群授权 ==========
// https://developers.e.qq.com/v3.0/docs/api/audience_grant_relations/add

// 授权类型枚举
const (
	AudienceGrantTypeGrantTypeBusiness = "GRANT_TYPE_BUSINESS" // 授权给商务管家账号
)

// 人群授权范围枚举
const (
	AudienceGrantScopeTypeBusiness = "GRANT_SCOPE_TYPE_BUSINESS" // 授权给 BM 内全部账号
	AudienceGrantScopeTypeAccount  = "GRANT_SCOPE_TYPE_ACCOUNT"  // 授权给指定账号
)

// 人群授权权限类型枚举
const (
	AudienceGrantPermissionTypeTarget  = "GRANT_PERMISSION_TYPE_TARGET"  // 定向使用权限
	AudienceGrantPermissionTypeInsight = "GRANT_PERMISSION_TYPE_INSIGHT" // 洞察使用权限
)

// AudienceGrantRelationsAddReq 添加人群授权请求
type AudienceGrantRelationsAddReq struct {
	GlobalReq
	AccountID      int64      `json:"account_id"`       // 推广帐号 id (必填)
	AudienceIDList []int64    `json:"audience_id_list"` // 人群 id 列表，1-20 个 (必填)
	GrantType      string     `json:"grant_type"`       // 授权类型 (必填)，可选值：GRANT_TYPE_BUSINESS
	GrantSpec      *GrantSpec `json:"grant_spec"`       // 人群授权信息 (必填)
}

func (p *AudienceGrantRelationsAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证添加人群授权请求参数
func (p *AudienceGrantRelationsAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.AudienceIDList) == 0 {
		return errors.New("audience_id_list为必填")
	}
	if len(p.AudienceIDList) > 20 {
		return errors.New("audience_id_list最大长度为20")
	}
	for i, id := range p.AudienceIDList {
		if id < 1 {
			return errors.New("audience_id_list[" + itoa(i) + "]最小值为1")
		}
	}
	if p.GrantType == "" {
		return errors.New("grant_type为必填")
	}
	if p.GrantType != AudienceGrantTypeGrantTypeBusiness {
		return errors.New("grant_type值无效，可选值：GRANT_TYPE_BUSINESS")
	}
	if p.GrantSpec == nil {
		return errors.New("grant_spec为必填")
	}
	if validateErr := p.GrantSpec.Validate(); validateErr != nil {
		return validateErr
	}
	return nil
}

// GrantSpec 人群授权信息
type GrantSpec struct {
	GrantToBusinessSpec *GrantToBusinessSpec `json:"grant_to_business_spec,omitempty"` // 授权给商务管家账号的授权信息
}

func (p *GrantSpec) Validate() error {
	if p.GrantToBusinessSpec == nil {
		return errors.New("grant_spec.grant_to_business_spec为必填")
	}
	return p.GrantToBusinessSpec.Validate()
}

// GrantToBusinessSpec 授权给商务管家账号的授权信息
type GrantToBusinessSpec struct {
	GrantBusinessID         int64                     `json:"grant_business_id"`                   // 商务管家账号 (必填)
	GrantScopeType          string                    `json:"grant_scope_type"`                    // 人群授权范围 (必填)，仅支持 GRANT_SCOPE_TYPE_ACCOUNT
	GrantBusinessPermission *GrantBusinessPermission  `json:"grant_business_permission,omitempty"` // 授权给 BM 内全部账号的权限信息
	GrantAccountPermission  []*GrantAccountPermission `json:"grant_account_permission,omitempty"`  // 授权给指定账号的权限信息
}

func (p *GrantToBusinessSpec) Validate() error {
	if p.GrantBusinessID < 0 {
		return errors.New("grant_business_id最小值为0")
	}
	if p.GrantScopeType == "" {
		return errors.New("grant_scope_type为必填")
	}
	if p.GrantScopeType != AudienceGrantScopeTypeBusiness && p.GrantScopeType != AudienceGrantScopeTypeAccount {
		return errors.New("grant_scope_type值无效，可选值：GRANT_SCOPE_TYPE_BUSINESS, GRANT_SCOPE_TYPE_ACCOUNT")
	}
	if p.GrantScopeType == AudienceGrantScopeTypeAccount && len(p.GrantAccountPermission) == 0 {
		return errors.New("grant_scope_type=GRANT_SCOPE_TYPE_ACCOUNT 时，grant_account_permission为必填")
	}
	for i, perm := range p.GrantAccountPermission {
		if perm == nil {
			return errors.New("grant_account_permission[" + itoa(i) + "]不能为空")
		}
		if validateErr := perm.Validate(i); validateErr != nil {
			return validateErr
		}
	}
	return nil
}

// GrantBusinessPermission 授权给 BM 内全部账号的权限信息
type GrantBusinessPermission struct {
	GrantPermissionTypeList []string `json:"grant_permission_type_list"` // 人群授权权限列表 (必填)
}

// GrantAccountPermission 授权给指定账号的权限信息
type GrantAccountPermission struct {
	AccountID               int64    `json:"account_id"`                 // 广告主帐号 id (必填)
	GrantPermissionTypeList []string `json:"grant_permission_type_list"` // 人群授权权限列表 (必填)
}

func (p *GrantAccountPermission) Validate(idx int) error {
	if p.AccountID == 0 {
		return errors.New("grant_account_permission[" + itoa(idx) + "].account_id为必填")
	}
	if len(p.GrantPermissionTypeList) == 0 {
		return errors.New("grant_account_permission[" + itoa(idx) + "].grant_permission_type_list为必填")
	}
	for _, pt := range p.GrantPermissionTypeList {
		if pt != AudienceGrantPermissionTypeTarget && pt != AudienceGrantPermissionTypeInsight {
			return errors.New("grant_permission_type_list值无效，可选值：GRANT_PERMISSION_TYPE_TARGET, GRANT_PERMISSION_TYPE_INSIGHT")
		}
	}
	return nil
}

// AudienceGrantRelationsAddResp 添加人群授权响应（应答字段为空）
type AudienceGrantRelationsAddResp struct{}
