package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// 构建基础组件创建落地页请求
func buildXijingPageByCompBaseReq() *model.XijingPageByCompAddReq {
	req := &model.XijingPageByCompAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageByCompPage{
		{
			PageType:    model.XijingTemplatePageTypeXijingAndroid,
			PageName:    "测试落地页",
			PageTitle:   "测试标题",
			MobileAppID: "1104790111",
			BgColor:     "rgba(189, 16, 224, 1)",
			BgImageID:   "",
			ComponentSpecList: []*model.XijingPageByCompComponentSpec{
				{
					Type: model.XijingTemplateComponentTypeFixedButton,
					FixedButtonSpec: &model.XijingPageByCompFixedButtonSpec{
						ButtonStyle: "fixedBtn-1",
						CommonSetting: &model.XijingPageByCompCommonSetting{
							Position:           "bottom",
							WhiteSpace:         80,
							DistanceToViewPort: 1,
						},
						BtnSetting: &model.XijingPageByCompBtnSetting{
							Desc:            "立即下载",
							BackgroundColor: "#1890ff",
							Color:           "rgb(255, 255, 255)",
						},
					},
				},
			},
		},
	}
	return req
}

// ========== 蹊径基于组件创建落地页接口调用测试用例 ==========

// TestXijingPageByCompAddBasicSelf 测试基本创建（固定按钮组件）
func TestXijingPageByCompAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageByCompBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageByCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageByCompAddWithImagesSelf 测试带图片列表组件
func TestXijingPageByCompAddWithImagesSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].ComponentSpecList = append(req.Pages[0].ComponentSpecList,
		&model.XijingPageByCompComponentSpec{
			Type: model.XijingTemplateComponentTypeImages,
			ImageListSpec: &model.XijingPageByCompImageListSpec{
				ImageList: []*model.XijingPageByCompImageItem{
					{
						ImageID: "1000437",
						HotArea: []*model.XijingPageByCompHotArea{
							{Width: 30, Height: 30, Top: 20, Left: 20},
						},
						Padding: &model.XijingPageByCompPadding{
							Top: 0, Right: 0, Bottom: 0, Left: 0,
						},
						ProgressBar: &model.XijingPageByCompProgressBar{
							Color:           "rgb(255, 255, 255)",
							BackgroundColor: "#1890ff",
						},
					},
				},
			},
		},
	)
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageByCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageByCompAddWithDeeplinkSelf 测试带 deeplink 和剪贴板内容
func TestXijingPageByCompAddWithDeeplinkSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].Clipboard = "剪贴板内容"
	req.Pages[0].PageDeeplink = "deeplink://deep"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageByCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageByCompAddIosSelf 测试创建 iOS 落地页
func TestXijingPageByCompAddIosSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageByCompAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageByCompPage{
		{
			PageType:    model.XijingTemplatePageTypeXijingIos,
			PageName:    "iOS落地页",
			PageTitle:   "iOS标题",
			MobileAppID: "987654321",
			BgColor:     "rgba(0, 0, 0, 1)",
			BgImageID:   "bg_img_001",
			ComponentSpecList: []*model.XijingPageByCompComponentSpec{
				{
					Type: model.XijingTemplateComponentTypeFixedButton,
					FixedButtonSpec: &model.XijingPageByCompFixedButtonSpec{
						ButtonStyle: "fixedBtn-2",
						CommonSetting: &model.XijingPageByCompCommonSetting{
							Position:           "bottom",
							WhiteSpace:         60,
							DistanceToViewPort: 0,
						},
						BtnSetting: &model.XijingPageByCompBtnSetting{
							Desc:            "下载App",
							BackgroundColor: "#ff4d4f",
							Color:           "rgb(255, 255, 255)",
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageByCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageByCompAddMultiplePagesSelf 测试批量创建
func TestXijingPageByCompAddMultiplePagesSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageByCompBaseReq()
	req.Pages = append(req.Pages, &model.XijingPageByCompPage{
		PageType:    model.XijingTemplatePageTypeXijingAndroid,
		PageName:    "第二落地页",
		PageTitle:   "第二标题",
		MobileAppID: "1104790111",
		BgColor:     "rgba(255, 255, 255, 1)",
		BgImageID:   "",
		ComponentSpecList: []*model.XijingPageByCompComponentSpec{
			{
				Type: model.XijingTemplateComponentTypeFixedButton,
				FixedButtonSpec: &model.XijingPageByCompFixedButtonSpec{
					ButtonStyle: "fixedBtn-1",
					CommonSetting: &model.XijingPageByCompCommonSetting{
						Position: "bottom", WhiteSpace: 80, DistanceToViewPort: 1,
					},
					BtnSetting: &model.XijingPageByCompBtnSetting{
						Desc: "立即安装", BackgroundColor: "#52c41a", Color: "rgb(255,255,255)",
					},
				},
			},
		},
	})
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageByCompAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径基于组件创建落地页参数验证测试用例 ==========

// TestXijingPageByCompAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageByCompAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.AccountID = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateInvalidAutoSubmitSelf 测试 is_auto_submit 非法值
func TestXijingPageByCompAddValidateInvalidAutoSubmitSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.IsAutoSubmit = 2
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：is_auto_submit非法")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateEmptyPagesSelf 测试 pages 为空
func TestXijingPageByCompAddValidateEmptyPagesSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages = []*model.XijingPageByCompPage{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：pages为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidatePagesTooManySelf 测试 pages 超过10个
func TestXijingPageByCompAddValidatePagesTooManySelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	for i := 0; i < 10; i++ {
		req.Pages = append(req.Pages, &model.XijingPageByCompPage{
			PageType:          model.XijingTemplatePageTypeXijingAndroid,
			PageName:          "额外页",
			PageTitle:         "额外标题",
			MobileAppID:       "1104790111",
			BgColor:           "rgba(0,0,0,1)",
			ComponentSpecList: []*model.XijingPageByCompComponentSpec{},
		})
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：pages超过10个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateMissingPageTypeSelf 测试缺少 page_type
func TestXijingPageByCompAddValidateMissingPageTypeSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].PageType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidatePageNameTooLongSelf 测试 page_name 超过20字节
func TestXijingPageByCompAddValidatePageNameTooLongSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].PageName = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidatePageTitleTooLongSelf 测试 page_title 超过20字节
func TestXijingPageByCompAddValidatePageTitleTooLongSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].PageTitle = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_title超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateMissingMobileAppIDSelf 测试缺少 mobile_app_id
func TestXijingPageByCompAddValidateMissingMobileAppIDSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].MobileAppID = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：mobile_app_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateMissingBgColorSelf 测试缺少 bg_color
func TestXijingPageByCompAddValidateMissingBgColorSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].BgColor = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：bg_color为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateBgImageIDTooLongSelf 测试 bg_image_id 超过256字节
func TestXijingPageByCompAddValidateBgImageIDTooLongSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].BgImageID = strings.Repeat("a", 257)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：bg_image_id超过256字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateClipboardTooLongSelf 测试 clipboard 超过300字节
func TestXijingPageByCompAddValidateClipboardTooLongSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].Clipboard = strings.Repeat("a", 301)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：clipboard超过300字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateComponentTypeMissingSelf 测试组件缺少 type
func TestXijingPageByCompAddValidateComponentTypeMissingSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Pages[0].ComponentSpecList[0].Type = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component.type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageByCompAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageByCompAddValidateFullParamsSelf(t *testing.T) {
	req := buildXijingPageByCompBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
