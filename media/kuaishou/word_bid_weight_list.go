package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordBidWeightList 获取优词提量列表
func (a *KuaishouAdapter) WordBidWeightList(ctx context.Context, req *kuaishouModel.WordBidWeightListReq) (resp *kuaishouModel.WordBidWeightListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordBidWeightListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/unit/bid_weight/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
