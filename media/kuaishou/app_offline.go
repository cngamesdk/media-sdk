package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppOffline 应用下架
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/offline
func (a *KuaishouAdapter) AppOffline(ctx context.Context, req *kuaishouModel.AppOfflineReq) (resp *kuaishouModel.AppOfflineResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppOfflineResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/offline", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
