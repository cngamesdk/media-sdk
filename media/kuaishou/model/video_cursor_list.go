package model

import "errors"

// VideoCursorListReq 游标查询视频信息请求
type VideoCursorListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`   // 广告主ID，必填
	Cursor       int64 `json:"cursor"`          // 游标值，必填，首次传0，后续传上一轮返回的最大cursor
	Limit        int   `json:"limit,omitempty"` // 每轮返回的数据量，默认200，最大1000
}

func (receiver *VideoCursorListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoCursorListReq) Validate() (err error) {
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

// VideoCursorListResp 游标查询视频信息响应数据（仅data部分）
type VideoCursorListResp struct {
	TotalCount int64               `json:"total_count"` // 视频总数
	Details    []VideoCursorDetail `json:"details"`     // 视频详情列表
}

// VideoCursorDetail 游标查询视频详情（在VideoListDetail基础上增加cursor字段）
type VideoCursorDetail struct {
	VideoListDetail
	Cursor int64 `json:"cursor"` // 游标值，用于下一轮查询
}
