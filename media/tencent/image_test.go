package tencent

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取图片信息测试用例 ==========

// TestImageGetByAccountIDSelf 测试按 account_id 查询图片列表
func TestImageGetByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetByOrganizationIDSelf 测试按 organization_id 查询图片列表
func TestImageGetByOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.OrganizationID = 222222
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageWidthSelf 测试按图片宽度过滤
func TestImageGetFilterByImageWidthSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageIDSelf 测试按 image_id IN 过滤
func TestImageGetFilterByImageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageID, Operator: model.OperatorIn, Values: []string{"img001", "img002", "img003"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageSignatureSelf 测试按图片签名过滤
func TestImageGetFilterByImageSignatureSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageSignature, Operator: model.OperatorEquals, Values: []string{"f4c8a3bc4deb305fb74cb08ed395b98c"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByCreatedTimeSelf 测试按创建时间过滤
func TestImageGetFilterByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldCreatedTime, Operator: model.OperatorGreaterEquals, Values: []string{"2024-01-01"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterBySourceTypeSelf 测试按来源类型过滤
func TestImageGetFilterBySourceTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeLocal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByStatusSelf 测试按状态过滤
func TestImageGetFilterByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByAigcFlagSelf 测试按 AIGC 标记过滤
func TestImageGetFilterByAigcFlagSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldAigcFlag, Operator: model.OperatorEquals, Values: []string{model.ImageAigcFlagUseMuseAI}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithMultipleFiltersSelf 测试多个过滤条件（最多4个）
func TestImageGetWithMultipleFiltersSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"1920"}},
		{Field: model.ImageFilterFieldImageHeight, Operator: model.OperatorEquals, Values: []string{"1080"}},
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeAPI}},
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithPaginationSelf 测试自定义分页
func TestImageGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithLabelIDSelf 测试携带 label_id 参数
func TestImageGetWithLabelIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.LabelID = 5001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithBusinessScenarioSelf 测试携带 business_scenario 参数
func TestImageGetWithBusinessScenarioSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.BusinessScenario = 2 // 投放素材包
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestImageGetValidateMissingAccountAndOrgSelf 测试 account_id 和 organization_id 均未填写
func TestImageGetValidateMissingAccountAndOrgSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id 和 organization_id 需必填其一")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringExceedMaxSelf 测试 filtering 超过4条
func TestImageGetValidateFilteringExceedMaxSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
		{Field: model.ImageFilterFieldImageHeight, Operator: model.OperatorEquals, Values: []string{"360"}},
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeLocal}},
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
		{Field: model.ImageFilterFieldAigcFlag, Operator: model.OperatorEquals, Values: []string{model.ImageAigcFlagNotAI}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过4条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringMissingFieldSelf 测试 filtering 缺少 field
func TestImageGetValidateFilteringMissingFieldSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringMissingOperatorSelf 测试 filtering 缺少 operator
func TestImageGetValidateFilteringMissingOperatorSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Values: []string{"640"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringEmptyValuesSelf 测试 filtering values 为空
func TestImageGetValidateFilteringEmptyValuesSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestImageGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 10
	req.Format()
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestImageGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 添加图片文件测试用例（文件上传） ==========

// TestImageAddByFileWithAccountIDSelf 测试通过 account_id 上传图片文件
func TestImageAddByFileWithAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_image.jpg")
	if err != nil {
		t.Skip("跳过：测试图片文件 /tmp/test_image.jpg 不存在")
	}
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c" // 32字节 md5
	req.ImageFile = imageData
	req.ImageFileName = "test_image.jpg"
	req.Description = "测试图片上传"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageAddByFileWithOrganizationIDSelf 测试通过 organization_id 上传图片文件
func TestImageAddByFileWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_image.jpg")
	if err != nil {
		t.Skip("跳过：测试图片文件 /tmp/test_image.jpg 不存在")
	}
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.OrganizationID = 222222
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = imageData
	req.ImageFileName = "test_image.jpg"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageAddByFileWithImageUsageSelf 测试上传图片并指定图片用途
func TestImageAddByFileWithImageUsageSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_image.jpg")
	if err != nil {
		t.Skip("跳过：测试图片文件 /tmp/test_image.jpg 不存在")
	}
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = imageData
	req.ImageFileName = "test_image.jpg"
	req.ImageUsage = model.ImageUsageMarketingPendant
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageAddByFileWithResizeSelf 测试上传图片并指定裁剪尺寸
func TestImageAddByFileWithResizeSelf(t *testing.T) {
	ctx := context.Background()
	imageData, err := os.ReadFile("/tmp/test_image.jpg")
	if err != nil {
		t.Skip("跳过：测试图片文件 /tmp/test_image.jpg 不存在")
	}
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = imageData
	req.ImageFileName = "test_image.jpg"
	req.ResizeWidth = 640
	req.ResizeHeight = 480
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 添加图片文件测试用例（base64 上传） ==========

// TestImageAddByBytesSelf 测试通过 base64 编码上传图片
func TestImageAddByBytesSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeBytes
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.Bytes = "/9j/4AAQSkZJRgABAQAAAQABAAD/fake_base64_image_data"
	req.Description = "base64图片上传"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageAddByBytesWithOrganizationIDSelf 测试通过 organization_id 和 base64 编码上传图片
func TestImageAddByBytesWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.OrganizationID = 222222
	req.UploadType = model.ImageUploadTypeBytes
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.Bytes = "/9j/4AAQSkZJRgABAQAAAQABAAD/fake_base64_image_data"
	req.ImageUsage = model.ImageUsageShopImg
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, addErr := adapter.ImageAddSelf(ctx, req)
	if addErr != nil {
		t.Fatal(addErr)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 添加图片文件验证测试用例 ==========

// TestImageAddValidateMissingAccountAndOrgSelf 测试 account_id 和 organization_id 均未填写
func TestImageAddValidateMissingAccountAndOrgSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id 和 organization_id 需必填其一")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateMissingUploadTypeSelf 测试缺少 upload_type
func TestImageAddValidateMissingUploadTypeSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：upload_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateInvalidUploadTypeSelf 测试 upload_type 取值非法
func TestImageAddValidateInvalidUploadTypeSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = "UPLOAD_TYPE_INVALID"
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：upload_type取值非法")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateMissingSignatureSelf 测试缺少 signature
func TestImageAddValidateMissingSignatureSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：signature为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateSignatureWrongLengthSelf 测试 signature 长度不是 32 字节
func TestImageAddValidateSignatureWrongLengthSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "short_sig" // 不足32字节
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：signature长度必须为32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateMissingFileSelf 测试 upload_type=FILE 时缺少 file
func TestImageAddValidateMissingFileSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFileName = "test.jpg"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：UPLOAD_TYPE_FILE时file为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateMissingFileNameSelf 测试 upload_type=FILE 时缺少 image_file_name
func TestImageAddValidateMissingFileNameSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：UPLOAD_TYPE_FILE时image_file_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateMissingBytesSelf 测试 upload_type=BYTES 时缺少 bytes
func TestImageAddValidateMissingBytesSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeBytes
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：UPLOAD_TYPE_BYTES时bytes为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateDescriptionTooLongSelf 测试 description 超过 255 字节
func TestImageAddValidateDescriptionTooLongSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.Description = "这是一段超长的图片描述，用于测试字段长度校验是否正确生效。" +
		"描述内容需要超过255个字节，所以需要写足够多的文字来触发校验错误。" +
		"继续添加更多内容以确保超过255字节的限制。abcdefghijklmnopqrstuvwxyz0123456789"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：description超过255字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateResizeWidthOutOfRangeSelf 测试 resize_width 超出范围
func TestImageAddValidateResizeWidthOutOfRangeSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.ResizeWidth = 5000 // 超过最大值4000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：resize_width超过4000")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageAddValidateResizeHeightOutOfRangeSelf 测试 resize_height 超出范围
func TestImageAddValidateResizeHeightOutOfRangeSelf(t *testing.T) {
	req := &model.ImageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.UploadType = model.ImageUploadTypeFile
	req.Signature = "f4c8a3bc4deb305fb74cb08ed395b98c"
	req.ImageFile = []byte("fake image data")
	req.ImageFileName = "test.jpg"
	req.ResizeHeight = 5000 // 超过最大值4000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：resize_height超过4000")
	}
	fmt.Printf("验证错误: %v\n", err)
}
