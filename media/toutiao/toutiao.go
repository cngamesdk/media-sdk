package toutiao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/cngamesdk/media-sdk/utils"
	"net/http"
	"time"
)

func init() {
	adapter.Register(config.MediaToutiao, &ToutiaoFactory{})
}

// ToutiaoFactory 巨量引擎工厂
type ToutiaoFactory struct{}

func (f *ToutiaoFactory) Create(config *config.Config) (adapter.MediaSDK, error) {
	client := utils.NewHTTPClient(&utils.HTTPConfig{
		Timeout:    config.Timeout,
		Proxy:      config.Proxy,
		MaxRetries: config.MaxRetries,
		RetryWait:  config.RetryWait,
		Debug:      config.Debug,
	})

	return &ToutiaoAdapter{media.Media{Config: config, Client: client}}, nil
}

// ToutiaoAdapter 巨量引擎适配器
type ToutiaoAdapter struct {
	media.Media
}

func (a *ToutiaoAdapter) Code() config.MediaType {
	return config.MediaToutiao
}

func (a *ToutiaoAdapter) Name() string {
	return "巨量引擎"
}

// GetAccount 获取账户
func (a *ToutiaoAdapter) GetAccount(ctx context.Context, req *model.AccountReq) (*model.AccountResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			ID      string  `json:"id"`
			Name    string  `json:"name"`
			Balance float64 `json:"balance"`
			Status  string  `json:"status"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodGet, "/open_api/2/advertiser/info/", params, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, fmt.Errorf("toutiao api error: code=%d", result.Code)
	}

	return &model.AccountResp{
		ID:      result.Data.ID,
		Name:    result.Data.Name,
		Balance: result.Data.Balance,
		Status:  result.Data.Status,
	}, nil
}

// RefreshToken 刷新Token
func (a *ToutiaoAdapter) RefreshToken(ctx context.Context) error {
	params := map[string]interface{}{
		"app_id":        a.Config.AppID,
		"secret":        a.Config.AppSecret,
		"refresh_token": a.Config.RefreshToken,
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int64  `json:"expires_in"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodPost, "/open_api/2/oauth2/refresh_token/", params, &result)
	if err != nil {
		return err
	}

	a.Config.AccessToken = result.Data.AccessToken
	a.Config.RefreshToken = result.Data.RefreshToken
	a.Config.ExpireTime = time.Now().Add(time.Duration(result.Data.ExpiresIn) * time.Second)

	return nil
}

// CreateCampaign 创建广告计划
func (a *ToutiaoAdapter) CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"name":          req.Name,
		"budget":        req.Budget,
		"budget_mode":   req.BudgetMode,
		"status":        req.Status,
		"start_time":    req.StartTime.Unix(),
	}

	if !req.EndTime.IsZero() {
		params["end_time"] = req.EndTime.Unix()
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			ID string `json:"campaign_id"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodPost, "/open_api/2/campaign/create/", params, &result)
	if err != nil {
		return nil, err
	}

	return &model.CampaignResp{
		ID: result.Data.ID,
	}, nil
}

// UpdateCampaign 更新广告计划
func (a *ToutiaoAdapter) UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"campaign_id":   req.ID,
		"name":          req.Name,
		"budget":        req.Budget,
		"status":        req.Status,
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			ID string `json:"campaign_id"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodPost, "/open_api/2/campaign/update/", params, &result)
	if err != nil {
		return nil, err
	}

	return &model.CampaignResp{
		ID: result.Data.ID,
	}, nil
}

// GetCampaign 获取广告计划
func (a *ToutiaoAdapter) GetCampaign(ctx context.Context, req *model.GetCampaignReq) (*model.GetCampaignResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"campaign_id":   req.CampaignID,
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			ID         string  `json:"campaign_id"`
			Name       string  `json:"name"`
			Budget     float64 `json:"budget"`
			BudgetMode string  `json:"budget_mode"`
			Status     string  `json:"status"`
			StartTime  string  `json:"start_time"`
			EndTime    string  `json:"end_time"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodGet, "/open_api/2/campaign/get/", params, &result)
	if err != nil {
		return nil, err
	}

	return &model.GetCampaignResp{}, nil
}

// ListCampaigns 列出广告计划
func (a *ToutiaoAdapter) ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}

	if len(req.Status) > 0 {
		params["filtering"] = map[string]interface{}{
			"status": req.Status,
		}
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			List     []map[string]interface{} `json:"list"`
			PageInfo struct {
				Page        int `json:"page"`
				PageSize    int `json:"page_size"`
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodGet, "/open_api/2/campaign/list/", params, &result)
	if err != nil {
		return nil, err
	}

	campaigns := make([]*model.CampaignResp, 0, len(result.Data.List))
	for _, item := range result.Data.List {
		campaigns = append(campaigns, &model.CampaignResp{
			ID:   item["campaign_id"].(string),
			Name: item["name"].(string),
		})
	}

	return &model.ListCampaignsResp{
		List:     campaigns,
		Total:    result.Data.PageInfo.TotalNumber,
		Page:     result.Data.PageInfo.Page,
		PageSize: result.Data.PageInfo.PageSize,
	}, nil
}

// CreateUnit 创建广告组
func (a *ToutiaoAdapter) CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"campaign_id":   req.CampaignID,
		"name":          req.Name,
		"pricing":       req.Pricing,
		"bid":           req.BidAmount,
		"budget":        req.DailyBudget,
		"status":        req.Status,
	}

	// 处理定向
	if req.Target != nil {
		params["audience"] = a.buildTargeting(req.Target)
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			ID string `json:"ad_id"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodPost, "/open_api/2/ad/create/", params, &result)
	if err != nil {
		return nil, err
	}

	return &model.UnitResp{
		ID: result.Data.ID,
	}, nil
}

// UpdateUnit 更新广告组
func (a *ToutiaoAdapter) UpdateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// UpdateUnit 获取广告组
func (a *ToutiaoAdapter) GetUnit(ctx context.Context, req *model.GetUnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// ListUnits 获取广告组列表
func (a *ToutiaoAdapter) ListUnits(ctx context.Context, req *model.ListUnitsReq) (*model.ListUnitsResp, error) {
	return nil, nil
}

// CreateCreative 创建广告创意
func (a *ToutiaoAdapter) CreateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// UpdateCreative 更新广告创意
func (a *ToutiaoAdapter) UpdateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// GetCreative 获取广告创意
func (a *ToutiaoAdapter) GetCreative(ctx context.Context, req *model.GetCreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// GetCreative 获取广告创意列表
func (a *ToutiaoAdapter) ListCreatives(ctx context.Context, req *model.ListCreativesReq) (*model.ListCreativesResp, error) {
	return nil, nil
}

// GetReport 获取报表
func (a *ToutiaoAdapter) GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_date":    req.StartDate,
		"end_date":      req.EndDate,
		"group_by":      req.GroupBy,
		"page":          req.Page,
		"page_size":     req.PageSize,
	}

	var result struct {
		Code int `json:"code"`
		Data struct {
			List     []map[string]interface{} `json:"list"`
			PageInfo struct {
				Page        int `json:"page"`
				PageSize    int `json:"page_size"`
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}

	err := a.Request(ctx, http.MethodGet, "/open_api/2/report/ad/get/", params, &result)
	if err != nil {
		return nil, err
	}

	var reports []*model.ReportData
	for _, item := range result.Data.List {
		report := &model.ReportData{
			Impressions: int64(item["show_cnt"].(float64)),
			Clicks:      int64(item["click_cnt"].(float64)),
			Cost:        item["stat_cost"].(float64),
		}
		reports = append(reports, report)
	}

	return &model.ReportResp{
		List:     reports,
		Total:    result.Data.PageInfo.TotalNumber,
		Page:     result.Data.PageInfo.Page,
		PageSize: result.Data.PageInfo.PageSize,
	}, nil
}

// Req 发送请求
func (a *ToutiaoAdapter) Request(ctx context.Context, method, path string, params interface{}, result interface{}) error {
	headers := map[string]string{
		"Access-Token": a.Config.AccessToken,
		"Content-Type": "application/json",
	}

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	resp, err := a.Client.Request(ctx, method, a.Config.BaseURL+path, bytes.NewReader(body), headers)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, result)
}

// buildTargeting 构建定向
func (a *ToutiaoAdapter) buildTargeting(target *model.Targeting) map[string]interface{} {
	audience := make(map[string]interface{})

	if len(target.Gender) > 0 {
		audience["gender"] = target.Gender[0]
	}

	if len(target.Age) > 0 {
		audience["age"] = target.Age
	}

	if len(target.Region) > 0 {
		audience["region"] = target.Region
	}

	if len(target.Interests) > 0 {
		audience["interest"] = target.Interests
	}

	return audience
}
