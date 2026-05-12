package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppRelease 发布应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release
func (a *KuaishouAdapter) AppRelease(ctx context.Context, req *kuaishouModel.AppReleaseReq) (resp *kuaishouModel.AppReleaseResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppReleaseResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/release", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
