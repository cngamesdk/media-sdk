package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SubpkgRetryBuild 分包失败重新构建
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/retryBuildSubPackage
func (a *KuaishouAdapter) SubpkgRetryBuild(ctx context.Context, req *kuaishouModel.SubpkgRetryBuildReq) (resp *kuaishouModel.SubpkgRetryBuildResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SubpkgRetryBuildResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/retryBuildSubPackage", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
