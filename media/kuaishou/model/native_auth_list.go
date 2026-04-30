package model

import "errors"

// AuthListPageInfo 授权列表分页信息
type AuthListPageInfo struct {
	CurrentPage int64 `json:"current_page"` // 当前页面值，必填
	PageSize    int64 `json:"page_size"`    // 页面大小
	TotalCount  int64 `json:"total_count"`  // 总共数目
}

// NativeAuthListReq 获取快手号授权列表请求
type NativeAuthListReq struct {
	accessTokenReq
	AdvertiserId int64            `json:"advertiser_id"` // 广告主id，必填
	PageInfo     AuthListPageInfo `json:"page_info"`     // 分页信息，必填
	KolUserId    int64            `json:"kol_user_id"`   // 快手号用户id查询
	UserName     string           `json:"user_name"`     // 快手号名称查询
	AuthStatus   []int            `json:"auth_status"`   // 授权状态：1-新建待确认 2-已确认未开始 3-生效中 4-拒绝 5-超时 6-过期 7-终止 8-已删除
	KolUserType  []int            `json:"kol_user_type"` // 达人类型：1-普通快手号，2-蓝V服务号，3-聚星作品达人，5-聚星直播达人
}

func (receiver *NativeAuthListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeAuthListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PageInfo.CurrentPage <= 0 {
		err = errors.New("page_info.current_page is empty")
		return
	}
	if len(receiver.KolUserType) == 0 {
		err = errors.New("kol_user_type is empty")
		return
	}
	return
}

// KwaiUserAuthorizeInfo 快手号授权信息
type KwaiUserAuthorizeInfo struct {
	UserId          int64           `json:"user_id"`          // 用户id
	KolStatus       int             `json:"kol_status"`       // 授权状态：1-待达人确认 2-达人已确认 3-授权生效中 4-达人拒绝 5-超时 6-过期 7-终止
	ValidStartTime  int64           `json:"valid_start_time"` // 授权开始时间，不限时为0
	ValidEndTime    int64           `json:"valid_end_time"`   // 授权结束时间，不限时为0
	UserInfo        UserProfileView `json:"user_info"`        // 用户信息
	ValidType       int             `json:"valid_type"`       // 授权类型：1-固定时间，2-不限时间
	KolUserType     int             `json:"kol_user_type"`    // 快手号用户类型：1-普通，2-蓝V，3-聚星达人
	AuthId          int64           `json:"auth_id"`          // 授权记录ID
	ComponentStatus int             `json:"component_status"` // 组件授权状态：0-待审核，1-审核通过，2-审核拒绝，3-未授权，4-已终止
}

// NativeAuthListResp 获取快手号授权列表响应数据（仅data部分）
type NativeAuthListResp struct {
	Data     []KwaiUserAuthorizeInfo `json:"data"`      // 授权列表
	PageInfo AuthListPageInfo        `json:"page_info"` // 分页信息
}
