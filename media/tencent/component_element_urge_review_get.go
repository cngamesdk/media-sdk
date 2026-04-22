package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ComponentElementUrgeReviewGet 获取创意组件元素催审状态
// https://developers.e.qq.com/v3.0/docs/api/component_element_urge_review/get
func (a *TencentAdapter) ComponentElementUrgeReviewGet(ctx context.Context, req *model.ComponentElementUrgeReviewGetReq) (
	resp *model.ComponentElementUrgeReviewGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.ComponentElementUrgeReviewGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/component_element_urge_review/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
