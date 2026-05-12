package model

import "errors"

// AppShareCanShareAccountListReq 获取可共享的账号列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/canShare/accountList
// 备注：此接口已提供分页能力，分页参数需要进行填写
type AppShareCanShareAccountListReq struct {
	accessTokenReq
	AdvertiserId    int64   `json:"advertiser_id"`               // 广告主ID，必填
	SearchAccountId []int64 `json:"search_account_id,omitempty"` // 精确查找的账号，可选，上限500个
	CurrentPage     int     `json:"current_page"`                // 当前页，必填，默认填1
	PageSize        int     `json:"page_size"`                   // 分页大小，必填，单页上限500
}

func (receiver *AppShareCanShareAccountListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareCanShareAccountListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CurrentPage <= 0 {
		err = errors.New("current_page must be greater than 0")
		return
	}
	if receiver.PageSize <= 0 {
		err = errors.New("page_size must be greater than 0")
		return
	}
	if len(receiver.SearchAccountId) > 500 {
		err = errors.New("search_account_id must not exceed 500 items")
		return
	}
	return
}

// AppShareCanShareAccountItem 可共享账号条目
type AppShareCanShareAccountItem struct {
	AccountId   int64  `json:"account_id"`   // 账号ID
	AccountName string `json:"account_name"` // 账号名称
}

// AppShareCanShareAccountListResp 获取可共享的账号列表响应数据（仅data部分）
type AppShareCanShareAccountListResp struct {
	CurrentPage int                           `json:"current_page"` // 当前页
	PageSize    int                           `json:"page_size"`    // 页码大小
	TotalCount  int64                         `json:"total_count"`  // 总数
	List        []AppShareCanShareAccountItem `json:"list"`         // 可共享账号列表
}
