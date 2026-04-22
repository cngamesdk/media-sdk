package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ComponentReviewResultsGet 查询组件审核结果
// https://developers.e.qq.com/v3.0/docs/api/component_review_results/get
func (a *TencentAdapter) ComponentReviewResultsGet(ctx context.Context, req *model.ComponentReviewResultsGetReq) (
	resp *model.ComponentReviewResultsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentReviewResultsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/component_review_results/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
