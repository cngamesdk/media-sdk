package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LeadsActionTypeReportAddSelf 线索上报DMP平台
// https://developers.e.qq.com/v3.0/docs/api/leads_action_type_report/add
func (a *TencentAdapter) LeadsActionTypeReportAddSelf(ctx context.Context, req *model.LeadsActionTypeReportAddReq) (
	resp *model.LeadsActionTypeReportAddResp, err error) {
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
	var result model.LeadsActionTypeReportAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/leads_action_type_report/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
