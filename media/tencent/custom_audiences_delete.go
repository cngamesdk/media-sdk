package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// CustomAudiencesDelete 删除客户人群
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/delete
func (a *TencentAdapter) CustomAudiencesDelete(ctx context.Context, req *model.CustomAudiencesDeleteReq) (
	resp *model.CustomAudiencesDeleteResp, err error) {
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
	var result model.CustomAudiencesDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/custom_audiences/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
