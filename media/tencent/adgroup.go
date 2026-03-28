package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// AdgroupsGetSelf 获取广告
// https://developers.e.qq.com/v3.0/docs/api/adgroups/get
func (a *TencentAdapter) AdgroupsGetSelf(ctx context.Context, req *model.AdgroupsGetReq) (
	resp *model.AdgroupsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AdgroupsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/adgroups/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
