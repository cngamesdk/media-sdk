package model

import "errors"

// CreativeUpdateStatusReq 修改创意状态请求
type CreativeUpdateStatusReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`          // 广告主ID，必填
	CreativeId   int64   `json:"creative_id"`            // 广告创意ID，必填；与creative_ids可同时填，总数最多20个
	CreativeIds  []int64 `json:"creative_ids,omitempty"` // 广告创意ID列表，与creative_id可同时填，不得重复，总数最多20个
	PutStatus    int     `json:"put_status"`             // 操作码，必填：1=投放 2=暂停 3=删除
}

func (receiver *CreativeUpdateStatusReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeUpdateStatusReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CreativeId <= 0 && len(receiver.CreativeIds) == 0 {
		err = errors.New("creative_id or creative_ids is required")
		return
	}
	if receiver.PutStatus <= 0 {
		err = errors.New("put_status is empty")
		return
	}
	return
}

// CreativeUpdateStatusError 修改状态错误详情
type CreativeUpdateStatusError struct {
	CreativeId int64  `json:"creative_id"` // 创意ID
	ErrorCode  int    `json:"error_code"`  // 错误码
	ErrorMsg   string `json:"error_msg"`   // 错误信息
}

// CreativeUpdateStatusResp 修改创意状态响应数据（仅data部分）
type CreativeUpdateStatusResp struct {
	CreativeId  int64                       `json:"creative_id"`  // 创意ID
	CreativeIds []int64                     `json:"creative_ids"` // 所有修改状态成功的创意ID
	Errors      []CreativeUpdateStatusError `json:"errors"`       // 错误详情列表
}
