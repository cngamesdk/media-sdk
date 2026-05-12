package model

import "errors"

// AppServiceCategoryReq 获取APP服务类目详情请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/service/category
type AppServiceCategoryReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *AppServiceCategoryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppServiceCategoryReq) Validate() (err error) {
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

// AppServiceCategoryItem APP服务类目条目
type AppServiceCategoryItem struct {
	Id       int64                    `json:"id"`       // APP类目ID
	Name     string                   `json:"name"`     // APP类目名称
	Children []AppServiceCategoryItem `json:"children"` // 子节点
}

// AppServiceCategoryResp 获取APP服务类目详情响应数据（仅data部分）
type AppServiceCategoryResp []AppServiceCategoryItem
