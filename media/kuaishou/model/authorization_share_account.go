package model

import "errors"

// AuthorizationShareAccountReq 获取已共享授权记录请求
type AuthorizationShareAccountReq struct {
	accessTokenReq
	KolUserId    int64 `json:"kol_user_id"`   // 达人ID，必填
	AdvertiserId int64 `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *AuthorizationShareAccountReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuthorizationShareAccountReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.KolUserId <= 0 {
		err = errors.New("kol_user_id is empty")
		return
	}
	return
}

// ShareAccountAuthInfo 已共享授权记录
type ShareAccountAuthInfo struct {
	UserId          int64           `json:"user_id"`          // 用户id
	KolStatus       int             `json:"kol_status"`       // 快手号授权状态：1-待达人确认 2-已确认 3-生效中 4-拒绝 5-超时 6-过期 7-终止
	ValidStartTime  int64           `json:"valid_start_time"` // 授权开始时间，不限时为0
	ValidEndTime    int64           `json:"valid_end_time"`   // 授权结束时间，不限时为0
	UserInfo        UserProfileView `json:"user_info"`        // 用户信息
	ValidType       int             `json:"valid_type"`       // 授权类型：1-固定时间，2-不限时间
	KolUserType     int             `json:"kol_user_type"`    // 达人类型：1-普通，2-蓝V服务号
	AuthId          int64           `json:"auth_id"`          // 共享授权记录ID
	AccountId       int64           `json:"account_id"`       // 账户ID
	ComponentStatus int             `json:"component_status"` // 组件授权状态：0-待审核，1-通过，2-拒绝，3-未授权，4-已终止
}

// AuthorizationShareAccountResp 获取已共享授权记录响应数据（仅data部分，返回数组）
type AuthorizationShareAccountResp []ShareAccountAuthInfo
