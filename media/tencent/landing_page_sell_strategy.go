package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// LandingPageSellStrategyGetSelf 短剧售卖策略-获取售卖策略列表
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/get
func (a *TencentAdapter) LandingPageSellStrategyGetSelf(ctx context.Context, req *model.LandingPageSellStrategyGetReq) (
	resp *model.LandingPageSellStrategyGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.LandingPageSellStrategyGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/landing_page_sell_strategy/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
