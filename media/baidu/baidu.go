package baidu

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media"
	model2 "github.com/cngamesdk/media-sdk/media/baidu/model"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/cngamesdk/media-sdk/utils"
)

func init() {
	adapter.Register(config.MediaBaidu, &BaiduFactory{})
}

// BaiduFactory 百度营销工厂
type BaiduFactory struct{}

func (f *BaiduFactory) Create(config *config.Config) (adapter.MediaSDK, error) {
	return NewBaiduAdapter(config), nil
}

// NewBaiduAdapter 初始化百度营销适配器
func NewBaiduAdapter(config *config.Config) *BaiduAdapter {
	client := utils.NewHTTPClient(&utils.HTTPConfig{
		Timeout:    config.Timeout,
		Proxy:      config.Proxy,
		MaxRetries: config.MaxRetries,
		RetryWait:  config.RetryWait,
		Debug:      config.Debug,
	})

	return &BaiduAdapter{media.Media{Config: config, Client: client}}
}

// BaiduAdapter 百度营销适配器
type BaiduAdapter struct {
	media.Media
}

func (a *BaiduAdapter) Code() config.MediaType {
	return config.MediaBaidu
}

func (a *BaiduAdapter) Name() string {
	return "百度营销"
}

// AuthorizationLinkSelf 获取授权链接
// 程序化拼接授权链接方式
// 模板: https://u.baidu.com/oauth/page/index?platformId={platformId}&appId={appId}&scope={scope}&state={state}&callback={callback}
func (a *BaiduAdapter) AuthorizationLinkSelf(ctx context.Context, req *model2.AuthorizationLinkReq) (resp *model2.AuthorizationLinkResp, err error) {
	_ = ctx
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}

	resp = &model2.AuthorizationLinkResp{
		AuthorizationURL: req.BuildURL(),
		PlatformID:       model2.PlatformID,
		AppID:            req.AppID,
		Scope:            req.Scope,
		State:            req.State,
		Callback:         req.Callback,
	}
	return
}

// AccessTokenSelf 换取授权令牌
// POST https://u.baidu.com/oauth/accessToken
func (a *BaiduAdapter) AccessTokenSelf(ctx context.Context, req *model2.AccessTokenReq) (resp *model2.AccessTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.AccessTokenResp
	errRequest := a.RequestPostJson(ctx, nil, model2.BaseUrlOAuthAPI+"/accessToken", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// RefreshTokenSelf 更新授权令牌
// POST https://u.baidu.com/oauth/refreshToken
func (a *BaiduAdapter) RefreshTokenSelf(ctx context.Context, req *model2.RefreshTokenReq) (resp *model2.RefreshTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model2.RefreshTokenResp
	errRequest := a.RequestPostJson(ctx, nil, model2.BaseUrlOAuthAPI+"/refreshToken", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// Auth 授权
func (a *BaiduAdapter) Auth(ctx context.Context, req *model.AuthReq) (resp interface{}, err error) {
	return nil, fmt.Errorf("not implemented")
}

// AccessToken 获取AccessToken（统一接口，建议直接使用 AccessTokenSelf）
func (a *BaiduAdapter) AccessToken(ctx context.Context, req *model.AccessTokenReq) (*model.AccessTokenResp, error) {
	return nil, fmt.Errorf("not implemented, use AccessTokenSelf instead")
}

// RefreshToken 刷新Token（统一接口，建议直接使用 RefreshTokenSelf）
func (a *BaiduAdapter) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (*model.RefreshTokenResp, error) {
	return nil, fmt.Errorf("not implemented, use RefreshTokenSelf instead")
}

// GetAccount 获取账户信息
func (a *BaiduAdapter) GetAccount(ctx context.Context, req *model.AccountReq) (*model.AccountResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// CreateCampaign 创建广告计划
func (a *BaiduAdapter) CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// UpdateCampaign 更新广告计划
func (a *BaiduAdapter) UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListCampaigns 列出广告计划
func (a *BaiduAdapter) ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// CreateUnit 创建广告组
func (a *BaiduAdapter) CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// UpdateUnit 更新广告组
func (a *BaiduAdapter) UpdateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListUnits 列出广告组
func (a *BaiduAdapter) ListUnits(ctx context.Context, req *model.ListUnitsReq) (*model.ListUnitsResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// CreateCreative 创建广告创意
func (a *BaiduAdapter) CreateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// UpdateCreative 更新广告创意
func (a *BaiduAdapter) UpdateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListCreatives 列出广告创意
func (a *BaiduAdapter) ListCreatives(ctx context.Context, req *model.ListCreativesReq) (*model.ListCreativesResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetReport 获取报表
func (a *BaiduAdapter) GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error) {
	return nil, fmt.Errorf("not implemented")
}

// RequestGet 发送GET请求（带dealResponse处理）
func (a *BaiduAdapter) RequestGet(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model2.BaseResp
	if err = a.Media.RequestGet(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

// RequestPostJson 发送POST JSON请求（带dealResponse处理）
func (a *BaiduAdapter) RequestPostJson(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model2.BaseResp
	if err = a.Media.RequestPostJson(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *BaiduAdapter) dealResponse(req model2.BaseResp, result interface{}) (err error) {
	if req.Code != 0 {
		err = fmt.Errorf("baidu api error: code=%d, message:%s", req.Code, req.Message)
		return
	}
	dataJson, dataJsonErr := json.Marshal(req.Data)
	if dataJsonErr != nil {
		err = fmt.Errorf("baidu response to json error: %s", dataJsonErr.Error())
		return
	}
	if unJsonErr := json.Unmarshal(dataJson, result); unJsonErr != nil {
		err = fmt.Errorf("baidu json to target error: %s", unJsonErr.Error())
		return
	}
	return
}

// RequestPostJsonBusiness 发送业务API POST JSON请求
// 自动将body包装为 {header: {userName, accessToken}, body: {...}} 格式
// 自动解析外层header并检查status，将body数据填充到result
func (a *BaiduAdapter) RequestPostJsonBusiness(ctx context.Context, userName string, accessToken string, url string, body interface{}, result interface{}) (err error) {
	req := &model2.ApiReq{
		Header: model2.ApiReqHeader{
			UserName:    userName,
			AccessToken: accessToken,
		},
		Body: body,
	}
	var resp model2.ApiResp
	if err = a.Media.RequestPostJson(ctx, nil, url, req, &resp); err != nil {
		return
	}
	err = a.dealBusinessResponse(resp, result)
	return
}

func (a *BaiduAdapter) dealBusinessResponse(resp model2.ApiResp, result interface{}) (err error) {
	if resp.Header.Status != 0 {
		var failureMsgs []string
		for _, f := range resp.Header.Failures {
			failureMsgs = append(failureMsgs, fmt.Sprintf("code=%d, message=%s, position=%s", f.Code, f.Message, f.Position))
		}
		err = fmt.Errorf("baidu api error: status=%d, desc=%s, failures=[%v]", resp.Header.Status, resp.Header.Desc, failureMsgs)
		return
	}
	bodyJson, bodyJsonErr := json.Marshal(resp.Body)
	if bodyJsonErr != nil {
		err = fmt.Errorf("baidu response body to json error: %s", bodyJsonErr.Error())
		return
	}
	if unJsonErr := json.Unmarshal(bodyJson, result); unJsonErr != nil {
		err = fmt.Errorf("baidu json to target error: %s", unJsonErr.Error())
		return
	}
	return
}
