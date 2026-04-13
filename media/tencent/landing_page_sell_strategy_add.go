package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LandingPageSellStrategyAddSelf 短剧售卖策略-创建售卖策略
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/add
func (a *TencentAdapter) LandingPageSellStrategyAddSelf(ctx context.Context, req *model.LandingPageSellStrategyAddReq) (
	resp *model.LandingPageSellStrategyAddResp, err error) {
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
	var result model.LandingPageSellStrategyAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/landing_page_sell_strategy/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
