package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取创意（基础）
func TestDynamicCreativesGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按创意ID过滤
func TestDynamicCreativesGetByIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldDynamicCreativeID,
			Operator: model.OperatorIn,
			Values:   []string{"40958977", "40958978"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按广告ID过滤
func TestDynamicCreativesGetByAdgroupIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldAdgroupID,
			Operator: model.OperatorEquals,
			Values:   []string{"456"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按配置状态过滤
func TestDynamicCreativesGetByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldConfiguredStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.CreativeConfiguredStatusNormal},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 指定返回字段 + 游标分页
func TestDynamicCreativesGetWithCursorSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"dynamic_creative_id", "dynamic_creative_name", "adgroup_id", "configured_status", "creative_components"}
	req.PaginationMode = model.CreativePaginationModeCursor
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 创建创意（图片+H5落地页）
func TestDynamicCreativesAddSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.DynamicCreativeName = "test_creative_001"
	req.CreativeComponents = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				ComponentID: 111,
				Value:       &model.TextComponentValue{Content: "测试广告标题"},
			},
		},
		Image: []*model.CreativeComponent{
			{
				ComponentID: 222,
				Value: &model.ImageComponentValue{
					ImageID: "image_id_001",
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeH5,
						PageSpec: &model.PageSpec{
							H5Spec: &model.H5Spec{
								PageURL: "https://www.example.com/landing",
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 创建创意（视频+微信小程序落地页）
func TestDynamicCreativesAddWithVideoSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.DynamicCreativeName = "test_video_creative_001"
	req.DeliveryMode = model.DeliveryModeComponent
	req.ClickTrackingURL = "https://track.example.com/click"
	req.ImpressionTrackingURL = "https://track.example.com/impression"
	req.CreativeComponents = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				ComponentID: 111,
				Value:       &model.TextComponentValue{Content: "视频广告标题"},
			},
		},
		Video: []*model.CreativeComponent{
			{
				ComponentID: 333,
				Value: &model.VideoComponentValue{
					VideoID: "video_id_001",
					CoverID: "cover_id_001",
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeWechatMiniProgram,
						PageSpec: &model.PageSpec{
							WechatMiniProgramSpec: &model.WechatMiniProgramSpec{
								MiniProgramID:   "wx1234567890",
								MiniProgramPath: "pages/index/index",
							},
						},
					},
				},
			},
		},
		Brand: []*model.CreativeComponent{
			{
				ComponentID: 444,
				Value: &model.BrandComponentValue{
					BrandName:    "测试品牌",
					BrandImageID: "brand_image_001",
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新创意（基础：改名）
func TestDynamicCreativesUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DynamicCreativeID = 40958977
	req.DynamicCreativeName = "updated_creative_name_001"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新创意（替换创意组件：标题+图片）
func TestDynamicCreativesUpdateComponentsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DynamicCreativeID = 40958977
	req.CreativeComponents = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				ComponentID: 111,
				Value:       &model.TextComponentValue{Content: "更新后的广告标题"},
			},
		},
		Image: []*model.CreativeComponent{
			{
				ComponentID: 222,
				Value: &model.ImageComponentValue{
					ImageID: "image_id_new",
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeH5,
						PageSpec: &model.PageSpec{
							H5Spec: &model.H5Spec{
								PageURL: "https://www.example.com/new-landing",
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新创意（修改配置状态）
func TestDynamicCreativesUpdateStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DynamicCreativeID = 40958977
	req.ConfiguredStatus = model.CreativeConfiguredStatusNormal
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新创意（更新监测链接）
func TestDynamicCreativesUpdateTrackingURLSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DynamicCreativeID = 40958977
	req.ClickTrackingURL = "https://track.example.com/new-click"
	req.ImpressionTrackingURL = "https://track.example.com/new-impression"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 创建创意（图集+Android应用落地页）
func TestDynamicCreativesAddWithImageListSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.DynamicCreativeName = "test_imagelist_creative_001"
	req.CreativeComponents = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				ComponentID: 111,
				Value:       &model.TextComponentValue{Content: "图集广告标题"},
			},
		},
		Description: []*model.CreativeComponent{
			{
				ComponentID: 555,
				Value:       &model.TextComponentValue{Content: "广告描述文案"},
			},
		},
		ImageList: []*model.CreativeComponent{
			{
				ComponentID: 666,
				Value: &model.ImageListComponentValue{
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeAndroidApp,
						PageSpec: &model.PageSpec{
							AndroidAppSpec: &model.AndroidAppSpec{
								AndroidAppID: "com.example.app",
							},
						},
					},
					List: []*model.ImageListItem{
						{ImageID: "img_001"},
						{ImageID: "img_002"},
						{ImageID: "img_003"},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// TestDynamicCreativesDeleteSelf 测试删除动态创意
func TestDynamicCreativesDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DynamicCreativeID = 40958977
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// TestDynamicCreativesDeleteValidateAccountID 测试删除创意缺少account_id时的校验
func TestDynamicCreativesDeleteValidateAccountID(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesDeleteReq{}
	req.AccessToken = "123"
	req.DynamicCreativeID = 40958977
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.DynamicCreativesDeleteSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v", err)
}

// TestDynamicCreativesDeleteValidateCreativeID 测试删除创意缺少dynamic_creative_id时的校验
func TestDynamicCreativesDeleteValidateCreativeID(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.DynamicCreativesDeleteSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v", err)
}
