package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// CustomAudienceEstimationsGet 人群覆盖数预估
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_estimations/get
func (a *TencentAdapter) CustomAudienceEstimationsGet(ctx context.Context, req *model.CustomAudienceEstimationsGetReq) (
	resp *model.CustomAudienceEstimationsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.CustomAudienceEstimationsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/custom_audience_estimations/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
