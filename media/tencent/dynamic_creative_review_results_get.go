package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// DynamicCreativeReviewResultsGet 查询动态创意审核结果
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creative_review_results/get
func (a *TencentAdapter) DynamicCreativeReviewResultsGet(ctx context.Context, req *model.DynamicCreativeReviewResultsGetReq) (
	resp *model.DynamicCreativeReviewResultsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.DynamicCreativeReviewResultsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/dynamic_creative_review_results/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
