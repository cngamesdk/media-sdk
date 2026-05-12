package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppIosReportUpdate iOS应用上报更新
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/ios/update
func (a *KuaishouAdapter) AppIosReportUpdate(ctx context.Context, req *kuaishouModel.AppIosReportUpdateReq) (resp *kuaishouModel.AppIosReportUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppIosReportUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/ios/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
