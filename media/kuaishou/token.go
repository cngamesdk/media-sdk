package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AuthSelf 生成授权URL
func (a *KuaishouAdapter) AuthSelf(ctx context.Context, req *kuaishouModel.AuthReq) (resp *kuaishouModel.AuthResp, err error) {
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
	result := kuaishouModel.AuthResp(kuaishouModel.DevelopersUrl + "/tools/authorize?" + convertResult)
	resp = &result
	return
}

// AccessTokenSelf 获取access_token
func (a *KuaishouAdapter) AccessTokenSelf(ctx context.Context, req *kuaishouModel.AccessTokenReq) (resp *kuaishouModel.AccessTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result kuaishouModel.AccessTokenResp
	if errRequest := a.RequestPostJson(ctx, nil, kuaishouModel.AdUrl+"/rest/openapi/oauth2/authorize/access_token", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// RefreshTokenSelf 刷新access_token
func (a *KuaishouAdapter) RefreshTokenSelf(ctx context.Context, req *kuaishouModel.RefreshTokenReq) (resp *kuaishouModel.RefreshTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result kuaishouModel.RefreshTokenResp
	if errRequest := a.RequestPostJson(ctx, nil, kuaishouModel.AdUrl+"/rest/openapi/oauth2/authorize/refresh_token", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// ApprovalListSelf 拉取token下授权广告账户
func (a *KuaishouAdapter) ApprovalListSelf(ctx context.Context, req *kuaishouModel.ApprovalListReq) (resp *kuaishouModel.ApprovalListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result kuaishouModel.ApprovalListResp
	if errRequest := a.RequestPostJson(ctx, nil, kuaishouModel.AdUrl+"/rest/openapi/oauth2/authorize/approval/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
