package model

import "errors"

// VideoGetReq 获取视频信息get请求
type VideoGetReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID，必填
	PhotoIds     []string `json:"photo_ids"`     // 视频ID集，必填
}

func (receiver *VideoGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoGetReq) Validate() (err error) {
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

// VideoGetResp 获取视频信息get响应数据（仅data部分，为数组）
type VideoGetResp []VideoGetDetail

// VideoGetDetail 视频详情（get接口）
type VideoGetDetail struct {
	PhotoId               string                 `json:"photo_id"`              // 视频ID
	PhotoName             string                 `json:"photo_name"`            // 视频名称
	CoverUrl              string                 `json:"cover_url"`             // 视频首帧图片链接
	Url                   string                 `json:"url"`                   // 视频预览链接
	Signature             string                 `json:"signature"`             // 视频MD5
	NewStatus             int                    `json:"new_status"`            // 视频状态：0删除 1可用 2转码中 3审核中 11转码失败 12审核失败 13已屏蔽 14客户端删除
	Status                int                    `json:"status"`                // 视频状态：0正常 1删除
	Source                int                    `json:"source"`                // 视频来源：0本地 1开眼 2素造 4mapi 7聚星 8LA上传 10个人主页 11代理商自制
	Authorization         int                    `json:"authorization"`         // 授权状态：1授权 2取消授权
	Duration              int64                  `json:"duration"`              // 视频时长，单位毫秒
	Width                 int                    `json:"width"`                 // 视频宽度
	Height                int                    `json:"height"`                // 视频高度
	CreateTime            string                 `json:"create_time"`           // 创建时间
	UploadTime            string                 `json:"upload_time"`           // 上传时间
	PhotoTag              []string               `json:"photo_tag"`             // 视频标签
	PhotoUserId           int64                  `json:"photoUserId"`           // 视频所属的UserId
	OuterLoopNative       int                    `json:"outer_loop_native"`     // 是否是原生视频
	AtlasPicIds           []string               `json:"atlas_pic_ids"`         // 图集图片ID列表
	AtlasAudioBsKey       string                 `json:"atlas_audio_bs_key"`    // 图集音频bs_key
	AtlasAudioUrl         string                 `json:"atlas_audio_url"`       // 音频播放链接
	ShieldStatus          int                    `json:"shieldStatus"`          // 是否在个人主页隐藏：0非隐藏 1隐藏
	PhotoDupStatus        int                    `json:"photo_dup_status"`      // 素材创新度：0原创 1重复
	LowQualityStatus      int                    `json:"low_quality_status"`    // 素材低质标签：0非低质 1低质
	PhotoWakeStatus       int                    `json:"photo_wake_status"`     // 视频唤醒状态
	AdPhotoValuateInfo    *VideoGetValuateInfo   `json:"adPhotoValuateInfo"`    // 视频素材评价
	PhotoTagIdentifyItems []VideoTagIdentifyItem `json:"photoTagIdentifyItems"` // 素材内容标签
}

// VideoGetValuateInfo 视频素材评价（get接口，camelCase字段）
type VideoGetValuateInfo struct {
	SimLabel                string `json:"simLabel"`                // 视频重复信息（风控）
	QualityLabel            string `json:"qualityLabel"`            // 视频质量信息
	QuotaMsg                string `json:"quotaMsg"`                // 视频内卷份额信息
	IsDupPhoto              bool   `json:"isDupPhoto"`              // 视频是否重复
	IsDelayReview           *bool  `json:"isDelayReview"`           // 视频是否延审，可为null
	RunningScore            int    `json:"runningScore"`            // 视频跑量分：1优质 2低质
	HitTagCombination       int    `json:"hitTagCombination"`       // 是否命中标签组合
	OptimizationSuggestions string `json:"optimizationSuggestions"` // 优化建议
}
