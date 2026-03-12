# 广告渠道SDK

## 简介
这是一个统一的广告渠道SDK，支持主流广告平台的快速对接，包括巨量引擎、腾讯广告、磁力引擎（快手）、百度、UC、爱奇艺、哔哩哔哩等。

## 特性
- 统一接口：所有媒体使用相同的API接口
- 一键调用：通过配置文件切换媒体
- 扩展性强：支持自定义适配器
- 性能稳定：支持限流、重试、超时控制
- 文档清晰：完整的代码注释和示例

## 安装

```bash
go get -u github.com/cngamesdk/media-sdk
```

## 快速开始

### 1. 创建配置

```Go
import (
    "github.com/cngamesdk/media-sdk"
    "time"
)

config := adsdk.DefaultConfig(adsdk.MediaToutiao)
config.AppID = "your_app_id"
config.AppSecret = "your_app_secret"
```

### 2. 创建客户端

```Go
client, err := media-sdk.NewClient(config)
if err != nil {
    panic(err)
}
```

### 3. 调用API

```Go
// 获取账户信息
account, err := client.GetAccount(ctx, &model.AccountReq{
    AccessToken: "your_access_token",
    AdvertiserID: your_advertiser_id,
})

// 创建广告计划
campaign, err := client.CreateCampaign(ctx, &model.CampaignReq{
    AdvertiserID: "your_advertiser_id",
    Name:         "测试计划",
    Budget:       1000,
    BudgetMode:   "DAY",
    Status:       "ENABLE",
})
```

### 4. 多媒体管理

```go
// 创建多客户端管理器
multiClient := media-sdk.NewMultiClient()

// 注册多个媒体客户端
multiClient.RegisterClient(config.MediaToutiao, toutiaoClient)
multiClient.RegisterClient(config.MediaTencent, tencentClient)

// 获取指定媒体客户端
client, _ := multiClient.GetClient(adsdk.MediaToutiao)

// 批量执行
multiClient.BatchExecute(ctx, func(client *media-sdk.Client) error {
    _, err := client.GetAccount(ctx, req)
    return err
})
```

## 配置说明

| 参数        | 说明         | 默认值  |
|------------|--------------|--------|
| MediaType  | 媒体类型       | -      |
| AppID      | 应用ID        | -      |
| AppSecret  | 应用密钥       | -      |
| Timeout    | 超时时间       | 30s    |
| RateLimit  | QPS限制       | 10     |
| MaxRetries | 最大重试次数   | 3      |
| RetryWait  | 重试等待时间   | 1s     |
| Debug      | 调试模式       | false  |

## 支持的媒体

| 媒体     | MediaType | 文档 |
|---------|-----------|------|
| 巨量引擎 | toutiao   | [文档](https://open.oceanengine.com/) |
| 腾讯广告 | tencent   | [文档](https://developers.e.qq.com/) |
| 磁力引擎 | kuaisou   | [文档](https://developers.e.kuaishou.com/welcome) |
| 百度    | baidu     | [文档](https://dev2.baidu.com/home) |
| UC      | uc        | [文档](https://e.uc.cn/) |
| 爱奇艺   | iqiyi     | [文档](https://api.iqiyi.com/) |
| 哔哩哔哩 | bilibili  | [文档](https://api.bilibili.com/) |

## 扩展开发

### 添加新的媒体

1. 在media目录下创建新的适配器文件
2. 实现MediaSDK接口
3. 在init函数中注册适配器

```go
package adapter

import (
    "media-sdk"
)

func init() {
    Register(media-sdk.MediaNew, &NewMediaFactory{})
}

type NewMediaFactory struct{}

func (f *NewMediaFactory) Create(config *media-sdk.Config) (media-sdk.MediaSdk, error) {
    return &NewMediaAdapter{config: config}, nil
}

type NewMediaAdapter struct {
    config *media-sdk.Config
}

// 实现所有接口方法...
```

## 性能优化

 - 连接池复用
 - 限流控制
 - 自动重试
 - 并发安全
 - 缓存支持

## 错误处理

```go
if err != nil {
    if errors.Is(err, media-sdk.ErrTokenExpired) {
        // 刷新token
        client.RefreshToken(ctx)
    } else if errors.Is(err, media-sdk.ErrRateLimit) {
        // 等待后重试
        time.Sleep(time.Second)
    }
}
```

## 注意事项

 1. Token管理：注意token过期时间，及时刷新
 2. 限流控制：遵守各媒体的QPS限制
 3. 并发安全：所有接口都是并发安全的
 4. 错误处理：建议实现完整的错误处理逻辑

## License

 MIT

## 使用示例

### main.go

```go
package main

import (
    "media-sdk"
    "media-sdk/model"
    "context"
    "fmt"
    "log"
    "time"
)

func main() {
    // 创建巨量引擎客户端
    toutiaoConfig := media-sdk.DefaultConfig(model.MediaToutiao)
    toutiaoConfig.AppID = "your_app_id"
    toutiaoConfig.AppSecret = "your_app_secret"
    
    toutiaoClient, err := media-sdk.NewClient(toutiaoConfig)
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建腾讯广告客户端
    tencentConfig := media-sdk.DefaultConfig(model.MediaTencent)
    tencentConfig.AppID = "your_app_id"
    tencentConfig.AppSecret = "your_app_secret"
    
    tencentClient, err := media-sdk.NewClient(tencentConfig)
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建多媒体管理器
    multiClient := media-sdk.NewMultiClient()
    multiClient.RegisterClient(media-sdk.MediaToutiao, toutiaoClient)
    multiClient.RegisterClient(media-sdk.MediaTencent, tencentClient)
    
    ctx := context.Background()
    
    // 获取所有媒体账户信息
    multiClient.BatchExecute(ctx, func(client *media-sdk.Client) error {
        account, err := client.GetAccount(ctx, &model.AccountRequest{
            AdvertiserID: client.config.AdvertiserID,
        })
        if err != nil {
            return err
        }
        fmt.Printf("账户: %+v\n", account)
        return nil
    })
    
    // 创建广告计划
    campaign, err := toutiaoClient.CreateCampaign(ctx, &model.CampaignRequest{
        AdvertiserID: toutiaoConfig.AdvertiserID,
        Name:         "测试计划_" + time.Now().Format("20060102150405"),
        Budget:       1000,
        BudgetMode:   "DAY",
        Status:       "ENABLE",
        StartTime:    time.Now(),
        EndTime:      time.Now().AddDate(0, 1, 0),
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("创建计划成功: %s\n", campaign.ID)
    
    // 创建广告组
    unit, err := toutiaoClient.CreateUnit(ctx, &model.UnitRequest{
        CampaignID:   campaign.ID,
        AdvertiserID: toutiaoConfig.AdvertiserID,
        Name:         "测试组",
        Pricing:      "OCPM",
        BidAmount:    10.0,
        DailyBudget:  100.0,
        Status:       "ENABLE",
        Target: &types.Targeting{
            Gender: []string{"MALE", "FEMALE"},
            Age:    []string{"18-24", "25-35"},
            Region: []string{"110000", "310000"},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("创建组成功: %s\n", unit.ID)
    
    // 获取报表
    report, err := toutiaoClient.GetReport(ctx, &model.ReportRequest{
        AdvertiserID: toutiaoConfig.AdvertiserID,
        StartDate:    time.Now().AddDate(0, 0, -7).Format("2006-01-02"),
        EndDate:      time.Now().Format("2006-01-02"),
        Level:        "CAMPAIGN",
        Page:         1,
        PageSize:     10,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("报表数据: %+v\n", report)
}
```

这个SDK封装提供了：

统一接口：所有媒体使用相同的API接口

一键调用：通过配置文件轻松切换媒体

扩展性强：支持添加新的媒体适配器

性能稳定：支持限流、重试、超时控制

文档清晰：完整的代码注释和使用示例

可以直接使用这些代码搭建投放管理系统。