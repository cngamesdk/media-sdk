package model

import "errors"

// ChannelColumn 渠道分包信息
type ChannelColumn struct {
	ChannelName string `json:"channel_name"` // 渠道名
	Description string `json:"description"`  // 备注
}

// SubpkgAddReq 新建应用分包请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/add
type SubpkgAddReq struct {
	accessTokenReq
	AdvertiserId    int64           `json:"advertiser_id"`             // 广告主ID，必填
	ParentPackageId int64           `json:"parent_package_id"`         // 母包ID，必填
	Type            int             `json:"type"`                      // 分包方式：1-系统自动分包，2-上传渠道号列表，必填
	ChannelId       []string        `json:"channel_id,omitempty"`      // 上传的渠道号列表，当type=2时填写，单次最多100个
	Count           int             `json:"count,omitempty"`           // 分包数量，当type=1时填写，单次最多100
	ChannelColumns  []ChannelColumn `json:"channel_columns,omitempty"` // 渠道分包，与channel_id功能类似但不可同时传递
}

func (receiver *SubpkgAddReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SubpkgAddReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ParentPackageId <= 0 {
		err = errors.New("parent_package_id is empty")
		return
	}
	if receiver.Type != 1 && receiver.Type != 2 {
		err = errors.New("type must be 1 or 2")
		return
	}
	return
}

// SubpkgAddRespItem 新建应用分包响应条目
type SubpkgAddRespItem struct {
	BuildStatus     int    `json:"build_status"`      // 构建状态：0-创建中，1-构建中，2-构建成功，3-构建失败
	ChannelId       string `json:"channel_id"`        // 渠道号(分包号)
	PackageId       int64  `json:"package_id"`        // 分包ID
	ParentPackageId int64  `json:"parent_package_id"` // 绑定的母包ID
	Description     string `json:"description"`       // 分包备注
}

// SubpkgAddResp 新建应用分包响应数据（仅data部分）
type SubpkgAddResp []SubpkgAddRespItem
