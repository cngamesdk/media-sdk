package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// OfficialLandingPageDetailGetSelf 官方落地页-获取落地页详情
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_detail/get
func (a *TencentAdapter) OfficialLandingPageDetailGetSelf(ctx context.Context, req *model.OfficialLandingPageDetailGetReq) (
	resp *model.OfficialLandingPageDetailGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.OfficialLandingPageDetailGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/official_landing_page_detail/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
