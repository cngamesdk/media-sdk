package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserBudgetGet 账户日预算查询
func (a *KuaishouAdapter) AdvertiserBudgetGet(ctx context.Context, req *kuaishouModel.AdvertiserBudgetGetReq) (resp *kuaishouModel.AdvertiserBudgetGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserBudgetGetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/budget/get", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
