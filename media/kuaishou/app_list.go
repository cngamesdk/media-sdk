package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppList 获取应用列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/list
func (a *KuaishouAdapter) AppList(ctx context.Context, req *kuaishouModel.AppListReq) (resp *kuaishouModel.AppListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
