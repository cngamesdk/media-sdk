package adapter

import (
	"context"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/model"
)

// MediaSDK 媒体SDK统一接口
type MediaSDK interface {
	Code() config.MediaType
	Name() string

	Auth(req *model.AuthReq) (resp interface{}, err error)
	// 账户管理
	GetAccount(ctx context.Context, req *model.AccountReq) (*model.AccountResp, error)
	RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (*model.RefreshTokenResp, error)

	// 广告计划
	CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error)
	UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error)
	GetCampaign(ctx context.Context, req *model.GetCampaignReq) (*model.GetCampaignResp, error)
	ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error)

	// 广告组
	CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error)
	UpdateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error)
	GetUnit(ctx context.Context, req *model.GetUnitReq) (*model.UnitResp, error)
	ListUnits(ctx context.Context, req *model.ListUnitsReq) (*model.ListUnitsResp, error)

	// 广告创意
	CreateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error)
	UpdateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error)
	GetCreative(ctx context.Context, req *model.GetCreativeReq) (*model.CreativeResp, error)
	ListCreatives(ctx context.Context, req *model.ListCreativesReq) (*model.ListCreativesResp, error)

	// 数据报表
	GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error)

	// 批量操作
	//BatchCreate(ctx context.Context, req *model.BatchCreateReq) (*model.BatchCreateResp, error)
	//BatchUpdate(ctx context.Context, req *model.BatchUpdateReq) (*model.BatchUpdateResp, error)
}

// Factory SDK工厂接口
type Factory interface {
	Create(config *config.Config) (MediaSDK, error)
}
