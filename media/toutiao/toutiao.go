package toutiao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/cngamesdk/media-sdk/utils"
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

// Auth 授权
func (a *ToutiaoAdapter) Auth(ctx context.Context, req *model.AuthReq) (resp interface{}, err error) {
	_ = ctx
	myReq := &model2.AuthReq{}
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
	authResp := model2.AuthResp(model2.BaseUrlOpen + "/audit/oauth.html?" + convertResult)
	resp = authResp
	return
}

// Auth 授权
func (a *ToutiaoAdapter) AccessToken(ctx context.Context, req *model.AccessTokenReq) (resp *model.AccessTokenResp, err error) {
	myReq := model2.AccessTokenReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AccessTokenResp
	errRequest := a.RequestPostJson(ctx, nil, model2.BaseUrlApi+"/open_api/oauth2/access_token/", myReq, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp, err = result.Convert()
	return
}

// GetAccount 获取账户
func (a *ToutiaoAdapter) GetAccount(ctx context.Context, req *model.AccountReq) (resp *model.AccountResp, err error) {
	myReq := &model2.AccountReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AccountResp
	errRequest := a.RequestGet(ctx, myReq.Headers, model2.BaseUrlAd+"/open_api/2/advertiser/info/", myReq, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp, err = result.Convert()
	return
}

func (a *ToutiaoAdapter) RequestGet(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model2.BaseResp
	if err = a.Media.RequestGet(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *ToutiaoAdapter) RequestPostJson(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model2.BaseResp
	if err = a.Media.RequestPostJson(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *ToutiaoAdapter) dealResponse(req model2.BaseResp, result interface{}) (err error) {
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

// RefreshToken 刷新Token
func (a *ToutiaoAdapter) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (resp *model.RefreshTokenResp, err error) {
	myReq := &model2.RefreshTokenReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.RefreshTokenResp

	requestErr := a.RequestPostJson(ctx, nil, model2.BaseUrlApi+"/open_api/2/oauth2/refresh_token/", myReq, &result)
	if requestErr != nil {
		err = requestErr
		return
	}
	resp, err = result.Convert()
	return
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

	err := a.RequestPostJson(ctx, nil, "/open_api/2/campaign/create/", params, &result)
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

	err := a.RequestPostJson(ctx, nil, "/open_api/2/campaign/update/", params, &result)
	if err != nil {
		return nil, err
	}

	return &model.CampaignResp{
		ID: result.Data.ID,
	}, nil
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

	err := a.RequestGet(ctx, nil, model2.BaseUrlApi+"/open_api/v3.0/project/list/", params, &result)
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

	err := a.RequestPostJson(ctx, nil, model2.BaseUrlApi+"/open_api/2/ad/create/", params, &result)
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

	err := a.RequestGet(ctx, nil, "/open_api/2/report/ad/get/", params, &result)
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
