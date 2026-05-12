package model

import "errors"

// AppShareCanShareCorpListReq 获取可共享的主体列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/canShare/corpList
// 注意：此接口只有在跨主体共享白名单内才会有数据返回
type AppShareCanShareCorpListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *AppShareCanShareCorpListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareCanShareCorpListReq) Validate() (err error) {
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

// AppShareCanShareCorpItem 可共享主体条目
type AppShareCanShareCorpItem struct {
	CorpId   int64  `json:"corp_id"`   // 主体ID
	CorpName string `json:"corp_name"` // 主体名称
}

// AppShareCanShareCorpListResp 获取可共享的主体列表响应数据（仅data部分）
type AppShareCanShareCorpListResp struct {
	List []AppShareCanShareCorpItem `json:"list"` // 可共享主体列表
}
