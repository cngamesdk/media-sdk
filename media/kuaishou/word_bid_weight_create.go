package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordBidWeightCreate 创建优词提量信息
func (a *KuaishouAdapter) WordBidWeightCreate(ctx context.Context, req *kuaishouModel.WordBidWeightCreateReq) (resp *kuaishouModel.WordBidWeightCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordBidWeightCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/unit/bid_weight/create", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
