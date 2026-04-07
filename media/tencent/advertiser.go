package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// AdvertiserUpdateDailyBudgetSelf 批量修改广告主日限额
// https://developers.e.qq.com/v3.0/docs/api/advertiser/update_daily_budget
func (a *TencentAdapter) AdvertiserUpdateDailyBudgetSelf(ctx context.Context, req *model.AdvertiserUpdateDailyBudgetReq) (
	resp *model.AdvertiserUpdateDailyBudgetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()
	headers := make(model.Headers)
	headers.Json()
	var result model.AdvertiserUpdateDailyBudgetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/advertiser/update_daily_budget?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
