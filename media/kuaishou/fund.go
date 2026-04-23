package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AdvertiserFundGetSelf 获取广告账户余额信息
func (a *KuaishouAdapter) AdvertiserFundGetSelf(ctx context.Context, req *kuaishouModel.AdvertiserFundGetReq) (resp *kuaishouModel.AdvertiserFundGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AdvertiserFundGetResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/advertiser/fund/get", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
