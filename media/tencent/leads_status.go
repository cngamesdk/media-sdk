package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LeadsStatusUpdateSelf 更新线索状态
// https://developers.e.qq.com/v3.0/docs/api/leads_status/update
func (a *TencentAdapter) LeadsStatusUpdateSelf(ctx context.Context, req *model.LeadsStatusUpdateReq) (
	resp *model.LeadsStatusUpdateResp, err error) {
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
	var result model.LeadsStatusUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/leads_status/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
