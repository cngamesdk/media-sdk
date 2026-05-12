package model

import "errors"

// AppOfflineAppstoresReq 应用商店上下架请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/offline/appstores
type AppOfflineAppstoresReq struct {
	accessTokenReq
	AdvertiserId  int64    `json:"advertiser_id"`  // 广告主ID，必填
	AppIds        []int64  `json:"app_ids"`        // 应用ID列表，必填
	OfflineStores []string `json:"offline_stores"` // 应用商店列表，必填
}

func (receiver *AppOfflineAppstoresReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppOfflineAppstoresReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.AppIds) == 0 {
		err = errors.New("app_ids is empty")
		return
	}
	if len(receiver.OfflineStores) == 0 {
		err = errors.New("offline_stores is empty")
		return
	}
	return
}

// AppOfflineAppstoresResp 应用商店上下架响应数据（仅data部分）
type AppOfflineAppstoresResp struct {
	Result bool `json:"result"` // 结果
}
