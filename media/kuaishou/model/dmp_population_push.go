package model

import "errors"

// DmpPopulationPushReq 人群包跨账户推送请求
// https://ad.e.kuaishou.com/rest/openapi/v1/dmp/population/accounts/push
// 注意：人群包状态 status=4(已上线) 才能推送，建议单次推送账户数量<20个
type DmpPopulationPushReq struct {
	accessTokenReq
	AdvertiserId   int64   `json:"advertiser_id"`    // 广告主ID，必填
	OrientationId  int64   `json:"orientation_id"`   // 需要推送的人群包ID，必填，status=4才能推送
	DestAccountIds []int64 `json:"dest_account_ids"` // 要推送的账户ID列表，必填，建议单次<20个
}

func (receiver *DmpPopulationPushReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpPopulationPushReq) Validate() (err error) {
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
	if len(receiver.DestAccountIds) == 0 {
		err = errors.New("dest_account_ids is empty")
		return
	}
	return
}

// DmpPopulationPushResp 人群包跨账户推送响应数据（仅data部分）
type DmpPopulationPushResp struct {
	Success []int64 `json:"success"` // 推送成功的account_ids
	Fail    []int64 `json:"fail"`    // 推送失败的account_ids
}
