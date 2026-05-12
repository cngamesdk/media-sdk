package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppReleaseLatest 获取最新未发布应用包
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/latest
func (a *KuaishouAdapter) AppReleaseLatest(ctx context.Context, req *kuaishouModel.AppReleaseLatestReq) (resp *kuaishouModel.AppReleaseLatestResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppReleaseLatestResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/release/latest", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
