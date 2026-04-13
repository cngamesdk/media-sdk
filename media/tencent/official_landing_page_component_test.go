package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildOfficialLandingPageCompBaseReq 构建基础官方落地页组件创建请求（顶部外显素材 + 按钮组件）
func buildOfficialLandingPageCompBaseReq() *model.OfficialLandingPageCompAddReq {
	req := &model.OfficialLandingPageCompAddReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.ProtoVersion = 1
	req.PageConfig = &model.OfficialLandingPageConfig{
		PageName:     "落地页名称",
		PageTitle:    "落地页标题",
		IosAppId:     "11111111",
		AndroidAppId: "11111111",
	}
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType:               model.OfficialLandingPageElementTypeHeadOutsideMaterial,
			HeadOutsideMaterialConfig: []interface{}{},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:         "立即下载",
				Color:           "#FFFFFF",
				BackgroundColor: "#296BEF",
				HasIcon:         true,
				ButtonSize:      "large",
				MarginTop:       10,
				MarginBottom:    10,
				ConvertDownload: &model.OfficialLandingPageConvertDownload{
					ConvertType:        model.OfficialLandingPageConvertTypeDownload,
					DeeplinkUrlAndroid: "app://android",
					DeeplinkUrlIos:     "app://ios",
				},
			},
		},
	}
	return req
}

// ========== 官方落地页基于组件创建接口调用测试用例 ==========

// TestOfficialLandingPageCompAddBasicSelf 测试基本创建（顶部外显+按钮组件）
func TestOfficialLandingPageCompAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithHeadVideoSelf 测试带顶部视频组件
func TestOfficialLandingPageCompAddWithHeadVideoSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeHeadVideo,
			HeadVideoConfig: &model.OfficialLandingPageHeadVideoConfig{
				MaterialId: "video_material_id_001",
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:    "立即下载",
				ButtonSize: "large",
				ConvertDownload: &model.OfficialLandingPageConvertDownload{
					ConvertType: model.OfficialLandingPageConvertTypeDownload,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithHeadImageSelf 测试带顶部图片组件
func TestOfficialLandingPageCompAddWithHeadImageSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeHeadImage,
			HeadImageConfig: &model.OfficialLandingPageHeadImageConfig{
				MaterialId: "image_material_id_001",
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeFixedButton,
			FixedButtonConfig: &model.OfficialLandingPageFixedButtonConfig{
				Title:         "立即体验",
				Desc:          "下载享好礼",
				ButtonContent: "下载",
				ConvertDownload: &model.OfficialLandingPageConvertDownload{
					ConvertType: model.OfficialLandingPageConvertTypeDownload,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithCarouselSelf 测试带轮播图组件
func TestOfficialLandingPageCompAddWithCarouselSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeHeadCarousel,
			HeadCarouselConfig: &model.OfficialLandingPageHeadCarouselConfig{
				MaterialIdList: []string{"img_id_001", "img_id_002", "img_id_003"},
				Type:           "full",
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockCarousel,
			BlockCarouselConfig: &model.OfficialLandingPageBlockCarouselConfig{
				MaterialIdList: []string{"img_id_004", "img_id_005"},
				Type:           "center",
				MarginTop:      10,
				MarginBottom:   10,
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:    "立即下载",
				ButtonSize: "large",
				ConvertDownload: &model.OfficialLandingPageConvertDownload{
					ConvertType: model.OfficialLandingPageConvertTypeDownload,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithTextComponentsSelf 测试带文本和分割线组件
func TestOfficialLandingPageCompAddWithTextComponentsSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeHeadText,
			HeadTextConfig: &model.OfficialLandingPageHeadTextConfig{
				Title: &model.OfficialLandingPageTextContent{
					Content: "标题文案",
					Color:   "#000000",
				},
				Detail: &model.OfficialLandingPageTextContent{
					Content: "详情内容",
					Color:   "#666666",
				},
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockText,
			BlockTextConfig: &model.OfficialLandingPageBlockTextConfig{
				Content:      "这是一段普通文本内容",
				Color:        "#333333",
				FontSize:     "16",
				FontWeight:   "normal",
				FontStyle:    "normal",
				TextAlign:    "left",
				MarginTop:    10,
				MarginBottom: 10,
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockDivider,
			BlockDividerConfig: &model.OfficialLandingPageBlockDividerConfig{
				LineColor:    "#E5E5E5",
				Thickness:    1,
				Type:         "solid",
				MarginTop:    10,
				MarginBottom: 10,
			},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:    "立即下载",
				ButtonSize: "large",
				ConvertDownload: &model.OfficialLandingPageConvertDownload{
					ConvertType: model.OfficialLandingPageConvertTypeDownload,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithWeappConvertSelf 测试转化-打开小程序
func TestOfficialLandingPageCompAddWithWeappConvertSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType:               model.OfficialLandingPageElementTypeHeadOutsideMaterial,
			HeadOutsideMaterialConfig: []interface{}{},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:    "打开小程序",
				ButtonSize: "large",
				ConvertWeapp: &model.OfficialLandingPageConvertWeapp{
					ConvertType: model.OfficialLandingPageConvertTypeWeapp,
					WeappId:     "gh_xxxxxxxxxxxx",
					WeappPath:   "pages/index/index",
					BackupLink:  "https://example.com/backup",
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithFollowGhSelf 测试转化-关注公众号
func TestOfficialLandingPageCompAddWithFollowGhSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.IosAppId = ""
	req.PageConfig.AndroidAppId = ""
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType:               model.OfficialLandingPageElementTypeHeadOutsideMaterial,
			HeadOutsideMaterialConfig: []interface{}{},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockShelf,
			BlockShelfConfig: &model.OfficialLandingPageBlockShelfConfig{
				CardType:      "card",
				Title:         "关注我们",
				Desc:          "获取更多资讯",
				ButtonContent: "关注",
				MarginTop:     10,
				MarginBottom:  10,
				ConvertGh: &model.OfficialLandingPageConvertGh{
					ConvertType: model.OfficialLandingPageConvertTypeGh,
					Appid:       "wxffffffffffff",
					OneClick:    true,
					DirectFocus: true,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithSideButtonSelf 测试侧边悬浮按钮组件
func TestOfficialLandingPageCompAddWithSideButtonSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = append(req.PageElements, &model.OfficialLandingPageElement{
		ElementType: model.OfficialLandingPageElementTypeRightFixedButton,
		RightFixedButtonConfig: &model.OfficialLandingPageRightFixedButtonConfig{
			BackgroundColor: "#FFFFFF",
			Type:            "with-button",
			Content: []*model.OfficialLandingPageRightFixedButtonContent{
				{
					Title:         "下载",
					TitleColor:    "#000000",
					ButtonContent: "下载",
					ButtonBgColor: "#296BEF",
					ConvertDownload: &model.OfficialLandingPageConvertDownload{
						ConvertType: model.OfficialLandingPageConvertTypeDownload,
					},
				},
			},
		},
	})
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithLinkConvertSelf 测试转化-跳转链接
func TestOfficialLandingPageCompAddWithLinkConvertSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType:               model.OfficialLandingPageElementTypeHeadOutsideMaterial,
			HeadOutsideMaterialConfig: []interface{}{},
		},
		{
			ElementType: model.OfficialLandingPageElementTypeBlockButton,
			BlockButtonConfig: &model.OfficialLandingPageBlockButtonConfig{
				Content:    "了解详情",
				ButtonSize: "large",
				ConvertLink: &model.OfficialLandingPageConvertLink{
					ConvertType:        model.OfficialLandingPageConvertTypeLink,
					Src:                "https://example.com/landing",
					AppIdIos:           "com.example.app",
					DeeplinkUrlIos:     "app://example/ios",
					AppIdAndroid:       "com.example.android",
					DeeplinkUrlAndroid: "app://example/android",
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithImageHotAreaSelf 测试带热区的图片组件
func TestOfficialLandingPageCompAddWithImageHotAreaSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeBlockImage,
			BlockImageConfig: &model.OfficialLandingPageBlockImageConfig{
				MaterialId: "image_material_id_002",
				Areas: []*model.OfficialLandingPageHotArea{
					{
						Left:   10,
						Top:    20,
						Width:  30,
						Height: 20,
						ConvertDownload: &model.OfficialLandingPageConvertDownload{
							ConvertType: model.OfficialLandingPageConvertTypeDownload,
						},
					},
				},
				MarginTop:    5,
				MarginBottom: 5,
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithShareConfigSelf 测试带分享配置
func TestOfficialLandingPageCompAddWithShareConfigSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.ShareTitle = "分享标题"
	req.PageConfig.ShareDescription = "这是分享描述内容"
	req.PageConfig.ShareThumbUrlMaterialId = "thumb_material_id"
	req.PageConfig.BgColor = "#F5F5F5"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageCompAddWithDoubleShelfSelf 测试双图文复合组件
func TestOfficialLandingPageCompAddWithDoubleShelfSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{
		{
			ElementType: model.OfficialLandingPageElementTypeBlockShelfGroup,
			BlockShelfGroupConfig: &model.OfficialLandingPageBlockShelfGroupConfig{
				TextAlign:    "left",
				MarginTop:    10,
				MarginBottom: 10,
				Content: []*model.OfficialLandingPageShelfGroupItem{
					{
						Title:          "产品一",
						Desc:           "产品一描述内容",
						ButtonContent:  "下载",
						IconMaterialId: "icon_material_id_001",
						ConvertDownload: &model.OfficialLandingPageConvertDownload{
							ConvertType: model.OfficialLandingPageConvertTypeDownload,
						},
					},
					{
						Title:          "产品二",
						Desc:           "产品二描述内容",
						ButtonContent:  "下载",
						IconMaterialId: "icon_material_id_002",
						ConvertDownload: &model.OfficialLandingPageConvertDownload{
							ConvertType: model.OfficialLandingPageConvertTypeDownload,
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 官方落地页基于组件创建参数验证测试用例 ==========

// TestOfficialLandingPageCompAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestOfficialLandingPageCompAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateMissingPageConfigSelf 测试缺少 page_config
func TestOfficialLandingPageCompAddValidateMissingPageConfigSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig = nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_config为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidatePageNameEmptySelf 测试 page_name 为空
func TestOfficialLandingPageCompAddValidatePageNameEmptySelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.PageName = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidatePageNameTooLongSelf 测试 page_name 超过20字节
func TestOfficialLandingPageCompAddValidatePageNameTooLongSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.PageName = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidatePageTitleEmptySelf 测试 page_title 为空
func TestOfficialLandingPageCompAddValidatePageTitleEmptySelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.PageTitle = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_title为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidatePageTitleTooLongSelf 测试 page_title 超过20字节
func TestOfficialLandingPageCompAddValidatePageTitleTooLongSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.PageTitle = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_title超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateShareTitleTooLongSelf 测试 share_title 超过14字节
func TestOfficialLandingPageCompAddValidateShareTitleTooLongSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.ShareTitle = strings.Repeat("a", 15)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：share_title超过14字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateShareDescTooLongSelf 测试 share_description 超过20字节
func TestOfficialLandingPageCompAddValidateShareDescTooLongSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageConfig.ShareDescription = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：share_description超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateElementsTooManySelf 测试 page_elements 超过10个
func TestOfficialLandingPageCompAddValidateElementsTooManySelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	for i := 0; i < 9; i++ {
		req.PageElements = append(req.PageElements, &model.OfficialLandingPageElement{
			ElementType: model.OfficialLandingPageElementTypeBlockDivider,
			BlockDividerConfig: &model.OfficialLandingPageBlockDividerConfig{
				LineColor: "#E5E5E5",
			},
		})
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_elements超过10个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateMissingElementTypeSelf 测试组件缺少 element_type
func TestOfficialLandingPageCompAddValidateMissingElementTypeSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements[0].ElementType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateInvalidProtoVersionSelf 测试非法的 proto_version
func TestOfficialLandingPageCompAddValidateInvalidProtoVersionSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.ProtoVersion = 2
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：proto_version须为0或1")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageCompAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestOfficialLandingPageCompAddValidateFullParamsSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestOfficialLandingPageCompAddValidateEmptyPageElementsSelf 测试 page_elements 为空切片（允许）
func TestOfficialLandingPageCompAddValidateEmptyPageElementsSelf(t *testing.T) {
	req := buildOfficialLandingPageCompBaseReq()
	req.PageElements = []*model.OfficialLandingPageElement{}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("空 page_elements 应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("空 page_elements 验证通过")
}
