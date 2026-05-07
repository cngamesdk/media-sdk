package model

import "errors"

// PhotoPackageCreateReq 新建素材包请求
type PhotoPackageCreateReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID，必填
	Name         string   `json:"name"`                // 素材包名称，必填
	PhotoIds     []string `json:"photo_ids,omitempty"` // 视频ids
}

func (receiver *PhotoPackageCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Name == "" {
		err = errors.New("name is empty")
		return
	}
	return
}

// PhotoPackageCreateResp 新建素材包响应数据（仅data部分）
type PhotoPackageCreateResp struct {
	PhotoPackageId int64    `json:"photo_package_id"` // 素材包ID
	UserId         int64    `json:"user_id"`          // 快手ID
	Name           string   `json:"name"`             // 素材包名称
	PhotoIds       []string `json:"photo_ids"`        // 视频ID集合
	Status         int      `json:"status"`           // 状态，1-有效，0-删除
	CreatTime      int64    `json:"creat_time"`       // 创建时间
	UpdateTime     int64    `json:"update_time"`      // 更新时间
	PhotoAddQuota  int      `json:"photo_add_quota"`  // 可继续追加的视频数量
	AdvertiserId   int64    `json:"advertiser_id"`    // 广告主id
}
