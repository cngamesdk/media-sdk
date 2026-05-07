package model

import "errors"

// VideoListReq 获取视频信息list请求
type VideoListReq struct {
	accessTokenReq
	AdvertiserId    int64    `json:"advertiser_id"`               // 广告主ID，必填
	PhotoIds        []string `json:"photo_ids,omitempty"`         // 视频ID列表，最多100个
	StartDate       string   `json:"start_date,omitempty"`        // 起始时间，与end_date同时传，格式yyyy-MM-dd
	EndDate         string   `json:"end_date,omitempty"`          // 结束时间，与start_date同时传，格式yyyy-MM-dd
	Page            int      `json:"page,omitempty"`              // 当前页码，默认1
	PageSize        int      `json:"page_size,omitempty"`         // 每页行数，默认20，最大200
	Signature       string   `json:"signature,omitempty"`         // 视频MD5
	PhotoName       string   `json:"photo_name,omitempty"`        // 视频名称
	OuterLoopNative int      `json:"outer_loop_native,omitempty"` // 是否获取原生视频，0=False，1=True
	PhotoUserId     int64    `json:"photo_user_id,omitempty"`     // 视频所属的UserId
	UpdateStartDate string   `json:"update_start_date,omitempty"` // 起始更新时间，与update_end_date同时传，格式yyyy-MM-dd
	UpdateEndDate   string   `json:"update_end_date,omitempty"`   // 结束更新时间，与update_start_date同时传，格式yyyy-MM-dd
}

func (receiver *VideoListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoListReq) Validate() (err error) {
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

// VideoListResp 获取视频信息list响应数据（仅data部分）
type VideoListResp struct {
	TotalCount int64             `json:"total_count"` // 视频总数
	Details    []VideoListDetail `json:"details"`     // 视频详情列表
}

// VideoListDetail 视频详情
type VideoListDetail struct {
	PhotoId               string                 `json:"photo_id"`                 // 视频ID
	PhotoName             string                 `json:"photo_name"`               // 视频名称
	CoverUrl              string                 `json:"cover_url"`                // 视频首帧图片链接
	Url                   string                 `json:"url"`                      // 视频预览链接
	Signature             string                 `json:"signature"`                // 视频MD5
	NewStatus             int                    `json:"new_status"`               // 视频状态：0删除 1可用 2转码中 3审核中 11转码失败 12审核失败 13已屏蔽 14客户端删除
	Status                int                    `json:"status"`                   // 视频状态：0正常 1删除
	Source                int                    `json:"source"`                   // 视频来源：0本地 1开眼 2素造 4mapi 7聚星 8LA上传 10个人主页 11代理商自制
	Authorization         int                    `json:"authorization"`            // 授权状态：1授权 2取消授权
	Duration              int64                  `json:"duration"`                 // 视频时长，单位毫秒
	Width                 int                    `json:"width"`                    // 视频宽度
	Height                int                    `json:"height"`                   // 视频高度
	CreateTime            string                 `json:"create_time"`              // 创建时间
	UploadTime            string                 `json:"upload_time"`              // 上传时间
	PhotoTag              []string               `json:"photo_tag"`                // 视频标签
	PhotoUserId           int64                  `json:"photo_user_id"`            // 视频所属的UserId
	OuterLoopNative       int                    `json:"outer_loop_native"`        // 是否是原生视频：0或2=非原生 1或3=原生
	AtlasPicIds           []string               `json:"atlas_pic_ids"`            // 图文视频图片ID列表
	AtlasAudioBsKey       string                 `json:"atlas_audio_bs_key"`       // 图文视频音频bs_key
	AtlasAudioUrl         string                 `json:"atlas_audio_url"`          // 图文视频音频URL
	ShieldStatus          int                    `json:"shield_status"`            // 是否已同步个人主页：1否 0是
	PhotoDupStatus        int                    `json:"photo_dup_status"`         // 素材创新度：0原创 1重复
	LowQualityStatus      int                    `json:"low_quality_status"`       // 素材低质标签：0非低质 1低质
	PhotoWakeStatus       int                    `json:"photo_wake_status"`        // 视频唤醒状态（已废弃）
	PhotoValuateInfo      *VideoValuateInfo      `json:"photo_valuate_info"`       // 视频素材评价
	PhotoTagIdentifyItems []VideoTagIdentifyItem `json:"photo_tag_identify_items"` // 素材内容标签
}

// VideoValuateInfo 视频素材评价
type VideoValuateInfo struct {
	SimLabel          string `json:"sim_label"`           // 视频重复信息（风控）
	QualityLabel      string `json:"quality_label"`       // 视频质量信息
	IsDupPhoto        bool   `json:"is_dup_photo"`        // 视频是否重复
	IsDelayReview     *bool  `json:"is_delay_review"`     // 视频是否延审
	RunningScore      int    `json:"running_score"`       // 视频跑量分：1优质 2低质
	HitTagCombination int    `json:"hit_tag_combination"` // 是否命中标签组合
}

// VideoTagIdentifyItem 素材内容标签项
type VideoTagIdentifyItem struct {
	Dimension             string               `json:"dimension"`             // 标签维度描述
	AdAssetTagActionItems []VideoTagActionItem `json:"adAssetTagActionItems"` // 标签操作列表
}

// VideoTagActionItem 标签操作项
type VideoTagActionItem struct {
	Action                 string             `json:"action"`                 // 操作类型
	AdAssetTagInfoItemList []VideoTagInfoItem `json:"adAssetTagInfoItemList"` // 标签信息列表
}

// VideoTagInfoItem 标签信息项
type VideoTagInfoItem struct {
	TagId     int64  `json:"tagId"`     // 标签ID
	TagName   string `json:"tagName"`   // 标签名称
	TagDetail string `json:"tagDetail"` // 标签详情
}
