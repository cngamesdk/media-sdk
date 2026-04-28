package model

import "errors"

// UnitUpdateBidReq 修改广告组出价请求
type UnitUpdateBidReq struct {
	accessTokenReq
	AdvertiserId      int64   `json:"advertiser_id"`                 // 广告主ID，必填
	UnitId            int64   `json:"unit_id,omitempty"`             // 广告组ID，与unit_ids至少填一个
	UnitIds           []int64 `json:"unit_ids,omitempty"`            // 广告组ID列表，可批量修改成相同出价
	Bid               int64   `json:"bid"`                           // 出价，必填，单位：厘；CPC/eCPC不低于0.2元不高于100元；OCPC行为出价不低于1元
	DeepConversionBid int64   `json:"deep_conversion_bid,omitempty"` // 深度转化目标出价，单位：厘
	RoiRatio          float64 `json:"roi_ratio,omitempty"`           // 付费ROI系数，优化目标为首日ROI时必填，范围(0,100]，最多三位小数
}

func (receiver *UnitUpdateBidReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitUpdateBidReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 && len(receiver.UnitIds) == 0 {
		err = errors.New("unit_id or unit_ids is required")
		return
	}
	if receiver.Bid <= 0 {
		err = errors.New("bid is empty")
		return
	}
	return
}

// UnitUpdateBidError 修改出价错误详情
type UnitUpdateBidError struct {
	Id       int64  `json:"id"`        // ID
	ErrorMsg string `json:"error_msg"` // 错误信息
}

// UnitUpdateBidResp 修改广告组出价响应数据（仅data部分）
type UnitUpdateBidResp struct {
	UnitIds []int64              `json:"unit_ids"` // 所有修改成功的广告组ID
	Errors  []UnitUpdateBidError `json:"errors"`   // 错误详情列表
}
