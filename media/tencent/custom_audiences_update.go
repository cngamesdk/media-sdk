package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// CustomAudiencesUpdate 更新客户人群
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/update
func (a *TencentAdapter) CustomAudiencesUpdate(ctx context.Context, req *model.CustomAudiencesUpdateReq) (
	resp *model.CustomAudiencesUpdateResp, err error) {
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
	var result model.CustomAudiencesUpdateResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/custom_audiences/update?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
