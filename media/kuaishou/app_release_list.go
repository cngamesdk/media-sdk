package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppReleaseList 获取新版应用发布列表【单元创编】
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/list
func (a *KuaishouAdapter) AppReleaseList(ctx context.Context, req *kuaishouModel.AppReleaseListReq) (resp *kuaishouModel.AppReleaseListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppReleaseListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/release/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
