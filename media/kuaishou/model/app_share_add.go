package model

import "errors"

// AppShareAddReq 添加应用共享请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/share/add
type AppShareAddReq struct {
	accessTokenReq
	AdvertiserId       int64   `json:"advertiser_id"`                  // 广告主ID，必填
	AppId              int64   `json:"app_id"`                         // 应用ID，必填
	ShareType          int     `json:"share_type"`                     // 共享类型，必填：0=不共享 1=账号 2=主体
	ShareAdvertiserIds []int64 `json:"share_advertiser_ids,omitempty"` // 要共享的账号ID列表，share_type=1时必填，单次最多200个
	ShareCorpIds       []int64 `json:"share_corp_ids,omitempty"`       // 要共享的主体ID列表，share_type=2时必填
}

var validAppShareTypes = map[int]bool{0: true, 1: true, 2: true}

func (receiver *AppShareAddReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareAddReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if !validAppShareTypes[receiver.ShareType] {
		err = errors.New("share_type must be 0(不共享) 1(账号) 2(主体)")
		return
	}
	if receiver.ShareType == 1 {
		if len(receiver.ShareAdvertiserIds) == 0 {
			err = errors.New("share_advertiser_ids is required when share_type=1")
			return
		}
		if len(receiver.ShareAdvertiserIds) > 200 {
			err = errors.New("share_advertiser_ids must not exceed 200 items")
			return
		}
	}
	if receiver.ShareType == 2 && len(receiver.ShareCorpIds) == 0 {
		err = errors.New("share_corp_ids is required when share_type=2")
		return
	}
	return
}

// AppShareAddResp 添加应用共享响应数据（仅data部分）
type AppShareAddResp struct {
	Result bool `json:"result"` // 是否成功：true=成功 false=失败
}
