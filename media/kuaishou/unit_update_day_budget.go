package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// UnitUpdateDayBudget 修改广告组预算
func (a *KuaishouAdapter) UnitUpdateDayBudget(ctx context.Context, req *kuaishouModel.UnitUpdateDayBudgetReq) (resp *kuaishouModel.UnitUpdateDayBudgetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.UnitUpdateDayBudgetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/ad_unit/update/day_budget", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
