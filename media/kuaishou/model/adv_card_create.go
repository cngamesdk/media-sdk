package model

import "errors"

// AdvCardEmojiItem emoji信息
type AdvCardEmojiItem struct {
	EmojiCode string `json:"emoji_code"` // emoji编码（如"[加油]"）
	EmojiUrl  string `json:"emoji_url"`  // emoji图片URL
}

// AdvCardCreateItem 创建高级创意-单条卡片
type AdvCardCreateItem struct {
	CardType    int                `json:"card_type"`              // 卡片类型，必填：100=图片卡片 101=多利益卡-图文 102=多利益卡-多标签 103=电商促销样式 104=快捷评论卡 107=倒计时卡 108=优惠券卡 125=价格差卡 131=商品卡-汽车 132=商品卡-电商 133=商品卡-小说 134=商品卡-房产 200=推广位
	Url         string             `json:"url,omitempty"`          // 图片URL，card_type=100/101/103/125时必填
	Width       int                `json:"width,omitempty"`        // 图片宽度
	Height      int                `json:"height,omitempty"`       // 图片高度
	Title       string             `json:"title,omitempty"`        // 标题，card_type=101/102/103/107/108/125/200时必填，104时非必填
	SubTitle    string             `json:"sub_title,omitempty"`    // 副标题，card_type=102时必填
	Price       int                `json:"price,omitempty"`        // 原价格（单位：分），card_type=103/125时必填
	SalePrice   int                `json:"sale_price,omitempty"`   // 售卖价（单位：分），card_type=103/125时必填
	ContentType int                `json:"content_type,omitempty"` // 卡片内容类型，card_type=104时必填；emoji快捷评论卡填2
	EmojiList   []AdvCardEmojiItem `json:"emoji_list,omitempty"`   // emoji信息列表，card_type=104时必填
	Discount    int                `json:"discount,omitempty"`     // 折扣金额，card_type=108时必填
	StyleType   int                `json:"style_type,omitempty"`   // 高级创意样式类型：1=普通卡片 2=推广位
	BeginTime   int64              `json:"begin_time,omitempty"`   // 倒计时卡开始时间，card_type=107时必填
	EndTime     int64              `json:"end_time,omitempty"`     // 倒计时卡结束时间，card_type=107时必填
	LibraryId   int64              `json:"library_id,omitempty"`   // 商品库ID，card_type=131/132/133/134时必填
	ProductId   int64              `json:"product_id,omitempty"`   // 商品ID，card_type=131/132/133/134时必填
}

// AdvCardCreateReq 创建高级创意请求
type AdvCardCreateReq struct {
	accessTokenReq
	AdvertiserId int64               `json:"advertiser_id"` // 广告主ID，必填
	AdvList      []AdvCardCreateItem `json:"adv_list"`      // 卡片列表，必填
}

func (receiver *AdvCardCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvCardCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.AdvList) == 0 {
		err = errors.New("adv_list is empty")
		return
	}
	return
}

// AdvCardCreateResp 创建高级创意响应数据（仅data部分）
type AdvCardCreateResp struct {
	AdvList []int64 `json:"adv_list"` // 卡片ID数组
}
