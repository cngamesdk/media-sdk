package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// OAuth2AuthorizeSelf 授权
func (a *TencentAdapter) OAuth2AuthorizeSelf(ctx context.Context, req *model.OAuth2AuthorizeReq) (resp string, err error) {
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
	resp = model.DevelopersUrl + "/oauth/authorize?" + convertResult
	return
}
