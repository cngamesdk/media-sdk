package model

import "errors"

// AsyncTaskDownloadReq 数据下载请求（GET请求，返回CSV文件）
type AsyncTaskDownloadReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	TaskId       int64 `json:"task_id"`       // 任务ID，必填。只能获取最近6个月内创建的任务，超过6个月需重新创建
}

func (receiver *AsyncTaskDownloadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AsyncTaskDownloadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.TaskId <= 0 {
		err = errors.New("task_id is empty")
		return
	}
	return
}

// AsyncTaskDownloadResp 数据下载响应数据（仅data部分，CSV文件字节内容，字段同实时报表）
type AsyncTaskDownloadResp struct {
	FileData []byte `json:"-"` // CSV文件内容
}
