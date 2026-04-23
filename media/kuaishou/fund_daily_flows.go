package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserFundDailyFlowsSelf 获取广告账户流水信息
func (a *KuaishouAdapter) AdvertiserFundDailyFlowsSelf(ctx context.Context, req *kuaishouModel.AdvertiserFundDailyFlowsReq) (resp *kuaishouModel.AdvertiserFundDailyFlowsResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserFundDailyFlowsResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/fund/daily_flows", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
