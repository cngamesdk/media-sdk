package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppAndroidCreate 创建Android应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/android
func (a *KuaishouAdapter) AppAndroidCreate(ctx context.Context, req *kuaishouModel.AppAndroidCreateReq) (resp *kuaishouModel.AppAndroidCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppAndroidCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/create/android", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
