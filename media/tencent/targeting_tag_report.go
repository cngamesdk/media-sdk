package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// TargetingTagReportsGetSelf 获取定向标签报表
// https://developers.e.qq.com/v3.0/docs/api/targeting_tag_reports/get
func (a *TencentAdapter) TargetingTagReportsGetSelf(ctx context.Context, req *model.TargetingTagReportsGetReq) (
	resp *model.TargetingTagReportsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.TargetingTagReportsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/targeting_tag_reports/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
