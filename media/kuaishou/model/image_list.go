package model

import "errors"

// ImageListReq 查询图片list请求
type ImageListReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`          // 广告主ID，必填
	StartDate    string   `json:"start_date,omitempty"`   // 开始时间，格式yyyy-MM-dd，与end_date同时传或同时不传
	EndDate      string   `json:"end_date,omitempty"`     // 结束时间，格式yyyy-MM-dd，与start_date同时传或同时不传
	PicTypes     []int    `json:"pic_types,omitempty"`    // 图片类型：0-默认，5-竖版，6-横版，12-开屏，不填获取所有类型
	Page         int      `json:"page"`                   // 请求的页码数，必填，默认1
	PageSize     int      `json:"page_size"`              // 每页行数，必填，默认20，最高500
	ImageToken   string   `json:"image_token,omitempty"`  // 图片token
	Signature    string   `json:"signature,omitempty"`    // 图片MD5
	ImageTokens  []string `json:"image_tokens,omitempty"` // 图片token数组，最多100个
}

func (receiver *ImageListReq) Format() {
	receiver.accessTokenReq.Format()
	if receiver.Page <= 0 {
		receiver.Page = 1
	}
	if receiver.PageSize <= 0 {
		receiver.PageSize = 20
	}
}

func (receiver *ImageListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if (receiver.StartDate == "") != (receiver.EndDate == "") {
		err = errors.New("start_date and end_date must be set together")
		return
	}
	return
}

// ImageListDetail 图片列表详情
type ImageListDetail struct {
	Size       int64    `json:"size"`        // 图片大小
	Signature  string   `json:"signature"`   // 图片MD5
	ImageToken string   `json:"image_token"` // 图片token，创建创意时使用
	Format     string   `json:"format"`      // 图片格式
	Width      int64    `json:"width"`       // 图片宽度
	PicType    int      `json:"pic_type"`    // 图片类型：0-默认，5-竖版，6-横版，12-开屏，16-图集
	Url        string   `json:"url"`         // 图片预览地址
	Height     int64    `json:"height"`      // 图片高度
	PicId      string   `json:"pic_id"`      // 图片库图片ID（已加密）
	Name       string   `json:"name"`        // 图片名称
	PicTag     []string `json:"pic_tag"`     // 图片标签
	CreateTime int64    `json:"create_time"` // 创建时间（毫秒时间戳）
}

// ImageListResp 查询图片list响应数据（仅data部分）
type ImageListResp struct {
	TotalCount int64             `json:"total_count"` // 图片总数
	Details    []ImageListDetail `json:"details"`     // 图片详情列表
}
