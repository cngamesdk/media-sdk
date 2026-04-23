package model

import "errors"

// CompassAdvertisersReq 获取罗盘绑定广告主列表请求
type CompassAdvertisersReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 罗盘账户的「快手 id」，只能传递授权时使用的快手 id
}

func (receiver *CompassAdvertisersReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CompassAdvertisersReq) Validate() (err error) {
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

// CompassAdvertiserItem 罗盘绑定广告主信息
type CompassAdvertiserItem struct {
	AdvertiserId    int64  `json:"advertiser_id"`    // 广告主ID
	AdvertiserName  string `json:"advertiser_name"`  // 广告主名称
	CorporationName string `json:"corporation_name"` // 公司名称
	AppId           int    `json:"app_id"`           // 账户所属业务线
	CreateSource    int    `json:"create_source"`    // 账户创建来源
	ReviewStatus    int    `json:"review_status"`    // 账户审核状态
	CreateTime      int64  `json:"create_time"`      // 账户创建时间
	UserId          int64  `json:"user_id"`          // 快手ID
	ProductName     string `json:"product_name"`     // 产品名
	AgentId         int64  `json:"agent_id"`         // 代理商ID
	AccountType     int    `json:"account_type"`     // 投放类型
	FrozenStatus    int    `json:"frozen_status"`    // 冻结状态：1=正常 2=冻结
	FrozenReason    string `json:"frozen_reason"`    // 冻结原因
	AuthStatus      int    `json:"auth_status"`      // 认证状态
	AuthDetail      string `json:"auth_detail"`      // 认证详情
	UidBanned       bool   `json:"uid_banned"`       // 快手ID是否被封禁
	ReviewDetail    string `json:"review_detail"`    // 审核详情
}

// CompassAdvertisersResp 获取罗盘绑定广告主列表响应数据（仅data部分）
type CompassAdvertisersResp struct {
	Details []CompassAdvertiserItem `json:"details"` // 广告主列表
}
