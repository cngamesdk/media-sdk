package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// 上传客户人群文件-IMEI
func TestCustomAudienceFilesAddIMEI(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.File = []byte("867531012345678\n867531012345679\n")
	req.FileName = "audience_imei.txt"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceFilesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 上传客户人群文件-手机号 SHA256
func TestCustomAudienceFilesAddSHA256MobilePhone(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeSHA256MobilePhone
	req.File = []byte("8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92\n")
	req.FileName = "audience_mobile.txt"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceFilesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 上传客户人群文件-微信 OpenID（带 open_app_id）
func TestCustomAudienceFilesAddWechatOpenID(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeWechatOpenID
	req.File = []byte("oABC123xxxxopenid1\noABC123xxxxopenid2\n")
	req.FileName = "audience_openid.txt"
	req.OpenAppID = "wx1234567890abcdef"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceFilesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 上传客户人群文件-REDUCE 操作（删减人群）
func TestCustomAudienceFilesAddReduce(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeOAID
	req.File = []byte("oaid_value_001\noaid_value_002\n")
	req.FileName = "audience_oaid_reduce.txt"
	req.OperationType = model.CustomAudienceFileOperationTypeReduce
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceFilesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 上传客户人群文件-Format 默认 operation_type
func TestCustomAudienceFilesAddFormatDefaultOperationType(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.File = []byte("867531012345678\n")
	req.FileName = "audience.txt"
	req.Format()
	if req.OperationType != model.CustomAudienceFileOperationTypeAppend {
		t.Fatalf("期望 operation_type 默认值为 APPEND，实际为 %s", req.OperationType)
	}
	fmt.Printf("默认值验证通过: operation_type=%s\n", req.OperationType)
}

// 验证测试-缺少 account_id
func TestCustomAudienceFilesAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.File = []byte("867531012345678\n")
	req.FileName = "audience.txt"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 audience_id
func TestCustomAudienceFilesAddValidateAudienceIDEmpty(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.File = []byte("867531012345678\n")
	req.FileName = "audience.txt"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 user_id_type
func TestCustomAudienceFilesAddValidateUserIDTypeEmpty(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.File = []byte("867531012345678\n")
	req.FileName = "audience.txt"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：user_id_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-user_id_type 值无效
func TestCustomAudienceFilesAddValidateUserIDTypeInvalid(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = "INVALID_TYPE"
	req.File = []byte("867531012345678\n")
	req.FileName = "audience.txt"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：user_id_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 file
func TestCustomAudienceFilesAddValidateFileEmpty(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.FileName = "audience.txt"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 file_name
func TestCustomAudienceFilesAddValidateFileNameEmpty(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeIMEI
	req.File = []byte("867531012345678\n")
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-open_app_id 超过最大长度
func TestCustomAudienceFilesAddValidateOpenAppIDTooLong(t *testing.T) {
	req := &model.CustomAudienceFilesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.UserIDType = model.CustomAudienceFileUserIDTypeWechatOpenID
	req.File = []byte("oABC123xxxxopenid1\n")
	req.FileName = "audience.txt"
	req.OpenAppID = "wx" + string(make([]byte, 200))
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：open_app_id长度不能超过128字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}
