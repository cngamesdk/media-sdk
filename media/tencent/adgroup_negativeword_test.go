package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 新增广告否定词测试用例 ==========

// TestAdgroupNegativewordAddBasicSelf 测试新增短语否定词（最简参数）
func TestAdgroupNegativewordAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"短语否词1", "短语否词2"}
	req.ExactNegativeWords = []string{}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordAddExactSelf 测试新增精确否定词
func TestAdgroupNegativewordAddExactSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	req.ExactNegativeWords = []string{"精确否词1", "精确否词2"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordAddBothSelf 测试同时新增短语和精确否定词
func TestAdgroupNegativewordAddBothSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"短语否词1", "短语否词2", "短语否词3"}
	req.ExactNegativeWords = []string{"精确否词1", "精确否词2"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordAddEnglishSelf 测试新增英文否定词
func TestAdgroupNegativewordAddEnglishSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 789
	req.PhraseNegativeWords = []string{"free game", "download now"}
	req.ExactNegativeWords = []string{"exact match keyword"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordAddMaxWordSelf 测试新增长度接近上限的否定词（20等宽字符/40英文）
func TestAdgroupNegativewordAddMaxWordSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	// 20个中文字符 = 60字节 < 150字节限制
	req.PhraseNegativeWords = []string{"一二三四五六七八九十一二三四五六七八九十"}
	req.ExactNegativeWords = []string{}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestAdgroupNegativewordAddValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupNegativewordAddValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"否定词1"}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateMissingAdgroupIDSelf 测试缺少adgroup_id
func TestAdgroupNegativewordAddValidateMissingAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.PhraseNegativeWords = []string{"否定词1"}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateBothEmptySelf 测试短语词和精确词同时为空
func TestAdgroupNegativewordAddValidateBothEmptySelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：两个词列表不能同时为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateNilWordsSelf 测试未设置否定词字段
func TestAdgroupNegativewordAddValidateNilWordsSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	// PhraseNegativeWords 和 ExactNegativeWords 均为 nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：两个词列表不能同时为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateExceedPhraseCountSelf 测试短语词数组超过900
func TestAdgroupNegativewordAddValidateExceedPhraseCountSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	words := make([]string, 901)
	for i := range words {
		words[i] = "词"
	}
	req.PhraseNegativeWords = words
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：phrase_negative_words超过900")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateExceedExactCountSelf 测试精确词数组超过900
func TestAdgroupNegativewordAddValidateExceedExactCountSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	words := make([]string, 901)
	for i := range words {
		words[i] = "词"
	}
	req.ExactNegativeWords = words
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：exact_negative_words超过900")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidatePhraseWordTooLongSelf 测试短语词单词超过150字节
func TestAdgroupNegativewordAddValidatePhraseWordTooLongSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	// 51个中文字符 = 153字节 > 150字节限制
	req.PhraseNegativeWords = []string{"一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一"}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：短语词超过150字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordAddValidateExactWordTooLongSelf 测试精确词单词超过150字节
func TestAdgroupNegativewordAddValidateExactWordTooLongSelf(t *testing.T) {
	req := &model.AdgroupNegativewordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	// 51个中文字符 = 153字节 > 150字节限制
	req.ExactNegativeWords = []string{"一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：精确词超过150字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 更新广告否定词测试用例 ==========

// TestAdgroupNegativewordUpdateBasicSelf 测试更新短语否定词（最简参数）
func TestAdgroupNegativewordUpdateBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"短语否词1", "短语否词4"}
	req.ExactNegativeWords = []string{}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordUpdateExactSelf 测试更新精确否定词
func TestAdgroupNegativewordUpdateExactSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	req.ExactNegativeWords = []string{"精确否词1", "精确否词4"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordUpdateBothSelf 测试同时更新短语和精确否定词
func TestAdgroupNegativewordUpdateBothSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"短语否词1", "短语否词2", "短语否词3"}
	req.ExactNegativeWords = []string{"精确否词1", "精确否词2"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordUpdateEnglishSelf 测试更新英文否定词
func TestAdgroupNegativewordUpdateEnglishSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 789
	req.PhraseNegativeWords = []string{"free game", "download now"}
	req.ExactNegativeWords = []string{"exact match keyword"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 更新广告否定词验证测试用例 ==========

// TestAdgroupNegativewordUpdateValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupNegativewordUpdateValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{"否定词1"}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordUpdateValidateMissingAdgroupIDSelf 测试缺少adgroup_id
func TestAdgroupNegativewordUpdateValidateMissingAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.PhraseNegativeWords = []string{"否定词1"}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordUpdateValidateBothEmptySelf 测试短语词和精确词同时为空
func TestAdgroupNegativewordUpdateValidateBothEmptySelf(t *testing.T) {
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：两个词列表不能同时为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordUpdateValidateExceedCountSelf 测试短语词数组超过900
func TestAdgroupNegativewordUpdateValidateExceedCountSelf(t *testing.T) {
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	words := make([]string, 901)
	for i := range words {
		words[i] = "词"
	}
	req.PhraseNegativeWords = words
	req.ExactNegativeWords = []string{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：phrase_negative_words超过900")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordUpdateValidateWordTooLongSelf 测试精确词单词超过150字节
func TestAdgroupNegativewordUpdateValidateWordTooLongSelf(t *testing.T) {
	req := &model.AdgroupNegativewordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 456
	req.PhraseNegativeWords = []string{}
	// 51个中文字符 = 153字节 > 150字节限制
	req.ExactNegativeWords = []string{"一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：精确词超过150字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 查询广告否定词测试用例 ==========

// TestAdgroupNegativewordGetSingleSelf 测试查询单个广告的否定词
func TestAdgroupNegativewordGetSingleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupIDs = []int64{5076023598}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordGetMultipleSelf 测试查询多个广告的否定词
func TestAdgroupNegativewordGetMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupIDs = []int64{111, 222, 333}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupNegativewordGetMaxCountSelf 测试查询100个广告（上限）
func TestAdgroupNegativewordGetMaxCountSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	ids := make([]int64, 100)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.AdgroupIDs = ids
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupNegativewordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 查询广告否定词验证测试用例 ==========

// TestAdgroupNegativewordGetValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupNegativewordGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AdgroupIDs = []int64{123}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordGetValidateMissingAdgroupIDsSelf 测试缺少adgroup_ids
func TestAdgroupNegativewordGetValidateMissingAdgroupIDsSelf(t *testing.T) {
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_ids为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordGetValidateEmptyAdgroupIDsSelf 测试adgroup_ids为空数组
func TestAdgroupNegativewordGetValidateEmptyAdgroupIDsSelf(t *testing.T) {
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupIDs = []int64{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_ids至少包含1个广告id")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupNegativewordGetValidateExceedCountSelf 测试adgroup_ids超过100
func TestAdgroupNegativewordGetValidateExceedCountSelf(t *testing.T) {
	req := &model.AdgroupNegativewordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	ids := make([]int64, 101)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.AdgroupIDs = ids
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_ids超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}
