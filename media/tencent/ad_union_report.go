package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// AdUnionReportsGetSelf 获取联盟广告位报表
// https://developers.e.qq.com/v3.0/docs/api/ad_union_reports/get
func (a *TencentAdapter) AdUnionReportsGetSelf(ctx context.Context, req *model.AdUnionReportsGetReq) (
	resp *model.AdUnionReportsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.AdUnionReportsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/ad_union_reports/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
