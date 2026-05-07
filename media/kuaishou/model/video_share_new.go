package model

import "errors"

// VideoShareNewReq 视频库推送视频（新版）请求
type VideoShareNewReq struct {
	accessTokenReq
	AdvertiserId     int64   `json:"advertiser_id"`          // 当前账户ID，必填
	PhotoIds         []int64 `json:"photo_ids"`              // 被推送视频ID列表，必填
	AccountIds       []int64 `json:"account_ids"`            // 被推送账户ID列表，必填
	ShareAccountType int     `json:"share_account_type"`     // 分享账户类型，必填：1同主体同代理商 2同主体 9同代理 10同运营自定义客户
	SyncProfile      *bool   `json:"sync_profile,omitempty"` // 是否分享个人主页，默认不同步
}

func (receiver *VideoShareNewReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoShareNewReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PhotoIds) == 0 {
		err = errors.New("photo_ids is empty")
		return
	}
	if len(receiver.AccountIds) == 0 {
		err = errors.New("account_ids is empty")
		return
	}
	if receiver.ShareAccountType <= 0 {
		err = errors.New("share_account_type is empty")
		return
	}
	return
}

// VideoShareNewResp 视频库推送视频（新版）响应数据（仅data部分）
type VideoShareNewResp struct {
	ShareStatus               int                   `json:"share_status"`                 // 推送结果：-1全部失败 0部分失败 1全部成功 2等待异步结果
	SharePhotoExists          bool                  `json:"share_photo_exists"`           // 推送视频是否已存在于被推送账户中
	NotSupportedInternalPhoto bool                  `json:"not_supported_internal_photo"` // 被推送账户是否支持内部推送：true不支持 false支持
	MismatchedAccountList     []int64               `json:"mismatched_account_list"`      // 不符合条件的账户ID列表
	NeedToRetryList           []VideoShareRetryItem `json:"need_to_retry_list"`           // 未推送成功的结果
	DuplicatedPhotoList       []VideoSharePhotoItem `json:"duplicated_photo_list"`        // 已存在于被推送账户中的视频
	ShareSuccessList          []VideoSharePhotoItem `json:"share_success_list"`           // 推送成功的数据
}

// VideoShareRetryItem 推送失败/待重试项
type VideoShareRetryItem struct {
	AccountId   int64  `json:"account_id"`   // 账户ID
	PhotoId     string `json:"photo_id"`     // 加密后的视频ID
	ShareResult string `json:"share_result"` // 推送结果
}

// VideoSharePhotoItem 推送成功/重复视频项
type VideoSharePhotoItem struct {
	AccountId       int64  `json:"account_id"`        // 账户ID
	PhotoId         string `json:"photo_id"`          // 加密后的视频ID
	OriginalPhotoId string `json:"original_photo_id"` // 推送账户中的原始视频ID
}
