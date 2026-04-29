package model

import "errors"

// CreativeReviewDetailReq 获取创意审核详情请求
type CreativeReviewDetailReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	Ids          []int64 `json:"ids"`           // 创意ID集（自定义创意）或广告组ID集（程序化创意），必填
	CreativeMold int     `json:"creative_mold"` // 创意类型，必填：1=自定义创意 2=程序化创意
}

func (receiver *CreativeReviewDetailReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeReviewDetailReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.Ids) == 0 {
		err = errors.New("ids is empty")
		return
	}
	if receiver.CreativeMold <= 0 {
		err = errors.New("creative_mold is empty")
		return
	}
	return
}

// CreativeReviewReasonAndModify 审核拒绝/限流原因和建议
type CreativeReviewReasonAndModify struct {
	Reason string   `json:"reason"` // 审核拒绝原因
	Modify []string `json:"modify"` // 修改建议
}

// CreativeReviewElementReason 审核元素拒绝/限流详情
type CreativeReviewElementReason struct {
	Id              int                             `json:"id"`                // 视频ID或者创意ID
	Type            int                             `json:"type"`              // 审核元素：1=视频 2=封面 3=广告语 4=图片 6=创意 7=图集
	NegativeType    int                             `json:"negative_type"`     // 限流类型：1=低质 2=降级 3=负向 4=封面没过
	ReasonAndModify []CreativeReviewReasonAndModify `json:"reason_and_modify"` // 审核拒绝/限流原因和建议
}

// CreativeReviewCommunityDetail 社区审核拒绝详情
type CreativeReviewCommunityDetail struct {
	Id                    int64  `json:"id"`                      // 创意ID
	PhotoId               int64  `json:"photo_id"`                // 视频ID
	CoverId               int64  `json:"cover_id"`                // 封面ID
	CoverUrl              string `json:"cover_url"`               // 封面url
	Caption               string `json:"caption"`                 // 视频标题
	CommunityReviewDetail string `json:"community_review_detail"` // 审核拒绝理由
}

// CreativeReviewDetailItem 创意审核详情项
type CreativeReviewDetailItem struct {
	Id                                  int64                           `json:"id"`                                      // 程序化创意组ID（程序化创意）或创意ID（自定义创意）
	CommunityReviewStatus               int                             `json:"community_review_status"`                 // 社区审核状态：1=审核中 2=审核通过 3=审核拒绝 5=基本通过审核
	CommunityReviewDetail               string                          `json:"community_review_detail"`                 // 社区审核拒绝原因
	ReviewStatus                        int                             `json:"review_status"`                           // 商业审核状态：1=审核中 2=审核通过 3=审核拒绝 5=基本通过审核
	ReviewDetail                        string                          `json:"review_detail"`                           // 商业审核拒绝理由
	ReviewReason                        []CreativeReviewElementReason   `json:"review_reason"`                           // 商业审核拒绝详情
	LimitingReason                      []CreativeReviewElementReason   `json:"limiting_reason"`                         // 限流详情
	AdvCreativeCommunityReviewDetail    []CreativeReviewCommunityDetail `json:"adv_creative_community_review_detail"`    // 程序化创意社区审核拒绝详情
	CustomCreativeCommunityReviewDetail *CreativeReviewCommunityDetail  `json:"custom_creative_community_review_detail"` // 自定义创意社区审核拒绝详情
}

// CreativeReviewDetailResp 获取创意审核详情响应数据（仅data部分）
type CreativeReviewDetailResp []CreativeReviewDetailItem
