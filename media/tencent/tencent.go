package tencent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media"
	model3 "github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/cngamesdk/media-sdk/utils"
	"mime/multipart"
	"net/http"
)

func init() {
	adapter.Register(config.MediaTencent, &TencentFactory{})
}

// TencentFactory 腾讯广告工厂
type TencentFactory struct{}

func (f *TencentFactory) Create(config *config.Config) (adapter.MediaSDK, error) {
	return NewTencentAdapter(config), nil
}

// NewTencentAdapter 初始化巨量引擎适配器
func NewTencentAdapter(config *config.Config) *TencentAdapter {
	client := utils.NewHTTPClient(&utils.HTTPConfig{
		Timeout:    config.Timeout,
		Proxy:      config.Proxy,
		MaxRetries: config.MaxRetries,
		RetryWait:  config.RetryWait,
		Debug:      config.Debug,
	})

	return &TencentAdapter{media.Media{Config: config, Client: client}}
}

// TencentAdapter 腾讯广告适配器
type TencentAdapter struct {
	media.Media
}

func (a *TencentAdapter) Code() config.MediaType {
	return config.MediaTencent
}

func (a *TencentAdapter) Name() string {
	return "腾讯广告"
}

// Auth 授权
func (a *TencentAdapter) Auth(ctx context.Context, req *model.AuthReq) (resp interface{}, err error) {
	_ = ctx
	myReq := &model3.OAuth2AuthorizeReq{}
	if convertErr := myReq.Convert(req); convertErr != nil {
		err = convertErr
		return
	}
	result, resultErr := a.OAuth2AuthorizeSelf(ctx, myReq)
	if resultErr != nil {
		err = resultErr
		return
	}
	resp = result
	return
}

// Auth 授权
func (a *TencentAdapter) AccessToken(ctx context.Context, req *model.AccessTokenReq) (resp *model.AccessTokenResp, err error) {
	myReq := model3.AccessTokenReq{}
	myReq.Convert(req)
	myReq.Format()
	if validateErr := myReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model3.AccessTokenResp
	errRequest := a.RequestPostJson(ctx, nil, model3.ApiUrl+"/oauth/token", myReq, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp, err = result.Convert()
	return
}

// RefreshToken 刷新Token
func (a *TencentAdapter) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (resp *model.RefreshTokenResp, err error) {
	return
}

// GetAccount 获取账户
func (a *TencentAdapter) GetAccount(ctx context.Context, req *model.AccountReq) (resp *model.AccountResp, err error) {
	return
}

func (a *TencentAdapter) RequestGet(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model3.BaseResp
	if err = a.Media.RequestGet(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *TencentAdapter) RequestPostMultipart(ctx context.Context, url string, fields map[string]string, fileField, fileName string, fileData []byte, result interface{}) (err error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for k, v := range fields {
		if writeErr := w.WriteField(k, v); writeErr != nil {
			return writeErr
		}
	}
	fw, createErr := w.CreateFormFile(fileField, fileName)
	if createErr != nil {
		return createErr
	}
	if _, writeErr := fw.Write(fileData); writeErr != nil {
		return writeErr
	}
	if closeErr := w.Close(); closeErr != nil {
		return closeErr
	}
	headers := map[string]string{"Content-Type": w.FormDataContentType()}
	resp, reqErr := a.Client.Request(ctx, http.MethodPost, url, &buf, headers)
	if reqErr != nil {
		return reqErr
	}
	var response model3.BaseResp
	if unmarshalErr := json.Unmarshal(resp, &response); unmarshalErr != nil {
		return unmarshalErr
	}
	return a.dealResponse(response, result)
}

func (a *TencentAdapter) RequestPostJson(ctx context.Context, headers map[string]string, url string, data interface{}, result interface{}) (err error) {
	var response model3.BaseResp
	if err = a.Media.RequestPostJson(ctx, headers, url, data, &response); err != nil {
		return
	}
	err = a.dealResponse(response, result)
	return
}

func (a *TencentAdapter) dealResponse(req model3.BaseResp, result interface{}) (err error) {
	if req.Code != 0 {
		err = fmt.Errorf("api error: code=%d, message:%s, message_cn:%s, request_id:%s", req.Code, req.Message, req.MessageCn, req.RequestId)
		return
	}
	dataJson, dataJsonErr := json.Marshal(req.Data)
	if dataJsonErr != nil {
		err = fmt.Errorf("response to json error: %s", dataJsonErr.Error())
		return
	}
	if unJsonErr := json.Unmarshal(dataJson, result); unJsonErr != nil {
		err = fmt.Errorf("json to target error: %s", unJsonErr.Error())
		return
	}
	return
}

// CreateCampaign 创建广告计划
func (a *TencentAdapter) CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, nil
}

// UpdateCampaign 更新广告计划
func (a *TencentAdapter) UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	return nil, nil
}

// ListCampaigns 列出广告计划
func (a *TencentAdapter) ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error) {
	return nil, nil
}

// CreateUnit 创建广告组
func (a *TencentAdapter) CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// UpdateUnit 更新广告组
func (a *TencentAdapter) UpdateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	return nil, nil
}

// ListUnits 获取广告组列表
func (a *TencentAdapter) ListUnits(ctx context.Context, req *model.ListUnitsReq) (*model.ListUnitsResp, error) {
	return nil, nil
}

// CreateCreative 创建广告创意
func (a *TencentAdapter) CreateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// UpdateCreative 更新广告创意
func (a *TencentAdapter) UpdateCreative(ctx context.Context, req *model.CreativeReq) (*model.CreativeResp, error) {
	return nil, nil
}

// ListCreatives 获取广告创意列表
func (a *TencentAdapter) ListCreatives(ctx context.Context, req *model.ListCreativesReq) (*model.ListCreativesResp, error) {
	return nil, nil
}

// GetReport 获取报表
func (a *TencentAdapter) GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error) {
	return nil, nil
}
