package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppOnline 应用上架
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/online
func (a *KuaishouAdapter) AppOnline(ctx context.Context, req *kuaishouModel.AppOnlineReq) (resp *kuaishouModel.AppOnlineResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppOnlineResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/online", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
