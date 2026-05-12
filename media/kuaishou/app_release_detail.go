package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppReleaseDetail 获取新版应用发布详情
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/detail
func (a *KuaishouAdapter) AppReleaseDetail(ctx context.Context, req *kuaishouModel.AppReleaseDetailReq) (resp *kuaishouModel.AppReleaseDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppReleaseDetailResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/release/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
