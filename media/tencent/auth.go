package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// OAuth2AuthorizeSelf 授权
func (a *TencentAdapter) OAuth2AuthorizeSelf(ctx context.Context, req *model.OAuth2AuthorizeReq) (resp *model.OAuth2AuthorizeResp, err error) {
	_ = ctx
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	convertResult, convertErr := utils.ConvertStructToQueryString(req)
	if convertErr != nil {
		err = convertErr
		return
	}
	resp = &model.OAuth2AuthorizeResp{}
	resp.RedirectURL = model.DevelopersUrl + "/oauth/authorize?" + convertResult
	return
}

// OAuth2TokenSelf 授权码获取token
func (a *TencentAdapter) OAuth2TokenSelf(ctx context.Context, req *model.OAuth2TokenReq) (resp *model.OAuth2TokenResp, err error) {
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	convertResult, convertErr := utils.ConvertStructToQueryString(req)
	if convertErr != nil {
		err = convertErr
		return
	}
	var result model.OAuth2TokenResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl+"/oauth/token", convertResult, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}

// RefreshTokenSelf 刷新Refresh Token
func (a *TencentAdapter) RefreshTokenSelf(ctx context.Context, req *model.RefreshTokenReq) (resp *model.RefreshTokenResp, err error) {
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	convertResult, convertErr := utils.ConvertStructToQueryString(req)
	if convertErr != nil {
		err = convertErr
		return
	}
	var result model.RefreshTokenResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl+"/oauth/refresh_token", convertResult, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
