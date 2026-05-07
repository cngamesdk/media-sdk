package model

import "errors"

// PhotoPackageQueryReq 查询素材包请求
type PhotoPackageQueryReq struct {
	accessTokenReq
	AdvertiserId   int64  `json:"advertiser_id"`              // 广告主ID，必填
	PhotoPackageId int64  `json:"photo_package_id,omitempty"` // 素材包id
	NameLike       string `json:"name_like,omitempty"`        // 素材包名称（模糊搜索）
	Page           int    `json:"page"`                       // 当前页
	PageSize       int    `json:"page_size"`                  // 分页大小
}

func (receiver *PhotoPackageQueryReq) Format() {
	receiver.accessTokenReq.Format()
	if receiver.Page <= 0 {
		receiver.Page = 1
	}
	if receiver.PageSize <= 0 {
		receiver.PageSize = 20
	}
}

func (receiver *PhotoPackageQueryReq) Validate() (err error) {
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

// PhotoPackageInfo 素材包详情信息
type PhotoPackageInfo struct {
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

// PhotoPackageQueryResp 查询素材包响应数据（仅data部分）
type PhotoPackageQueryResp struct {
	TotalCount int64              `json:"total_count"` // 总数
	Details    []PhotoPackageInfo `json:"details"`     // 详情列表
}
