package model

import "errors"

// AdvancedCreativeReviewReq 获取程序化创意/智能创意审核信息请求
type AdvancedCreativeReviewReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	UnitIds      []int64 `json:"unit_ids"`      // 广告组ID集，必填，数量不超过20个
}

func (receiver *AdvancedCreativeReviewReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvancedCreativeReviewReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.UnitIds) == 0 {
		err = errors.New("unit_ids is empty")
		return
	}
	return
}

// AdvancedCreativeReviewReasonAndModify 审核拒绝/限流原因和建议
type AdvancedCreativeReviewReasonAndModify struct {
	Reason string   `json:"reason"` // 审核拒绝原因
	Modify []string `json:"modify"` // 修改建议
}

// AdvancedCreativeReviewElementReason 审核元素拒绝/限流详情
type AdvancedCreativeReviewElementReason struct {
	Id              int                                     `json:"id"`                // 视频ID或者创意ID
	Type            int                                     `json:"type"`              // 审核元素：1=视频 2=封面 3=广告语 4=图片 6=创意 7=图集
	NegativeType    int                                     `json:"negative_type"`     // 限流类型：1=低质 2=降级 3=负向 4=封面没过
	ReasonAndModify []AdvancedCreativeReviewReasonAndModify `json:"reason_and_modify"` // 审核拒绝/限流原因和建议
}

// AdvancedCreativeCommunityReview 社区审核详情
type AdvancedCreativeCommunityReview struct {
	Id                    int64  `json:"id"`                      // 创意ID
	PhotoId               string `json:"photo_id"`                // 视频ID
	CoverId               int64  `json:"cover_id"`                // 封面ID
	CoverUrl              string `json:"cover_url"`               // 封面url
	Caption               string `json:"caption"`                 // 视频标题
	CommunityReviewDetail string `json:"community_review_detail"` // 审核拒绝理由
}

// AdvancedCreativeAuditDetail 审核限流详情
type AdvancedCreativeAuditDetail struct {
	CommunityReviewStatus               int                                   `json:"community_review_status"`                 // 社区审核状态：0=未送审 1=待审核 2=审核通过 3=审核拒绝 5=部分审核通过
	CommunityReviewDetail               string                                `json:"community_review_detail"`                 // 社区审核拒绝原因
	ReviewStatus                        int                                   `json:"review_status"`                           // 商业审核状态：0=未送审 1=待审核 2=审核通过 3=审核拒绝 5=部分审核通过
	ReviewDetail                        string                                `json:"review_detail"`                           // 商业审核拒绝理由
	ReviewReason                        []AdvancedCreativeReviewElementReason `json:"review_reason"`                           // 商业审核拒绝详情
	LimitingReason                      []AdvancedCreativeReviewElementReason `json:"limiting_reason"`                         // 限流详情
	AdvCreativeCommunityReviewDetail    []AdvancedCreativeCommunityReview     `json:"adv_creative_community_review_detail"`    // 程序化创意社区审核详情
	CustomCreativeCommunityReviewDetail *AdvancedCreativeCommunityReview      `json:"custom_creative_community_review_detail"` // 自定义创意社区审核详情
}

// AdvancedCreativeReviewDetail 审核信息
type AdvancedCreativeReviewDetail struct {
	Reason string `json:"reason"` // 审核拒绝原因描述
}

// AdvancedCreativeCombineDetail 审核不通过和正在审核的创意组合
type AdvancedCreativeCombineDetail struct {
	Id            int64                         `json:"id"`              // 创意ID
	PhotoId       string                        `json:"photo_id"`        // 视频ID（已加密）
	CoverUrl      string                        `json:"cover_url"`       // 封面URL
	Caption       string                        `json:"caption"`         // 作品广告语
	ReviewStatus  int                           `json:"review_status"`   // 审核状态：1=审核中 2=审核通过 3=审核不通过
	ReviewDetail  *AdvancedCreativeReviewDetail `json:"review_detail"`   // 审核信息
	PutStatus     int                           `json:"put_status"`      // 操作状态：1=投放中 2=暂停 3=删除
	PicImageToken string                        `json:"pic_image_token"` // 联盟图片imageToken（横版/竖版）
	PicUrl        string                        `json:"pic_url"`         // 联盟图片url（横版/竖版）
	AppGradeType  int                           `json:"app_grade_type"`  // 审核分级：0=默认 1=降级（限制部分流量）
	AuditDetail   *AdvancedCreativeAuditDetail  `json:"audit_detail"`    // 审核限流详情
}

// AdvancedCreativeReviewDetailItem 程序化创意审核详情项
type AdvancedCreativeReviewDetailItem struct {
	UnitId             int64                           `json:"unit_id"`              // 广告组ID
	Slogans            []string                        `json:"slogans"`              // 审核不通过的封面广告语
	CombineDetailViews []AdvancedCreativeCombineDetail `json:"combine_detail_views"` // 审核不通过和正在审核的创意组合
}

// AdvancedCreativeReviewResp 获取程序化创意/智能创意审核信息响应数据（仅data部分）
type AdvancedCreativeReviewResp struct {
	Details []AdvancedCreativeReviewDetailItem `json:"details"` // 详情列表
}
