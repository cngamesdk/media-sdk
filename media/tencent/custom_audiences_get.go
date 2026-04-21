package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// CustomAudiencesGet 获取客户人群
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/get
func (a *TencentAdapter) CustomAudiencesGet(ctx context.Context, req *model.CustomAudiencesGetReq) (
	resp *model.CustomAudiencesGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.CustomAudiencesGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/custom_audiences/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
