package kuaishou

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media"
	model3 "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/cngamesdk/media-sdk/utils"
)

func init() {
	adapter.Register(config.MediaKuaiShou, &KuaishouFactory{})
}

// KuaishouFactory 腾讯广告工厂
type KuaishouFactory struct{}

func (f *KuaishouFactory) Create(config *config.Config) (adapter.MediaSDK, error) {
	return NewKuaishouAdapter(config), nil
}

// NewKuaishouAdapter 初始化巨量引擎适配器
func NewKuaishouAdapter(config *config.Config) *KuaishouAdapter {
	client := utils.NewHTTPClient(&utils.HTTPConfig{
		Timeout:    config.Timeout,
		Proxy:      config.Proxy,
		MaxRetries: config.MaxRetries,
		RetryWait:  config.RetryWait,
		Debug:      config.Debug,
	})

	return &KuaishouAdapter{media.Media{Config: config, Client: client}}
}

// KuaishouAdapter 腾讯广告适配器
type KuaishouAdapter struct {
	media.Media
}

func (a *KuaishouAdapter) Code() config.MediaType {
	return config.MediaKuaiShou
}

func (a *KuaishouAdapter) Name() string {
	return "磁力引擎"
}

// Auth 授权
func (a *KuaishouAdapter) Auth(ctx context.Context, req *model.AuthReq) (resp interface{}, err error) {
	_ = ctx
	myReq := &model3.AuthReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	convertResult, convertErr := utils.ConvertStructToQueryString(myReq)
	if convertErr != nil {
		err = convertErr
		return
	}
	authResp := model3.AuthResp(model3.DevelopersUrl + "/tools/authorize?" + convertResult)
	resp = authResp
	return
}

// Auth 授权
func (a *KuaishouAdapter) AccessToken(ctx context.Context, req *model.AccessTokenReq) (resp *model.AccessTokenResp, err error) {
	myReq := model3.AccessTokenReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model3.AccessTokenResp
	errRequest := a.RequestPostJson(ctx, nil, model3.AdUrl+"/oauth/token", myReq, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp, err = result.Convert()
	return
}

// RefreshToken 刷新Token
func (a *KuaishouAdapter) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (resp *model.RefreshTokenResp, err error) {
	myReq := &model3.RefreshTokenReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model3.RefreshTokenResp

	requestErr := a.RequestPostJson(ctx, nil, model3.AdUrl+"/oauth/refresh_token", myReq, &result)
	if requestErr != nil {
		err = requestErr
		return
	}
	resp, err = result.Convert()
	return
}

// GetAccount 获取账户
func (a *KuaishouAdapter) GetAccount(ctx context.Context, req *model.AccountReq) (resp *model.AccountResp, err error) {
	return
}

func (a *KuaishouAdapter) RequestGet(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model3.BaseResp
	if err = a.Media.RequestGet(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *KuaishouAdapter) RequestPostJson(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model3.BaseResp
	if err = a.Media.RequestPostJson(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *KuaishouAdapter) dealResponse(req model3.BaseResp, result interface{}) (err error) {
	if req.Code != 0 {
		err = fmt.Errorf("toutiao api error: code=%d, message:%s, request_id:%s", req.Code, req.Message, req.RequestId)
		return
	}
	dataJson, dataJsonErr := json.Marshal(req.Data)
	if dataJsonErr != nil {
		err = fmt.Errorf("toutiao response to json error: %s", dataJsonErr.Error())
		return
	}
	if unJsonErr := json.Unmarshal(dataJson, result); unJsonErr != nil {
		err = fmt.Errorf("toutiao json to target error: %s", unJsonErr.Error())
		return
	}
	return
}

// CreateCampaign 创建广告计划
func (a *KuaishouAdapter) CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, nil
}

// UpdateCampaign 更新广告计划
func (a *KuaishouAdapter) UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, nil
}

// ListCampaigns 列出广告计划
func (a *KuaishouAdapter) ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error) {
	return nil, nil
}

// CreateUnit 创建广告组
func (a *KuaishouAdapter) CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// UpdateUnit 更新广告组
func (a *KuaishouAdapter) UpdateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// ListUnits 获取广告组列表
func (a *KuaishouAdapter) ListUnits(ctx context.Context, req *model.ListUnitsReq) (*model.ListUnitsResp, error) {
	return nil, nil
}

// CreateCreative 创建广告创意
func (a *KuaishouAdapter) CreateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// UpdateCreative 更新广告创意
func (a *KuaishouAdapter) UpdateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// ListCreatives 获取广告创意列表
func (a *KuaishouAdapter) ListCreatives(ctx context.Context, req *model.ListCreativesReq) (*model.ListCreativesResp, error) {
	return nil, nil
}

// GetReport 获取报表
func (a *KuaishouAdapter) GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error) {
	return nil, nil
}
