package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppShareCancel 取消应用共享
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/share/cancel
func (a *KuaishouAdapter) AppShareCancel(ctx context.Context, req *kuaishouModel.AppShareCancelReq) (resp *kuaishouModel.AppShareCancelResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppShareCancelResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/share/cancel", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
