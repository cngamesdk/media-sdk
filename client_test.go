package media_sdk

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/model"
	"testing"
)

func TestClientToutiaoAuth(t *testing.T) {
	// 创建巨量引擎客户端
	client, err := NewClientDefault(config.MediaToutiao)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	req := &model.AuthReq{}
	req.AppId = 123
	req.RedirectUri = "https://www.test.com"
	account, accountErr := client.adapter.Auth(ctx, req)
	if accountErr != nil {
		t.Fatal(accountErr)
	}
	println(fmt.Sprintf("%+v", account))
}

func TestClientToutiao(t *testing.T) {
	client, err := NewClientDefault(config.MediaToutiao)
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
	client, err := NewClientDefault(config.MediaTencent)
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

	clientToutiao, errToutiao := NewClientDefault(config.MediaToutiao)
	if errToutiao != nil {
		t.Fatal(errToutiao)
	}

	clientTencent, errTencent := NewClientDefault(config.MediaTencent)
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

func TestAccessTokenToutiao(t *testing.T) {
	client, err := NewClientDefault(config.MediaToutiao)
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
