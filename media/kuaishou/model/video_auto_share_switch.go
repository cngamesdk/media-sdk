package model

import "errors"

// VideoAutoShareSwitchReq 查询账号共享视频库按钮是否开启请求
type VideoAutoShareSwitchReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *VideoAutoShareSwitchReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoAutoShareSwitchReq) Validate() (err error) {
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

// VideoAutoShareSwitchResp 查询账号共享视频库按钮是否开启响应数据（仅data部分）
type VideoAutoShareSwitchResp struct {
	SwitchStatus bool  `json:"switch_status"` // 开关状态：true开启，false关闭
	UserId       int64 `json:"user_id"`       // 所属的user_id
}
