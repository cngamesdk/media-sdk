package model

import "errors"

// AppShareCorpListReq 获取应用已共享主体列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/shareCorpList
// 注意：应用共享类型须为 share_type=2（主体共享）才会有数据
type AppShareCorpListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	AppId        int64 `json:"app_id"`        // 应用ID，必填
}

func (receiver *AppShareCorpListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareCorpListReq) Validate() (err error) {
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
	return
}

// AppShareCorpListItem 已共享主体列表条目
type AppShareCorpListItem struct {
	CorpId          int64  `json:"corp_id"`           // 主体ID
	CorpName        string `json:"corp_name"`         // 主体名称
	TotalAccountCnt int    `json:"total_account_cnt"` // 主体共享挂载的账号数量
}

// AppShareCorpListResp 获取应用已共享主体列表响应数据（仅data部分）
type AppShareCorpListResp struct {
	List []AppShareCorpListItem `json:"list"` // 已共享主体列表
}
