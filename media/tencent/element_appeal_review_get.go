package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ElementAppealReviewGet 获取元素申诉复审结果
// https://developers.e.qq.com/v3.0/docs/api/element_appeal_review/get
func (a *TencentAdapter) ElementAppealReviewGet(ctx context.Context, req *model.ElementAppealReviewGetReq) (
	resp *model.ElementAppealReviewGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ElementAppealReviewGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/element_appeal_review/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
