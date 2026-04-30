package model

import "errors"

// NativeAuthSaveReq 快手号授权请求
type NativeAuthSaveReq struct {
	accessTokenReq
	AdvertiserId     int64    `json:"advertiser_id"`       // 广告主ID，必填
	BatchUserIds     []int64  `json:"batch_user_ids"`      // 批量快手号id，最多10个（普通/蓝V时填写），必填
	UserId           int64    `json:"user_id"`             // 授权对应的快手号id
	ValidType        string   `json:"valid_type"`          // 授权类型：1-固定时间，2-不限时间（蓝V/普通必填），必填
	ValidStartTime   int64    `json:"valid_start_time"`    // 授权开始时间，毫秒时间戳
	ValidEndTime     int64    `json:"valid_end_time"`      // 授权结束时间，毫秒时间戳
	KolUserType      int      `json:"kol_user_type"`       // 达人类型：1-普通快手号，2-蓝V服务号（默认蓝V）
	AdSocialOrderIds []string `json:"ad_social_order_ids"` // 批量聚星订单（聚星直播达人授权时），最多10个
	BatchKwaiIds     []string `json:"batch_kwai_ids"`      // 批量自定义快手号id
	SyncBindType     int      `json:"sync_bind_type"`      // 同步授权子号
}

func (receiver *NativeAuthSaveReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeAuthSaveReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.BatchUserIds) == 0 {
		err = errors.New("batch_user_ids is empty")
		return
	}
	if len(receiver.BatchUserIds) > 10 {
		err = errors.New("batch_user_ids max length is 10")
		return
	}
	if len(receiver.ValidType) == 0 {
		err = errors.New("valid_type is empty")
		return
	}
	return
}

// NativeAuthSaveResult 授权结果
type NativeAuthSaveResult struct {
	AuthId          int64  `json:"auth_id"`            // 授权成功后的唯一id，失败为空
	UserId          int64  `json:"user_id"`            // 授权对应的快手号id
	Message         string `json:"message"`            // 错误信息
	AdSocialOrderId string `json:"ad_social_order_id"` // 聚星订单ID
	KwaiId          string `json:"kwai_id"`            // 授权成功自定义快手号
}

// NativeAuthSaveResp 快手号授权响应数据（仅data部分，返回数组）
type NativeAuthSaveResp []NativeAuthSaveResult
