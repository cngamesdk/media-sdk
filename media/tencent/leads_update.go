package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LeadsUpdateSelf 更新线索基本信息
// https://developers.e.qq.com/v3.0/docs/api/leads/update
func (a *TencentAdapter) LeadsUpdateSelf(ctx context.Context, req *model.LeadsUpdateReq) (
	resp *model.LeadsUpdateResp, err error) {
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
	var result model.LeadsUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/leads/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
