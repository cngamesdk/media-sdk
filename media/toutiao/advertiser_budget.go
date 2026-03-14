package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// EbpAppListSelf 帐户预算
// https://open.oceanengine.com/labels/7/docs/1696710531128335?origin=left_nav
func (a *ToutiaoAdapter) AdvertiserBudgetGetSelf(ctx context.Context, req *model2.AdvertiserBudgetGetReq) (resp *model2.AdvertiserBudgetGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.AdvertiserBudgetGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/2/advertiser/budget/get/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
