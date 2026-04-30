package model

import "errors"

// NativeReportUserListReq 查询原生快手号列表请求
type NativeReportUserListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *NativeReportUserListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeReportUserListReq) Validate() (err error) {
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

// UserProfileView 原生快手号信息
type UserProfileView struct {
	UserId      int64  `json:"user_id"`       // 用户id
	UserName    string `json:"user_name"`     // 用户名称
	UserSex     string `json:"user_sex"`      // 用户性别：M-男性，F-女性
	HeadUrl     string `json:"head_url"`      // 用户头像
	KolUserType int    `json:"kol_user_type"` // 达人用户类型：2-服务号达人，3-聚星达人
}

// NativeReportUserListResp 查询原生快手号列表响应数据（仅data部分）
type NativeReportUserListResp struct {
	UserList []UserProfileView `json:"user_list"` // 用户列表
}
