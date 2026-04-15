package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// DailyReportsGetSelf 获取日报表
// https://developers.e.qq.com/v3.0/docs/api/daily_reports/get
func (a *TencentAdapter) DailyReportsGetSelf(ctx context.Context, req *model.DailyReportsGetReq) (
	resp *model.DailyReportsGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	var result model.DailyReportsGetResp
	if requestErr := a.RequestGet(ctx, nil, model.ApiUrl3+"/daily_reports/get", req, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
