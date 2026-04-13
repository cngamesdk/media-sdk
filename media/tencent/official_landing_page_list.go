package tencent

import (
	"context"

	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// OfficialLandingPageListGetSelf 官方落地页-获取落地页列表
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_list/get
func (a *TencentAdapter) OfficialLandingPageListGetSelf(ctx context.Context, req *model.OfficialLandingPageListGetReq) (
	resp *model.OfficialLandingPageListGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.OfficialLandingPageListGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/official_landing_page_list/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
