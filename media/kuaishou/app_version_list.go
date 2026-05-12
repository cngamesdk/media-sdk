package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppVersionList 获取应用版本记录
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/version/list
func (a *KuaishouAdapter) AppVersionList(ctx context.Context, req *kuaishouModel.AppVersionListReq) (resp *kuaishouModel.AppVersionListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppVersionListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/version/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
