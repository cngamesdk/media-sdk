package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"github.com/cngamesdk/media-sdk/utils"
)

// OfficialLandingPageCompAddSelf 官方落地页-基于组件创建
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_component/add
func (a *TencentAdapter) OfficialLandingPageCompAddSelf(ctx context.Context, req *model.OfficialLandingPageCompAddReq) (
	resp *model.OfficialLandingPageCompAddResp, err error) {
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
	var result model.OfficialLandingPageCompAddResp
	if requestErr := a.RequestPostJson(ctx, headers, model.ApiUrl3+"/official_landing_page_component/add?"+globalQuery, req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
