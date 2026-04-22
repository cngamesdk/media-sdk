package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 新增广告主违规申述-基本调用
func TestIllegalComplaintAdd(t *testing.T) {
	ctx := context.Background()
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.IllegalOrderID = "ORDER_001"
	req.ComplaintReason = "该广告内容合规，已取得相关资质"
	req.File = []byte("fake zip content")
	req.FileName = "evidence.zip"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.IllegalComplaintAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestIllegalComplaintAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.IllegalOrderID = "ORDER_001"
	req.ComplaintReason = "合规"
	req.File = []byte("fake")
	req.FileName = "evidence.zip"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 illegal_order_id
func TestIllegalComplaintAddValidateIllegalOrderIDEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.ComplaintReason = "合规"
	req.File = []byte("fake")
	req.FileName = "evidence.zip"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：illegal_order_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 complaint_reason
func TestIllegalComplaintAddValidateComplaintReasonEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.IllegalOrderID = "ORDER_001"
	req.File = []byte("fake")
	req.FileName = "evidence.zip"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：complaint_reason为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 file
func TestIllegalComplaintAddValidateFileEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.IllegalOrderID = "ORDER_001"
	req.ComplaintReason = "合规"
	req.FileName = "evidence.zip"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 file_name
func TestIllegalComplaintAddValidateFileNameEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.IllegalOrderID = "ORDER_001"
	req.ComplaintReason = "合规"
	req.File = []byte("fake")
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestIllegalComplaintAddValidateAccessTokenEmpty(t *testing.T) {
	req := &model.IllegalComplaintAddReq{}
	req.AccountID = 2045867
	req.IllegalOrderID = "ORDER_001"
	req.ComplaintReason = "合规"
	req.File = []byte("fake")
	req.FileName = "evidence.zip"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
