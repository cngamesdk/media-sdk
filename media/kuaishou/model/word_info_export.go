package model

import "errors"

// WordInfoExportReq 批量导出关键词请求
type WordInfoExportReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
	CampaignId   int64 `json:"campaign_id"`   // 计划ID
	UnitId       int64 `json:"unit_id"`       // 单元ID
	StartTime    int64 `json:"start_time"`    // 起始时间，必填；关键词创建时间(毫秒时间戳)，起始～截止时间跨度不超过180天
	EndTime      int64 `json:"end_time"`      // 截止时间，必填；关键词创建时间(毫秒时间戳)，起始～截止时间跨度不超过180天
}

func (receiver *WordInfoExportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoExportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.StartTime <= 0 {
		err = errors.New("start_time is empty")
		return
	}
	if receiver.EndTime <= 0 {
		err = errors.New("end_time is empty")
		return
	}
	if receiver.EndTime <= receiver.StartTime {
		err = errors.New("end_time must be greater than start_time")
		return
	}
	return
}

// WordInfoExportResp 批量导出关键词响应数据（仅data部分）
type WordInfoExportResp struct {
	FileId int64 `json:"file_id"` // 文件ID
}
