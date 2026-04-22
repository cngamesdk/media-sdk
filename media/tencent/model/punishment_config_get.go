package model

// ========== 获取处罚系统配置 ==========
// https://developers.e.qq.com/v3.0/docs/api/punishment_config/get

// PunishmentConfigGetReq 获取处罚系统配置请求（无业务请求参数，仅全局参数）
type PunishmentConfigGetReq struct {
	GlobalReq
}

func (p *PunishmentConfigGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取处罚系统配置请求参数
func (p *PunishmentConfigGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	return nil
}

// PunishmentConfigGetResp 获取处罚系统配置响应
type PunishmentConfigGetResp struct {
	IllegalSceneList []*PunishmentConfigOption    `json:"illegal_scene_list,omitempty"` // 处罚场景
	IllegalNodeList  []*PunishmentConfigOption    `json:"illegal_node_list,omitempty"`  // 违规节点/对象
	ChannelList      []*PunishmentConfigOption    `json:"channel_list,omitempty"`       // 发现渠道
	ActionType       []*PunishmentConfigIntOption `json:"action_type,omitempty"`        // 账户处罚动作
	PunishLevel      []*PunishmentConfigIntOption `json:"punish_level,omitempty"`       // 处罚等级
	SceneList        []*PunishmentConfigIntOption `json:"scene_list,omitempty"`         // 违规原因
}

// PunishmentConfigOption 处罚枚举配置项（支持递归嵌套，最多四层）
type PunishmentConfigOption struct {
	Value    string                    `json:"value,omitempty"`     // 处罚枚举配置值
	Desc     string                    `json:"desc,omitempty"`      // 处罚枚举配置描述
	Level    int                       `json:"level,omitempty"`     // 处罚枚举配置层级，不超过四层
	ParentID string                    `json:"parent_id,omitempty"` // 处罚枚举配置父级值
	Options  []*PunishmentConfigOption `json:"options,omitempty"`   // 下级配置列表
}

// PunishmentConfigIntOption 处罚枚举配置项（整数值类型）
type PunishmentConfigIntOption struct {
	Value int64  `json:"value,omitempty"` // 处罚枚举配置整数值
	Desc  string `json:"desc,omitempty"`  // 处罚枚举配置描述
}
