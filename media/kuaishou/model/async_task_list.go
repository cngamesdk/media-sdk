package model

import "errors"

// AsyncTaskListReq 获取任务状态请求
type AsyncTaskListReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`       // 广告主ID，必填
	TaskIds      []int64 `json:"task_ids,omitempty"`  // 任务ID集，不超过10个
	Page         int     `json:"page,omitempty"`      // 页码，默认1
	PageSize     int     `json:"page_size,omitempty"` // 每页行数，默认20
}

func (receiver *AsyncTaskListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AsyncTaskListReq) Validate() (err error) {
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

// AsyncTaskListDetail 任务状态明细
type AsyncTaskListDetail struct {
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID
	TaskId       int64  `json:"task_id"`       // 任务ID
	TaskName     string `json:"task_name"`     // 任务名称
	CreateTime   string `json:"create_time"`   // 任务创建时间，格式 yyyy-MM-dd HH:mm:ss
	TaskStatus   int    `json:"task_status"`   // 任务状态：0=新建 1=处理中 2=处理成功 3=处理失败
	FileSize     int64  `json:"file_size"`     // 文件大小（字节数）
}

// AsyncTaskListResp 获取任务状态响应数据（仅data部分）
type AsyncTaskListResp struct {
	TotalCount int64                 `json:"total_count"` // 任务总数
	Details    []AsyncTaskListDetail `json:"details"`     // 详情
}
