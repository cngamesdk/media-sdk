package model

import "errors"

// PhotoPushToProfileReq 隐藏视频同步个人主页请求
type PhotoPushToProfileReq struct {
	accessTokenReq
	PhotoIds     string `json:"photo_ids"`     // 视频ID，英文逗号隔开，必填
	AdvertiserId int64  `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *PhotoPushToProfileReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPushToProfileReq) Validate() (err error) {
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
	return
}

// FailPhotoDetail 同步失败详情
type FailPhotoDetail struct {
	PhotoId   string `json:"photo_id"`   // 失败视频ID
	AuthorId  int64  `json:"author_id"`  // 作者ID
	PhotoName string `json:"photo_name"` // 视频名称
	Reason    string `json:"reason"`     // 同步失败具体原因
}

// PhotoPushToProfileResp 隐藏视频同步个人主页响应数据（仅data部分）
type PhotoPushToProfileResp struct {
	SuccessCount int               `json:"success_count"` // 成功个数
	FailCount    int               `json:"fail_count"`    // 失败个数
	FailDetails  []FailPhotoDetail `json:"fail_details"`  // 失败详情
}
