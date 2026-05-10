package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppAndroidUpdate 更新Android应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/update/android
func (a *KuaishouAdapter) AppAndroidUpdate(ctx context.Context, req *kuaishouModel.AppAndroidUpdateReq) (resp *kuaishouModel.AppAndroidUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppAndroidUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/update/android", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
