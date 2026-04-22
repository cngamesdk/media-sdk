package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// IllegalComplaintGet 获取直客广告主违规申述列表
// https://developers.e.qq.com/v3.0/docs/api/illegal_complaint/get
func (a *TencentAdapter) IllegalComplaintGet(ctx context.Context, req *model.IllegalComplaintGetReq) (
	resp *model.IllegalComplaintGetResp, err error) {
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
	var result model.IllegalComplaintGetResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/illegal_complaint/get?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
