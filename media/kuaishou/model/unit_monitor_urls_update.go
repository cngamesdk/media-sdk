package model

import "errors"

// UnitMonitorUrlUpdateItem 批量更新监测链接-单条请求项
type UnitMonitorUrlUpdateItem struct {
	UnitId              int64  `json:"unit_id"`                           // 广告组ID，必填
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // actionbar点击监测
	ClickUrl            string `json:"click_url,omitempty"`               // 点击监测（排除粉丝直播推广）
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 曝光监测
	LiveTrackUrl        string `json:"live_track_url,omitempty"`          // 点击监测（计划为粉丝直播推广）
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 3s播放监测链接
}

// UnitMonitorUrlsUpdateReq 批量更新监测链接请求
type UnitMonitorUrlsUpdateReq struct {
	accessTokenReq
	AdvertiserId    int64                      `json:"advertiser_id"`     // 广告主ID，必填
	UnitMonitorUrls []UnitMonitorUrlUpdateItem `json:"unit_monitor_urls"` // 监测链接详情，必填
}

func (receiver *UnitMonitorUrlsUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitMonitorUrlsUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.UnitMonitorUrls) == 0 {
		err = errors.New("unit_monitor_urls is empty")
		return
	}
	return
}

// UnitMonitorUrlUpdateResultItem 批量更新监测链接-单条返回项
type UnitMonitorUrlUpdateResultItem struct {
	UnitId              int64  `json:"unit_id"`                 // 广告组ID
	ActionbarClickUrl   string `json:"actionbar_click_url"`     // actionbar点击监测
	ClickUrl            string `json:"click_url"`               // 点击监测（排除粉丝直播推广）
	ImpressionUrl       string `json:"impression_url"`          // 曝光监测
	LiveTrackUrl        string `json:"live_track_url"`          // 点击监测（计划为粉丝直播推广）
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url"` // 3s播放监测链接
	Result              bool   `json:"result"`                  // 修改是否成功
	Message             string `json:"message"`                 // 修改成功与否提示信息
}

// UnitMonitorUrlsUpdateResp 批量更新监测链接响应数据（仅data部分）
type UnitMonitorUrlsUpdateResp struct {
	UnitMonitorUrls []UnitMonitorUrlUpdateResultItem `json:"unit_monitor_urls"` // 监测链接更新结果列表
}
