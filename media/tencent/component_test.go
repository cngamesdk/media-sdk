package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// TestComponentsGetSelf 测试获取全部创意组件（不加过滤条件）
func TestComponentsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByComponentIDSelf 测试按组件id过滤获取创意组件
func TestComponentsGetByComponentIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldComponentID,
			Operator: model.OperatorIn,
			Values:   []string{"111111", "222222"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetBySubTypeSelf 测试按组件子类型过滤获取创意组件
func TestComponentsGetBySubTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldComponentSubType,
			Operator: model.OperatorEquals,
			Values:   []string{"COMPONENT_SUB_TYPE_IMAGE"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByCreatedTimeSelf 测试按创建时间过滤获取创意组件
func TestComponentsGetByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldCreatedTime,
			Operator: model.OperatorGreaterEquals,
			Values:   []string{"1704067200"}, // 2024-01-01 00:00:00
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByVideoIDSelf 测试按视频id过滤获取创意组件
func TestComponentsGetByVideoIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldVideoID,
			Operator: model.OperatorIn,
			Values:   []string{"video_id_001", "video_id_002"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByImageIDSelf 测试按图片id过滤获取创意组件
func TestComponentsGetByImageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldImageID,
			Operator: model.OperatorEquals,
			Values:   []string{"image_id_001"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByPotentialStatusSelf 测试按潜力状态过滤获取创意组件
func TestComponentsGetByPotentialStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldPotentialStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.ComponentPotentialStatusHigh},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithPaginationSelf 测试分页获取创意组件
func TestComponentsGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithFieldsSelf 测试指定返回字段获取创意组件
func TestComponentsGetWithFieldsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"component_id", "component_sub_type", "component_custom_name", "created_time"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetDeletedSelf 测试获取已删除的创意组件
func TestComponentsGetDeletedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.IsDeleted = true
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithOrganizationIDSelf 测试通过业务单元获取创意组件
func TestComponentsGetWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.OrganizationID = 456
	req.ComponentIDFilteringMode = model.ComponentIDFilteringModeSharingByCustomerBusinessUnit
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetValidateAccountID 测试缺少account_id时的校验
func TestComponentsGetValidateAccountID(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsGetSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// ========== 创建创意组件测试用例 ==========

// TestComponentsAddTitleSelf 测试创建标题组件
func TestComponentsAddTitleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeTitle
	req.ComponentCustomName = "测试标题组件"
	req.ComponentValue = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				Value: &model.TextComponentValue{Content: "这是一个广告标题"},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddDescriptionSelf 测试创建描述组件
func TestComponentsAddDescriptionSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeDescription
	req.ComponentCustomName = "测试描述组件"
	req.ComponentValue = &model.CreativeComponents{
		Description: []*model.CreativeComponent{
			{
				Value: &model.TextComponentValue{Content: "这是广告描述文案内容"},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddImageSelf 测试创建单图组件（H5落地页）
func TestComponentsAddImageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeImage16X9
	req.ComponentCustomName = "测试单图16:9组件"
	req.ComponentValue = &model.CreativeComponents{
		Image: []*model.CreativeComponent{
			{
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
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddImageWithMiniProgramSelf 测试创建单图组件（微信小程序落地页）
func TestComponentsAddImageWithMiniProgramSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeImage1X1
	req.ComponentValue = &model.CreativeComponents{
		Image: []*model.CreativeComponent{
			{
				Value: &model.ImageComponentValue{
					ImageID: "image_id_001",
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
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddVideoSelf 测试创建视频组件（Android应用落地页）
func TestComponentsAddVideoSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeVideo16X9
	req.ComponentCustomName = "测试视频16:9组件"
	req.ComponentValue = &model.CreativeComponents{
		Video: []*model.CreativeComponent{
			{
				Value: &model.VideoComponentValue{
					VideoID: "video_id_001",
					CoverID: "cover_id_001",
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeAndroidApp,
						PageSpec: &model.PageSpec{
							AndroidAppSpec: &model.AndroidAppSpec{
								AndroidAppID: "com.example.app",
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddImageListSelf 测试创建图集组件
func TestComponentsAddImageListSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeImageList1X1_3
	req.ComponentValue = &model.CreativeComponents{
		ImageList: []*model.CreativeComponent{
			{
				Value: &model.ImageListComponentValue{
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeH5,
						PageSpec: &model.PageSpec{
							H5Spec: &model.H5Spec{PageURL: "https://www.example.com/landing"},
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
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddActionButtonSelf 测试创建行动按钮组件
func TestComponentsAddActionButtonSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeActionButton
	req.ComponentValue = &model.CreativeComponents{
		ActionButton: []*model.CreativeComponent{
			{
				Value: &model.ActionButtonComponentValue{
					ButtonText: "立即下载",
					JumpInfo: &model.JumpInfo{
						PageType: model.PageTypeAppDeepLink,
						PageSpec: &model.PageSpec{
							AppDeepLinkSpec: &model.AppDeepLinkSpec{
								AndroidDeepLinkURL: "myapp://main",
								IosDeepLinkURL:     "myapp://main",
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddLabelSelf 测试创建标签组件
func TestComponentsAddLabelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeLabel
	req.ComponentValue = &model.CreativeComponents{
		Label: []*model.CreativeComponent{
			{
				Value: &model.LabelComponentValue{
					List: []*model.LabelItem{
						{
							Content:        "限时优惠",
							Type:           model.LabelTypePromotional,
							DisplayContent: "限时特惠",
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddBarrageSelf 测试创建弹幕组件
func TestComponentsAddBarrageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeBarrage
	req.ComponentValue = &model.CreativeComponents{
		Barrage: []*model.CreativeComponent{
			{
				Value: &model.BarrageComponentValue{
					List: []*model.BarrageItem{
						{Text: "好产品"},
						{Text: "强烈推荐"},
						{Text: "性价比超高"},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddCountDownSelf 测试创建倒计时组件
func TestComponentsAddCountDownSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeCountDown
	req.ComponentValue = &model.CreativeComponents{
		CountDown: []*model.CreativeComponent{
			{
				Value: &model.CountDownComponentValue{
					Price:             "9900",
					TimeType:          model.CountdownTimeEnd,
					ExpiringTimestamp: 1800000000,
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddWithOrganizationIDSelf 测试通过业务单元创建组件
func TestComponentsAddWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.OrganizationID = 456
	req.ComponentSubType = model.ComponentSubTypeTitle
	req.ComponentValue = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{
				Value: &model.TextComponentValue{Content: "业务单元共享标题组件"},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsAddValidateSubTypeSelf 测试缺少component_sub_type时的校验
func TestComponentsAddValidateSubTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentValue = &model.CreativeComponents{
		Title: []*model.CreativeComponent{
			{Value: &model.TextComponentValue{Content: "测试"}},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsAddSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestComponentsAddValidateComponentValueSelf 测试缺少component_value时的校验
func TestComponentsAddValidateComponentValueSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentSubType = model.ComponentSubTypeTitle
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsAddSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// ========== 删除创意组件测试用例 ==========

// TestComponentsDeleteSelf 测试删除创意组件（默认策略）
func TestComponentsDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsDeleteWithForceStrategySelf 测试强制删除创意组件
func TestComponentsDeleteWithForceStrategySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentID = 111111
	req.DeleteStrategy = model.DeleteStrategyForce
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsDeleteWithRestrictedStrategySelf 测试受限删除创意组件
func TestComponentsDeleteWithRestrictedStrategySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentID = 111111
	req.DeleteStrategy = model.DeleteStrategyRestricted
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsDeleteWithOrganizationIDSelf 测试通过业务单元删除创意组件
func TestComponentsDeleteWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.OrganizationID = 456
	req.ComponentID = 111111
	req.DeleteStrategy = model.DeleteStrategyRestricted
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsDeleteValidateComponentIDSelf 测试缺少component_id时的校验
func TestComponentsDeleteValidateComponentIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestComponentsDeleteValidateStrategySelf 测试传入无效delete_strategy时的校验
func TestComponentsDeleteValidateStrategySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.ComponentID = 111111
	req.DeleteStrategy = "INVALID_STRATEGY"
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsDeleteSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}
