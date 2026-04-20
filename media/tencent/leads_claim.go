package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// LeadsClaimUpdateSelf 更新线索归因信息
// https://developers.e.qq.com/v3.0/docs/api/leads_claim/update
func (a *TencentAdapter) LeadsClaimUpdateSelf(ctx context.Context, req *model.LeadsClaimUpdateReq) (
	resp *model.LeadsClaimUpdateResp, err error) {
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
	var result model.LeadsClaimUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/leads_claim/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
