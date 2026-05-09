package model

import "errors"

// DmpPopulationOnlineReq 人群包上线请求
// https://ad.e.kuaishou.com/rest/openapi/v1/dmp/population/push
// 注意：只有 status=1(已生效) 或 status=6(上线失败) 的人群包可上线，最多6个同时处于"上线中"
type DmpPopulationOnlineReq struct {
	accessTokenReq
	AdvertiserId  int64 `json:"advertiser_id"`  // 广告主ID，必填
	OrientationId int64 `json:"orientation_id"` // 人群包ID，必填，status=1或6才可上线
}

func (receiver *DmpPopulationOnlineReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationOnlineReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.OrientationId <= 0 {
		err = errors.New("orientation_id is empty")
		return
	}
	return
}

// DmpPopulationOnlineResp 人群包上线响应数据（仅data部分）
// data为字符串，如"上线人群包成功"
type DmpPopulationOnlineResp struct {
	OnlineMsg string // 上线结果信息
}
