package model

import "errors"

// UnitMonitorUrlsGetReq 批量获取监测链接请求
type UnitMonitorUrlsGetReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	UnitIds      []int64 `json:"unit_ids"`      // 广告组ID列表，必填
}

func (receiver *UnitMonitorUrlsGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitMonitorUrlsGetReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.UnitIds) == 0 {
		err = errors.New("unit_ids is empty")
		return
	}
	return
}

// UnitMonitorUrlItem 广告组监测链接详情
type UnitMonitorUrlItem struct {
	UnitId              int64  `json:"unit_id"`                 // 广告组ID
	IsDpa               bool   `json:"is_dpa"`                  // 是否是DPA（DPA的监测链接在广告组上）
	ExistValidCreative  bool   `json:"exist_valid_creative"`    // 是否存在有效创意
	ActionbarClickUrl   string `json:"actionbar_click_url"`     // actionBar点击监测
	ClickUrl            string `json:"click_url"`               // 点击监测（排除粉丝直播推广）
	ImpressionUrl       string `json:"impression_url"`          // 曝光监测
	LiveTrackUrl        string `json:"live_track_url"`          // 粉丝直播推广点击监测
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url"` // 3s曝光监测链接
}

// UnitMonitorUrlsGetResp 批量获取监测链接响应数据（仅data部分）
type UnitMonitorUrlsGetResp struct {
	UnitMonitorUrls []UnitMonitorUrlItem `json:"unit_monitor_urls"` // 监测链接详情列表
}
