package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LeadsAddSelf 新增线索
// https://developers.e.qq.com/v3.0/docs/api/leads/add
func (a *TencentAdapter) LeadsAddSelf(ctx context.Context, req *model.LeadsAddReq) (
	resp *model.LeadsAddResp, err error) {
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
	var result model.LeadsAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/leads/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
