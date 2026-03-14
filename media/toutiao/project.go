package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// AdvertiserBudgetGetSelf 帐户预算
// https://open.oceanengine.com/labels/7/docs/1696710531128335?origin=left_nav
func (a *ToutiaoAdapter) ProjectCreateSelf(ctx context.Context, req *model2.ProjectCreateReq) (resp *model2.ProjectCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.ProjectCreateResp
	errRequest := a.RequestPostJson(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/project/create/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
