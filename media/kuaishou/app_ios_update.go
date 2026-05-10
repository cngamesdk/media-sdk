package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppIosUpdate 更新iOS应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/update/ios
func (a *KuaishouAdapter) AppIosUpdate(ctx context.Context, req *kuaishouModel.AppIosUpdateReq) (resp *kuaishouModel.AppIosUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppIosUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/update/ios", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
