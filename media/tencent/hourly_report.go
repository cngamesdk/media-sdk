package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// HourlyReportsGetSelf 获取小时报表
// https://developers.e.qq.com/v3.0/docs/api/hourly_reports/get
func (a *TencentAdapter) HourlyReportsGetSelf(ctx context.Context, req *model.HourlyReportsGetReq) (
	resp *model.HourlyReportsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.HourlyReportsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/hourly_reports/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
