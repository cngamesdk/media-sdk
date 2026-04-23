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
