package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// 构建基础页面请求（带顶部图片 + App 下载组件）
func buildWechatPagesCustomBaseReq() *model.WechatPagesCustomAddReq {
	req := &model.WechatPagesCustomAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageName = "测试组件原生页"
	req.PageSpecsList = []*model.WechatPagesCustomPageSpec{
		{
			PageElementsSpecList: []*model.WechatPagesCustomElementSpec{
				{
					ElementType: model.WechatPagesCustomElementTypeTopImage,
					TopImageSpec: &model.WechatPagesCustomTopImageSpec{
						ImageID:    "img001",
						Width:      800,
						Height:     800,
						AdLocation: model.WechatPagesCustomAdLocationSns,
					},
				},
			},
		},
	}
	req.ShareContentSpec = &model.WechatPageShareContentSpec{
		ShareTitle:       "分享标题",
		ShareDescription: "分享描述",
	}
	return req
}

// ========== 基于组件创建微信原生页接口调用测试用例 ==========

// TestWechatPagesCustomAddBasicSelf 测试基本创建（顶部图片+分享信息）
func TestWechatPagesCustomAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatPagesCustomBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesCustomAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesCustomAddWithAppDownloadSelf 测试带 App 下载组件
func TestWechatPagesCustomAddWithAppDownloadSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList[0].PageElementsSpecList = append(
		req.PageSpecsList[0].PageElementsSpecList,
		&model.WechatPagesCustomElementSpec{
			ElementType: model.WechatPagesCustomElementTypeAppDownload,
			AppDownloadSpec: &model.WechatPagesCustomAppDownloadSpec{
				Title: "立即下载",
				AppAndroidSpec: &model.WechatPagesCustomAppAndroidSpec{
					AppAndroidID: "com.example.app",
				},
				AppIosSpec: &model.WechatPagesCustomAppIosSpec{
					AppIosID: "123456789",
				},
			},
		},
	)
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesCustomAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesCustomAddWithGlobalFloatButtonSelf 测试带浮层组件
func TestWechatPagesCustomAddWithGlobalFloatButtonSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatPagesCustomBaseReq()
	req.GlobalSpec = &model.WechatPagesCustomGlobalSpec{
		GlobalElementsSpecList: []*model.WechatPagesCustomGlobalElementSpec{
			{
				ElementType: model.WechatPagesCustomGlobalElementTypeFloatButton,
				FloatButtonSpec: &model.WechatPagesCustomFloatButtonSpec{
					StyleType:   2,
					Title:       "立即下载",
					ElementType: model.WechatPagesCustomElementTypeAppDownload,
					AppDownloadSpec: &model.WechatPagesCustomAppDownloadSpec{
						Title: "立即下载",
						AppAndroidSpec: &model.WechatPagesCustomAppAndroidSpec{
							AppAndroidID: "com.example.app",
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesCustomAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesCustomAddWithSideBarFloatSelf 测试带侧边栏浮层组件（一键拨号）
func TestWechatPagesCustomAddWithSideBarFloatSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatPagesCustomBaseReq()
	req.GlobalSpec = &model.WechatPagesCustomGlobalSpec{
		GlobalElementsSpecList: []*model.WechatPagesCustomGlobalElementSpec{
			{
				ElementType: model.WechatPagesCustomGlobalElementTypeSideBarFloat,
				SideBarFloatSpec: &model.WechatPagesCustomSideBarFloatSpec{
					ElemType: model.WechatPagesCustomSideBarElemTypeTel,
					TelSpec: &model.WechatPagesCustomTelSpec{
						PhoneNumber: "13800138000",
						PhoneType:   "1",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesCustomAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 基于组件创建微信原生页参数验证测试用例 ==========

// TestWechatPagesCustomAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatPagesCustomAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.AccountID = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateEmptyPageNameSelf 测试 page_name 为空
func TestWechatPagesCustomAddValidateEmptyPageNameSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageName = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidatePageNameTooLongSelf 测试 page_name 超过120字节
func TestWechatPagesCustomAddValidatePageNameTooLongSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	name := ""
	for i := 0; i < 121; i++ {
		name += "a"
	}
	req.PageName = name
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过120字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateEmptyPageSpecsListSelf 测试 page_specs_list 为空
func TestWechatPagesCustomAddValidateEmptyPageSpecsListSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList = []*model.WechatPagesCustomPageSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_specs_list为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateEmptyElementsListSelf 测试 page_elements_spec_list 为空
func TestWechatPagesCustomAddValidateEmptyElementsListSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList[0].PageElementsSpecList = []*model.WechatPagesCustomElementSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_elements_spec_list为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateNilElementSelf 测试 page_elements_spec_list 包含 nil 元素
func TestWechatPagesCustomAddValidateNilElementSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList[0].PageElementsSpecList = append(
		req.PageSpecsList[0].PageElementsSpecList, nil,
	)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_elements_spec_list包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateMissingElementTypeSelf 测试 element_type 为空
func TestWechatPagesCustomAddValidateMissingElementTypeSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList[0].PageElementsSpecList[0].ElementType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateMissingShareContentSpecSelf 测试缺少 share_content_spec
func TestWechatPagesCustomAddValidateMissingShareContentSpecSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.ShareContentSpec = nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：share_content_spec为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateMissingShareTitleSelf 测试 share_title 为空
func TestWechatPagesCustomAddValidateMissingShareTitleSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.ShareContentSpec.ShareTitle = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：share_title为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateMissingShareDescriptionSelf 测试 share_description 为空
func TestWechatPagesCustomAddValidateMissingShareDescriptionSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.ShareContentSpec.ShareDescription = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：share_description为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesCustomAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatPagesCustomAddValidateFullParamsSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestWechatPagesCustomAddValidateMultiplePagesSelf 测试多页面配置通过验证
func TestWechatPagesCustomAddValidateMultiplePagesSelf(t *testing.T) {
	req := buildWechatPagesCustomBaseReq()
	req.PageSpecsList = append(req.PageSpecsList, &model.WechatPagesCustomPageSpec{
		BgColor: "#FFFFFF",
		PageElementsSpecList: []*model.WechatPagesCustomElementSpec{
			{
				ElementType: model.WechatPagesCustomElementTypeText,
				TextSpec: &model.WechatPagesCustomTextSpec{
					Text: "第二页文案内容",
				},
			},
		},
	})
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("多页面配置应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("多页面配置验证通过")
}
