package model

type commonPromotionStruct struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"`
}

type PromotionCreateReq struct {
	accessTokenReq
	ProjectId int64  `json:"project_id,omitempty"` // 项目ID
	Name      string `json:"name,omitempty"`       // 单元名称，长度是1-50个字（两个英文字符占1个字）。名称不可重复，否则会报错
	Operation string `json:"operation,omitempty"`  // 单元状态， 允许值: ENABLE开启(默认值）、DISABLE关闭

	DpaMaterials

	LiveAndPromotionMaterials
}

// 常量定义
const (
	// 素材类型
	ImageModeVideo         = "CREATIVE_IMAGE_MODE_VIDEO"          // 横版视频
	ImageModeVideoVertical = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL" // 竖版视频

	// 商品库视频生成类型
	VideoTemplateTypeSmart  = "DPA_VIDEO_TEMPLATE_SMART"  // 优选商品库视频
	VideoTemplateTypeCustom = "DPA_VIDEO_TEMPLATE_CUSTOM" // 自定义商品库视频

	// 视频模板ID长度限制
	MaxVideoTaskIDLength = 1
)

// DpaMaterials DPA素材
type DpaMaterials struct {
	PromotionMaterials struct {
		VideoMaterialList []struct {
			VideoID           string   `json:"video_id,omitempty"`            // 视频ID（条件必填）
			VideoCoverID      string   `json:"video_cover_id,omitempty"`      // 视频封面图片ID（条件必填）
			ImageMode         string   `json:"image_mode"`                    // 素材类型（条件必填）
			VideoTemplateType string   `json:"video_template_type,omitempty"` // 商品库视频生成类型
			VideoTaskIDs      []string `json:"video_task_ids,omitempty"`      // 自定义商品库视频模板ID
		} `json:"video_material_list"`
		ImageMaterialList []struct {
			ImageMode string `json:"image_mode,omitempty"` // 素材类型
			Images    []struct {
				ImageId          string `json:"image_id,omitempty"`    // 图片ID
				TemplateId       int64  `json:"template_id,omitempty"` // 图片素材类型-DPA模板ID
				TemplateDataList []struct {
					BackgroundImageId string `json:"background_image_id,omitempty"` // 自定义背景图片ID
				} `json:"template_data_list,omitempty"` // 图片素材类型-DPA模板ID
			} `json:"images"` // 图片ID数组，目前仅支持传入1个
		} `json:"image_material_list,omitempty"` // 创意图片素材，上限10个，video_material_list 与 image_material_list 最少传入一种  直播链路不支持图片素材
	} `json:"promotion_materials,omitempty"` // 素材组合
}

// 直播素材与营销素材组合
type LiveAndPromotionMaterials struct {
	MaterialsType           string `json:"materials_type,omitempty"` // 素材类型，直播场景必填
	PromotionRelatedProduct []struct {
		UniqueProductId   int64 `json:"unique_product_id,omitempty"` // 商品ID
		VideoMaterialList []struct {
			ImageMode         string `json:"image_mode"`                    // 素材类型 (必填)
			VideoID           string `json:"video_id,omitempty"`            // 视频ID
			VideoCoverID      string `json:"video_cover_id,omitempty"`      // 视频封面图片ID
			ItemID            string `json:"item_id,omitempty"`             // 抖音短视频ID
			VideoHpVisibility string `json:"video_hp_visibility,omitempty"` // 原生单元视频素材主页可见性设置
		} `json:"video_material_list"` // 视频素材信息，单商品最多设置视频素材30个
		ImageMaterialList []struct {
			ImageMode string `json:"image_mode,omitempty"` // 素材类型
			Images    []struct {
				ImageId string `json:"image_id,omitempty"` // 图片ID
			} `json:"images"` // 图片ID数组
		} `json:"image_material_list,omitempty"` // 创意图片素材
		TitleMaterialList []struct {
			Title    string  `json:"title,omitempty"`     // 创意标题
			WordList []int64 `json:"word_list,omitempty"` // 动态词包ID
		} `json:"title_material_list"`
		ExternalUrlMaterialList []string `json:"external_url_material_list,omitempty"` // 普通落地页链接素材
		CallToActionButtons     []string `json:"call_to_action_buttons,omitempty"`     // 行动号召文案
		ComponentMaterialList   []struct {
			ComponentId int64 `json:"component_id,omitempty"` // 组件id
		} `json:"component_material_list,omitempty"` // 创意组件信息
		AnchorMaterialList []struct {
			AnchorType string `json:"anchor_type,omitempty"` // 锚点类型
			AnchorId   string `json:"anchor_id,omitempty"`   // 原生锚点id
		} `json:"anchor_material_list,omitempty"` // 原生锚点素材
		OpenURL               string `json:"open_url,omitempty"`               // 直达链接
		UlinkURL              string `json:"ulink_url,omitempty"`              // 直达备用链接
		AnchorRelatedType     string `json:"anchor_related_type,omitempty"`    // 原生锚点启用开关
		IntelligentGeneration string `json:"intelligent_generation,omitempty"` // 智能生成行动号召按钮
		PromotionMaterials    struct {
			VideoMaterialList []struct {
				ImageMode         string `json:"image_mode"`                    // 素材类型 (必填)
				VideoID           string `json:"video_id,omitempty"`            // 视频ID
				VideoCoverID      string `json:"video_cover_id,omitempty"`      // 视频封面图片ID
				ItemID            int64  `json:"item_id,omitempty"`             // 抖音短视频ID
				VideoHpVisibility string `json:"video_hp_visibility,omitempty"` // 原生单元视频素材主页可见性设置
				GuideVideoID      string `json:"guide_video_id,omitempty"`      // 引导视频ID（游戏行业奖励关卡专用）
			} `json:"video_material_list,omitempty"` // 视频素材信息
			ImageMaterialList []struct {
				ImageMode string `json:"image_mode,omitempty"` // 素材类型
				Images    []struct {
					ImageId string `json:"image_id,omitempty"` // 图片ID
				} `json:"images"` // 图片ID数组
			} `json:"image_material_list,omitempty"` // 创意图片素材
			TextAbstractList []struct {
				AbstractText string `json:"abstract_text,omitempty"` // 文本摘要内容
				BidwordList  []struct {
					DefaultWord string `json:"default_word,omitempty"` // 关键词
				} `json:"bidword_list,omitempty"` // 搜索关键词列表
				WordList []int64 `json:"word_list,omitempty"` // 动态词包ID
			} `json:"text_abstract_list"` // 文本摘要信息
			OriginalVideoTitle string `json:"original_video_title,omitempty"` // 投放原视频标题
			TitleMaterialList  []struct {
				Title       string `json:"title,omitempty"` // 创意标题
				BidwordList []struct {
					DefaultWord string `json:"default_word,omitempty"` // 关键词
				} `json:"bidword_list,omitempty"` // 搜索关键词列表
				WordList []int64 `json:"word_list,omitempty"` // 动态词包ID
			} `json:"title_material_list"`
			PlayletSeriesUrlList []string `json:"playlet_series_url_list,omitempty"` // 短剧合集链接url
			ProductInfo          struct {
				Titles        []string `json:"titles,omitempty"`         // 产品名称（条件必填）
				ImageIDs      []string `json:"image_ids,omitempty"`      // 产品主图（条件必填）
				SellingPoints []string `json:"selling_points,omitempty"` // 产品卖点（条件必填）
			} `json:"product_info,omitempty"` // 产品信息
			DecorationMaterial struct {
				ImageMode  string `json:"image_mode,omitempty"`  // 素材类型（条件必填）
				ActivityID string `json:"activity_id,omitempty"` // 活动ID（条件必填）
			} `json:"decoration_material,omitempty"` // 家装卡券素材
			AnchorMaterialList struct {
				AnchorType string `json:"anchor_type,omitempty"` // 锚点类型（条件必填）
				AnchorID   string `json:"anchor_id,omitempty"`   // 原生锚点id（条件必填）
			} `json:"anchor_material_list,omitempty"` // 原生锚点素材
			ComponentMaterialList []struct {
				ComponentID int64 `json:"component_id,omitempty"` // 组件id
			} `json:"component_material_list,omitempty"` // 创意组件信息
			ExternalUrlMaterialList []string `json:"external_url_material_list,omitempty"` // 普通落地页链接素材
			MiniProgramInfo         struct {
				URL       string   `json:"url,omitempty"`              // 小程序链接
				AppID     string   `json:"app_id,omitempty"`           // 小程序APPID（条件必填）
				StartPath string   `json:"start_path,omitempty"`       // 小程序启动路径（条件必填）
				Params    string   `json:"params,omitempty,omitempty"` // 小程序参数
				URLs      []string `json:"urls,omitempty"`             // 小程序链接列表
				Auto      []struct {
					AppID     string `json:"app_id,omitempty"`           // 小程序APPID
					StartPath string `json:"start_path,omitempty"`       // 小程序启动路径
					Params    string `json:"params,omitempty,omitempty"` // 小程序参数
				} `json:"auto,omitempty"` // 自动配置列表
			} `json:"mini_program_info,omitempty"` // 字节小程序信息
			CarouselID              string                 `json:"carousel_id,omitempty"`             // 轮播图ID
			ItemID                  int64                  `json:"item_id,omitempty"`                 // 抖音短视频ID
			VideoHpVisibility       string                 `json:"video_hp_visibility,omitempty"`     // 视频主页可见性
			TrialPlayMaterialList   []*TrialPlayMaterial   `json:"trial_play_material_list"`          // 试玩素材列表（条件必填）
			InstantPlayMaterialList []*InstantPlayMaterial `json:"instant_play_material_list"`        // 快玩素材列表（条件必填）
			DynamicCreativeSwitch   string                 `json:"dynamic_creative_switch,omitempty"` // 动态创意开关
			AdvancedDcSettings      []string               `json:"advanced_dc_settings,omitempty"`    // 高级创意设置

			CallToActionButtons          []string                        `json:"call_to_action_buttons"`           // 行动号召按钮列表（必填）
			IntelligentGeneration        string                          `json:"intelligent_generation,omitempty"` // 智能生成开关
			PlantGrassSearchWordMaterial []*PlantGrassSearchWordMaterial `json:"plant_grass_search_word_material"` // 种草搜索词素材列表

		} `json:"promotion_materials,omitempty"` // 单元素材组合

	} `json:"promotion_related_product,omitempty"` // UBP多品单元素材组合
}

// TrialPlayMaterial 试玩素材
type TrialPlayMaterial struct {
	AppPlayURI   string `json:"app_play_uri,omitempty"`   // 应用试玩URI
	GuideVideoId string `json:"guide_video_id,omitempty"` // 引导视频ID
}

// InstantPlayMaterial 快玩素材
type InstantPlayMaterial struct {
	AppPlayURI string `json:"app_play_uri,omitempty"` // 应用快玩URI
}

// PlantGrassSearchWordMaterial 种草搜索词素材
type PlantGrassSearchWordMaterial struct {
	SearchWord string `json:"search_word,omitempty"` // 搜索词
}
