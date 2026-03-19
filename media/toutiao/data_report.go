package toutiao

import (
	"context"
	model2 "github.com/cngamesdk/media-sdk/media/toutiao/model"
)

// FileImageAdSelf 获取自定义报表可用指标和维度
// https://open.oceanengine.com/labels/7/docs/1755261744248832?origin=left_nav
func (a *ToutiaoAdapter) ReportCustomConfigGetSelf(ctx context.Context, req *model2.ReportCustomConfigGetReq) (
	resp *model2.ReportCustomConfigGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.ReportCustomConfigGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/report/custom/config/get/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}

// ReportCustomGetSelf 自定义报表
// https://open.oceanengine.com/labels/7/docs/1741387668314126?origin=left_nav
func (a *ToutiaoAdapter) ReportCustomGetSelf(ctx context.Context, req *model2.ReportCustomGetReq) (
	resp *model2.ReportCustomGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result model2.ReportCustomGetResp
	errRequest := a.RequestGet(ctx, headers, model2.BaseUrlApi+"/open_api/v3.0/report/custom/get/", req, &result)
	if errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
