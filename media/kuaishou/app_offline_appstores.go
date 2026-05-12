package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppOfflineAppstores 应用商店上下架
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/offline/appstores
func (a *KuaishouAdapter) AppOfflineAppstores(ctx context.Context, req *kuaishouModel.AppOfflineAppstoresReq) (resp *kuaishouModel.AppOfflineAppstoresResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppOfflineAppstoresResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/offline/appstores", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
