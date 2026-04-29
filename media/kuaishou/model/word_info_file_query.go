package model

import "errors"

// WordInfoFileQueryReq 获取关键词导出文件请求
type WordInfoFileQueryReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
	FileId       int64 `json:"file_id"`       // 文件ID，必填
}

func (receiver *WordInfoFileQueryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoFileQueryReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.FileId <= 0 {
		err = errors.New("file_id is empty")
		return
	}
	return
}

// WordInfoFileQueryResp 获取关键词导出文件响应数据（仅data部分）
type WordInfoFileQueryResp struct {
	FileId         int64  `json:"file_id"`          // 文件ID
	FileStatus     int    `json:"file_status"`      // 文件状态:0-未知，1-生成中，2-生成成功，3-导出失败，4-上传失败
	FileUrl        string `json:"file_url"`         // 文件地址
	FileUpdateTime string `json:"file_update_time"` // 文件更新时间,毫秒时间戳
}
