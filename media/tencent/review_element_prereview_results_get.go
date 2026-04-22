package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// ReviewElementPrereviewResultsGet 查询元素的预审结果
// https://developers.e.qq.com/v3.0/docs/api/review_element_prereview_results/get
func (a *TencentAdapter) ReviewElementPrereviewResultsGet(ctx context.Context, req *model.ReviewElementPrereviewResultsGetReq) (
	resp *model.ReviewElementPrereviewResultsGetResp, err error) {
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
	var result model.ReviewElementPrereviewResultsGetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/review_element_prereview_results/get?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
