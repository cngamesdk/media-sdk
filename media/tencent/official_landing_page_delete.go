package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// OfficialLandingPageDeleteSelf 官方落地页-删除落地页
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page/delete
func (a *TencentAdapter) OfficialLandingPageDeleteSelf(ctx context.Context, req *model.OfficialLandingPageDeleteReq) (
	resp *model.OfficialLandingPageDeleteResp, err error) {
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
	var result model.OfficialLandingPageDeleteResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/official_landing_page/delete?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
