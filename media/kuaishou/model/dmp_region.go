package model

import "errors"

// DmpRegionReq 人群包数据请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/dmp/region
type DmpRegionReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *DmpRegionReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DmpRegionReq) Validate() (err error) {
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

// DmpRegionItem 人群包地区条目
type DmpRegionItem struct {
	Children []string `json:"children"` // 下一级地区
	Id       int64    `json:"id"`       // 地区编号
	Name     string   `json:"name"`     // 地区名称
}

// DmpRegionResp 人群包数据响应数据（仅data部分）
type DmpRegionResp []DmpRegionItem
