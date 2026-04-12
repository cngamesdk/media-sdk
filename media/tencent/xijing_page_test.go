package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// 构建基础创建落地页请求
func buildXijingPageAddBaseReq() *model.XijingPageAddReq {
	req := &model.XijingPageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageAddPage{
		{
			PageType:       model.XijingTemplatePageTypeXijingAndroid,
			PageName:       "测试落地页",
			PageTitle:      "测试标题",
			PageTemplateID: "1006",
			MobileAppID:    "1104790111",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{
					Type:      model.XijingTemplateComponentTypeVideo,
					VideoSpec: &model.XijingTemplateVideoSpec{VideoID: "1540"},
				},
				{
					Type:     model.XijingTemplateComponentTypeText,
					TextSpec: &model.XijingTemplateTextSpec{Text: "hi 蹊径落地页"},
				},
			},
		},
	}
	return req
}

// ========== 蹊径基于模板创建落地页接口调用测试用例 ==========

// TestXijingPageAddBasicSelf 测试基本创建（Android 落地页）
func TestXijingPageAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageAddBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageAddWithAllComponentsSelf 测试带全部组件类型
func TestXijingPageAddWithAllComponentsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageAddPage{
		{
			PageType:       model.XijingTemplatePageTypeXijingAndroid,
			PageName:       "全组件落地页",
			PageTitle:      "全组件落地页标题",
			PageTemplateID: "1006",
			MobileAppID:    "1104790111",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{
					Type:      model.XijingTemplateComponentTypeVideo,
					VideoSpec: &model.XijingTemplateVideoSpec{VideoID: "1540"},
				},
				{
					Type:     model.XijingTemplateComponentTypeText,
					TextSpec: &model.XijingTemplateTextSpec{Text: "落地页描述文案"},
				},
				{
					Type: model.XijingTemplateComponentTypeImages,
					ImageListSpec: &model.XijingTemplateImageListSpec{
						ImageList: []*model.XijingTemplateImageItem{
							{ImageID: "1000437"},
						},
					},
				},
				{
					Type:       model.XijingTemplateComponentTypeButton,
					ButtonSpec: &model.XijingTemplateButtonSpec{Text: "立即下载"},
				},
				{
					Type:              model.XijingTemplateComponentTypeAppInfoButton,
					AppInfoButtonSpec: &model.XijingTemplateAppInfoButtonSpec{Text: "底部下载图的描述"},
				},
				{
					Type:            model.XijingTemplateComponentTypeFixedButton,
					FixedButtonSpec: &model.XijingTemplateFixedButtonSpec{Desc: "下载描述文案"},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageAddIosSelf 测试创建 iOS 落地页
func TestXijingPageAddIosSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageAddPage{
		{
			PageType:       model.XijingTemplatePageTypeXijingIos,
			PageName:       "iOS落地页",
			PageTitle:      "iOS落地页标题",
			PageTemplateID: "2001",
			MobileAppID:    "987654321",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{
					Type:     model.XijingTemplateComponentTypeText,
					TextSpec: &model.XijingTemplateTextSpec{Text: "iOS专属文案"},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageAddWebsiteSelf 测试创建网站落地页
func TestXijingPageAddWebsiteSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 1
	req.Pages = []*model.XijingPageAddPage{
		{
			PageType:       model.XijingPageTypeXijingWebsite,
			PageName:       "网站落地页",
			PageTitle:      "网站落地页标题",
			PageTemplateID: "3001",
			FormID:         "form001",
			Clipboard:      "剪贴板内容",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{
					Type:     model.XijingTemplateComponentTypeText,
					TextSpec: &model.XijingTemplateTextSpec{Text: "网站落地页文案"},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageAddMultiplePagesSelf 测试批量创建多个落地页
func TestXijingPageAddMultiplePagesSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.Pages = []*model.XijingPageAddPage{
		{
			PageType:       model.XijingTemplatePageTypeXijingAndroid,
			PageName:       "Android落地页A",
			PageTitle:      "标题A",
			PageTemplateID: "1006",
			MobileAppID:    "1104790111",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{Type: model.XijingTemplateComponentTypeText, TextSpec: &model.XijingTemplateTextSpec{Text: "文案A"}},
			},
		},
		{
			PageType:       model.XijingTemplatePageTypeXijingAndroid,
			PageName:       "Android落地页B",
			PageTitle:      "标题B",
			PageTemplateID: "1006",
			MobileAppID:    "1104790111",
			ComponentSpecList: []*model.XijingPageComponentSpec{
				{Type: model.XijingTemplateComponentTypeText, TextSpec: &model.XijingTemplateTextSpec{Text: "文案B"}},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径基于模板创建落地页参数验证测试用例 ==========

// TestXijingPageAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.AccountID = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateInvalidAutoSubmitSelf 测试 is_auto_submit 非法值
func TestXijingPageAddValidateInvalidAutoSubmitSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.IsAutoSubmit = 2
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：is_auto_submit非法")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateEmptyPagesSelf 测试 pages 为空
func TestXijingPageAddValidateEmptyPagesSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages = []*model.XijingPageAddPage{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：pages为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidatePagesTooManySelf 测试 pages 超过10个
func TestXijingPageAddValidatePagesTooManySelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	for i := 0; i < 10; i++ {
		req.Pages = append(req.Pages, &model.XijingPageAddPage{
			PageType:          model.XijingTemplatePageTypeXijingAndroid,
			PageName:          "额外页",
			PageTitle:         "额外标题",
			PageTemplateID:    "1006",
			ComponentSpecList: []*model.XijingPageComponentSpec{},
		})
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：pages超过10个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateMissingPageTypeSelf 测试缺少 page_type
func TestXijingPageAddValidateMissingPageTypeSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].PageType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidatePageNameTooLongSelf 测试 page_name 超过20字节
func TestXijingPageAddValidatePageNameTooLongSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].PageName = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidatePageTitleTooLongSelf 测试 page_title 超过20字节
func TestXijingPageAddValidatePageTitleTooLongSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].PageTitle = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_title超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateClipboardTooLongSelf 测试 clipboard 超过300字节
func TestXijingPageAddValidateClipboardTooLongSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].Clipboard = strings.Repeat("a", 301)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：clipboard超过300字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateMissingTemplateIDSelf 测试缺少 page_template_id
func TestXijingPageAddValidateMissingTemplateIDSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].PageTemplateID = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_template_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateTemplateIDTooLongSelf 测试 page_template_id 超过32字节
func TestXijingPageAddValidateTemplateIDTooLongSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].PageTemplateID = strings.Repeat("a", 33)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_template_id超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateComponentTypeMissingSelf 测试组件缺少 type
func TestXijingPageAddValidateComponentTypeMissingSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Pages[0].ComponentSpecList[0].Type = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component.type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageAddValidateFullParamsSelf(t *testing.T) {
	req := buildXijingPageAddBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// ========== 蹊径送审落地页接口调用测试用例 ==========

// TestXijingPageUpdateBasicSelf 测试送审单个落地页
func TestXijingPageUpdateBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = []string{"576460752303438398"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageUpdateMultiplePagesSelf 测试批量送审多个落地页
func TestXijingPageUpdateMultiplePagesSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = []string{
		"576460752303438398",
		"576460752303438399",
		"576460752303438400",
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageUpdateNotSubmitSelf 测试 is_submitted_for_review=false
func TestXijingPageUpdateNotSubmitSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = false
	req.PageIDList = []string{"576460752303438398"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageUpdateEmptyListSelf 测试空落地页列表（合法，最小长度为0）
func TestXijingPageUpdateEmptyListSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = []string{}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径送审落地页参数验证测试用例 ==========

// TestXijingPageUpdateValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageUpdateValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 0
	req.IsSubmittedForReview = true
	req.PageIDList = []string{"576460752303438398"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageUpdateValidateNilPageIDListSelf 测试 page_id_list 为 nil
func TestXijingPageUpdateValidateNilPageIDListSelf(t *testing.T) {
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageUpdateValidatePageIDListTooManySelf 测试 page_id_list 超过999个
func TestXijingPageUpdateValidatePageIDListTooManySelf(t *testing.T) {
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = make([]string, model.MaxXijingPageUpdatePageIDListCount+1)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id_list超过999个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageUpdateValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageUpdateValidateFullParamsSelf(t *testing.T) {
	req := &model.XijingPageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsSubmittedForReview = true
	req.PageIDList = []string{"576460752303438398"}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// ========== 蹊径删除落地页接口调用测试用例 ==========

// TestXijingPageDeleteBasicSelf 测试删除单个落地页
func TestXijingPageDeleteBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = []string{"576460752303438398"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageDeleteMultiplePagesSelf 测试批量删除多个落地页
func TestXijingPageDeleteMultiplePagesSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = []string{
		"576460752303438398",
		"576460752303438399",
		"576460752303438400",
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageDeleteEmptyListSelf 测试空落地页列表（合法，最小长度为0）
func TestXijingPageDeleteEmptyListSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = []string{}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径删除落地页参数验证测试用例 ==========

// TestXijingPageDeleteValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageDeleteValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 0
	req.PageIDList = []string{"576460752303438398"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageDeleteValidateNilPageIDListSelf 测试 page_id_list 为 nil
func TestXijingPageDeleteValidateNilPageIDListSelf(t *testing.T) {
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageDeleteValidatePageIDListTooManySelf 测试 page_id_list 超过999个
func TestXijingPageDeleteValidatePageIDListTooManySelf(t *testing.T) {
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = make([]string, model.MaxXijingPageDeletePageIDListCount+1)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id_list超过999个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageDeleteValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageDeleteValidateFullParamsSelf(t *testing.T) {
	req := &model.XijingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIDList = []string{"576460752303438398"}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
