package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ComponentElementUrgeReviewAdd 创意组件元素催审
// https://developers.e.qq.com/v3.0/docs/api/component_element_urge_review/add
func (a *TencentAdapter) ComponentElementUrgeReviewAdd(ctx context.Context, req *model.ComponentElementUrgeReviewAddReq) (
	resp *model.ComponentElementUrgeReviewAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	globalQuery, globalQueryErr := utils.ConvertStructToQueryString(req.GlobalReq)
	if globalQueryErr != nil {
		err = globalQueryErr
		return
	}
	req.GlobalReq.Clear()
	headers := make(model.Headers)
	headers.Json()
	var result model.ComponentElementUrgeReviewAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/component_element_urge_review/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
