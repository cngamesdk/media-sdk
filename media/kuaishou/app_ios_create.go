package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppIosCreate 创建iOS应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/ios
func (a *KuaishouAdapter) AppIosCreate(ctx context.Context, req *kuaishouModel.AppIosCreateReq) (resp *kuaishouModel.AppIosCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppIosCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/create/ios", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
