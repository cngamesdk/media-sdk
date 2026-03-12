package media_sdk

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/model"
	"testing"
)

func TestClientToutiao(t *testing.T) {
	// 创建巨量引擎客户端
	mediaConfig := config.DefaultConfig(config.MediaToutiao)

	client, err := NewClient(mediaConfig)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	req := &model.AccountReq{}
	req.AccessToken = "test"
	req.AdvertiserID = 123
	account, accountErr := client.adapter.GetAccount(ctx, req)
	if accountErr != nil {
		t.Fatal(accountErr)
	}
	println(fmt.Sprintf("%+v", account))
}

func TestClientTencent(t *testing.T) {
	// 创建巨量引擎客户端
	mediaConfig := config.DefaultConfig(config.MediaTencent)

	client, err := NewClient(mediaConfig)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	req := &model.AccountReq{}
	account, accountErr := client.adapter.GetAccount(ctx, req)
	if accountErr != nil {
		t.Fatal(accountErr)
	}
	println(fmt.Sprintf("%+v", account))
}

func TestMultiClient(t *testing.T) {

	mediaConfigToutiao := config.DefaultConfig(config.MediaToutiao)
	clientToutiao, errToutiao := NewClient(mediaConfigToutiao)
	if errToutiao != nil {
		t.Fatal(errToutiao)
	}

	mediaConfigTencent := config.DefaultConfig(config.MediaTencent)
	clientTencent, errTencent := NewClient(mediaConfigTencent)
	if errTencent != nil {
		t.Fatal(errTencent)
	}

	// 创建多媒体管理器
	multiClient := NewMultiClient()
	multiClient.RegisterClient(config.MediaToutiao, clientToutiao)
	multiClient.RegisterClient(config.MediaTencent, clientTencent)

	ctx := context.Background()

	// 获取所有媒体账户信息
	multiErr := multiClient.BatchExecute(ctx, func(client *Client) error {
		account, err := client.GetAccount(ctx, &model.AccountReq{
			AdvertiserID: 123,
		})
		if err != nil {
			return err
		}
		fmt.Printf("账户: %+v\n", account)
		return nil
	})
	if multiErr != nil {
		t.Fatal(multiErr)
	}
}

func TestAccessToken(t *testing.T) {
	// 创建巨量引擎客户端
	mediaConfig := config.DefaultConfig(config.MediaToutiao)
	mediaConfig.AppID = "123"
	mediaConfig.AppSecret = "test"

	client, err := NewClient(mediaConfig)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	req := &model.AccessTokenReq{}
	req.AuthCode = "123"
	account, accountErr := client.adapter.AccessToken(ctx, req)
	if accountErr != nil {
		t.Fatal(accountErr)
	}
	println(fmt.Sprintf("%+v", account))
}
