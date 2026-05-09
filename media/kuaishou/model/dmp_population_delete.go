package model

import "errors"

// DmpPopulationDeleteReq 人群包删除请求
// https://ad.e.kuaishou.com/rest/openapi/v1/dmp/population/delete
type DmpPopulationDeleteReq struct {
	accessTokenReq
	AdvertiserId  int64 `json:"advertiser_id"`  // 广告主ID，必填
	OrientationId int64 `json:"orientation_id"` // 人群包ID，必填
}

func (receiver *DmpPopulationDeleteReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationDeleteReq) Validate() (err error) {
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

// DmpPopulationDeleteResp 人群包删除响应数据（仅data部分）
// data为字符串，如"删除人群包成功"
type DmpPopulationDeleteResp struct {
	DeleteMsg string // 删除结果信息
}
