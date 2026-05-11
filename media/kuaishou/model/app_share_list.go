package model

import "errors"

// AppShareListReq 获取应用已共享账号列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/share/list
// 注意：应用共享类型须为 share_type=1（账号共享）才会有数据
type AppShareListReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"`       // 广告主ID，必填
	AppId        int64  `json:"app_id"`              // 应用ID，必填
	KeyWord      string `json:"key_word,omitempty"`  // 搜索关键词，账号关键词
	Page         int    `json:"page,omitempty"`      // 当前页码，默认1
	PageSize     int    `json:"page_size,omitempty"` // 分页大小，默认20
}

func (receiver *AppShareListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareListReq) Validate() (err error) {
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

// AppShareListItem 已共享账号列表条目
type AppShareListItem struct {
	AccountId   int64  `json:"account_id"`   // 账号ID
	AccountName string `json:"account_name"` // 账号名称
}

// AppShareListResp 获取应用已共享账号列表响应数据（仅data部分）
type AppShareListResp struct {
	CurrentPage int                `json:"current_page"` // 当前页码
	PageSize    int                `json:"page_size"`    // 分页大小
	TotalCount  int64              `json:"total_count"`  // 总数量
	List        []AppShareListItem `json:"list"`         // 已共享账号列表
}
