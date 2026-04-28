package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserUpdateBudget 修改账户预算
func (a *KuaishouAdapter) AdvertiserUpdateBudget(ctx context.Context, req *kuaishouModel.AdvertiserUpdateBudgetReq) (resp *kuaishouModel.AdvertiserUpdateBudgetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserUpdateBudgetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/update/budget", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
