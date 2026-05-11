package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppDetail 获取应用详情
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/detail
func (a *KuaishouAdapter) AppDetail(ctx context.Context, req *kuaishouModel.AppDetailReq) (resp *kuaishouModel.AppDetailResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppDetailResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/detail", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
