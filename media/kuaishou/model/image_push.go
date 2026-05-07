package model

import "errors"

// ImagePushReq 图片推送请求
type ImagePushReq struct {
	accessTokenReq
	AdvertiserId       int64    `json:"advertiser_id"`          // 广告主ID，必填
	ImageTokens        []string `json:"image_tokens,omitempty"` // 分享图片的tokens，与pic_ids二选一
	ShareAdvertiserIds []int64  `json:"share_advertiser_ids"`   // 推送账户，必填
	PicIds             []string `json:"pic_ids,omitempty"`      // 分享图片的ids，与image_tokens二选一
	ShareAccountType   int      `json:"share_account_type"`     // 分享账户类型，必填：1-同主体同代理商账户，2-同主体下账户，3-内部账户
}

func (receiver *ImagePushReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImagePushReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.ShareAdvertiserIds) == 0 {
		err = errors.New("share_advertiser_ids is empty")
		return
	}
	if receiver.ShareAccountType <= 0 {
		err = errors.New("share_account_type is empty")
		return
	}
	if len(receiver.ImageTokens) == 0 && len(receiver.PicIds) == 0 {
		err = errors.New("image_tokens or pic_ids is required")
		return
	}
	if len(receiver.ImageTokens) > 0 && len(receiver.PicIds) > 0 {
		err = errors.New("image_tokens and pic_ids cannot be set at the same time")
		return
	}
	return
}

// ImagePushDetail 图片推送详情
type ImagePushDetail struct {
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID
	ImageToken   string `json:"image_token"`   // 图片token
	PicId        string `json:"pic_id"`        // 加密的照片ID
}

// ImagePushResp 图片推送响应数据（仅data部分）
type ImagePushResp struct {
	Details []ImagePushDetail `json:"details"` // 推送图片详细信息列表
}
