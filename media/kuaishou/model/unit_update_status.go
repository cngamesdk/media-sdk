package model

import "errors"

// UnitUpdateStatusReq 修改广告组状态请求
type UnitUpdateStatusReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`      // 广告主ID，必填
	UnitId       int64   `json:"unit_id,omitempty"`  // 广告组ID，与unit_ids至少填一个，最多10个
	UnitIds      []int64 `json:"unit_ids,omitempty"` // 广告组ID列表，与unit_id可同时填，总数最多10个
	PutStatus    int     `json:"put_status"`         // 操作码，必填：1=投放 2=暂停 3=删除
}

func (receiver *UnitUpdateStatusReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitUpdateStatusReq) Validate() (err error) {
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
	if receiver.PutStatus <= 0 {
		err = errors.New("put_status is empty")
		return
	}
	return
}

// UnitUpdateStatusError 修改状态错误详情
type UnitUpdateStatusError struct {
	Id       int64  `json:"id"`        // ID
	ErrorMsg string `json:"error_msg"` // 错误信息
}

// UnitUpdateStatusResp 修改广告组状态响应数据（仅data部分）
type UnitUpdateStatusResp struct {
	UnitIds []int64                 `json:"unit_ids"` // 所有修改状态成功的广告组ID
	Errors  []UnitUpdateStatusError `json:"errors"`   // 错误详情列表
	UnitId  int64                   `json:"unit_id"`  // 广告组ID
}
