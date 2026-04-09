package tencent

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 创建品牌形象测试用例 ==========

// TestBrandAddByAccountIDSelf 测试通过 account_id 创建品牌形象
func TestBrandAddByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_brand.jpg")
	if err != nil {
		t.Skip("跳过：测试品牌图片 /tmp/test_brand.jpg 不存在")
	}
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = "测试品牌形象"
	req.BrandImageFile = imageData
	req.BrandImageFileName = "test_brand.jpg"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.BrandAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBrandAddWithPngFileSelf 测试上传 PNG 格式品牌形象
func TestBrandAddWithPngFileSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_brand.png")
	if err != nil {
		t.Skip("跳过：测试品牌图片 /tmp/test_brand.png 不存在")
	}
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = "PNG格式品牌形象"
	req.BrandImageFile = imageData
	req.BrandImageFileName = "test_brand.png"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.BrandAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBrandAddWithMaxNameSelf 测试名字长度恰好 100 字节（边界值）
func TestBrandAddWithMaxNameSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_brand.jpg")
	if err != nil {
		t.Skip("跳过：测试品牌图片 /tmp/test_brand.jpg 不存在")
	}
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = strings.Repeat("a", 100) // 恰好100字节
	req.BrandImageFile = imageData
	req.BrandImageFileName = "test_brand.jpg"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.BrandAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 创建品牌形象验证测试用例 ==========

// TestBrandAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestBrandAddValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.Name = "测试品牌"
	req.BrandImageFile = []byte("fake image data")
	req.BrandImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandAddValidateMissingNameSelf 测试缺少 name
func TestBrandAddValidateMissingNameSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.BrandImageFile = []byte("fake image data")
	req.BrandImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandAddValidateNameTooLongSelf 测试 name 超过 100 字节
func TestBrandAddValidateNameTooLongSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = strings.Repeat("a", 101) // 超过100字节
	req.BrandImageFile = []byte("fake image data")
	req.BrandImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name超过100字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandAddValidateMissingImageFileSelf 测试缺少 brand_image_file
func TestBrandAddValidateMissingImageFileSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = "测试品牌"
	req.BrandImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：brand_image_file为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandAddValidateImageFileTooLargeSelf 测试 brand_image_file 超过 400KB
func TestBrandAddValidateImageFileTooLargeSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = "测试品牌"
	req.BrandImageFile = make([]byte, model.MaxBrandImageFileBytes+1) // 超过400KB
	req.BrandImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：brand_image_file超过400KB")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandAddValidateMissingImageFileNameSelf 测试缺少 brand_image_file_name
func TestBrandAddValidateMissingImageFileNameSelf(t *testing.T) {
	req := &model.BrandAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Name = "测试品牌"
	req.BrandImageFile = []byte("fake image data")
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：brand_image_file_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 获取品牌形象列表测试用例 ==========

// TestBrandGetByAccountIDSelf 测试按 account_id 获取品牌形象列表
func TestBrandGetByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BrandGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBrandGetWithPaginationSelf 测试自定义分页参数
func TestBrandGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BrandGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBrandGetWithMaxPageSizeSelf 测试 page_size 最大值 100
func TestBrandGetWithMaxPageSizeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 100
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BrandGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBrandGetDefaultPaginationSelf 测试默认分页（不传 page/page_size）
func TestBrandGetDefaultPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BrandGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取品牌形象列表验证测试用例 ==========

// TestBrandGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestBrandGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestBrandGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestBrandGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBrandGetValidatePageTooSmallSelf 测试 page 小于最小值 1
func TestBrandGetValidatePageTooSmallSelf(t *testing.T) {
	req := &model.BrandGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.Page = 0
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page小于1")
	}
	fmt.Printf("验证错误: %v\n", err)
}
