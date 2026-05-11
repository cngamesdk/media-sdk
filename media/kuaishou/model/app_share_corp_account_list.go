package model

import "errors"

// AppShareCorpAccountPageInfo 分页信息
type AppShareCorpAccountPageInfo struct {
	CurrentPage int `json:"current_page"` // 当前页码，必填
	PageSize    int `json:"page_size"`    // 分页大小，必填
}

// AppShareCorpAccountListReq 获取单个主体下共享账号列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/share/listCorpAccount
// 注意：应用共享类型须为 share_type=2（主体共享）才会有数据
type AppShareCorpAccountListReq struct {
	accessTokenReq
	AdvertiserId int64                       `json:"advertiser_id"` // 广告主ID，必填
	AppId        int64                       `json:"app_id"`        // 应用ID，必填
	CorpId       int64                       `json:"corp_id"`       // 主体ID，必填
	PageInfo     AppShareCorpAccountPageInfo `json:"page_info"`     // 分页信息，必填
}

func (receiver *AppShareCorpAccountListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppShareCorpAccountListReq) Validate() (err error) {
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
	if receiver.CorpId <= 0 {
		err = errors.New("corp_id is empty")
		return
	}
	if receiver.PageInfo.CurrentPage <= 0 {
		err = errors.New("page_info.current_page must be greater than 0")
		return
	}
	if receiver.PageInfo.PageSize <= 0 {
		err = errors.New("page_info.page_size must be greater than 0")
		return
	}
	return
}

// AppShareCorpAccountItem 主体下共享账号
type AppShareCorpAccountItem struct {
	AccountId   int64  `json:"account_id"`   // 账号ID
	AccountName string `json:"account_name"` // 账号名称
}

// AppShareCorpAccountListItem 主体列表条目
type AppShareCorpAccountListItem struct {
	CorpId          int64                     `json:"corp_id"`           // 主体ID
	CorpName        string                    `json:"corp_name"`         // 主体名称
	TotalAccountCnt int                       `json:"total_account_cnt"` // 主体共享挂载的账号数量
	ShareAccountVos []AppShareCorpAccountItem `json:"share_account_vos"` // 账号列表
}

// AppShareCorpAccountListResp 获取单个主体下共享账号列表响应数据（仅data部分）
type AppShareCorpAccountListResp struct {
	List []AppShareCorpAccountListItem `json:"list"` // 主体及其账号列表
}
