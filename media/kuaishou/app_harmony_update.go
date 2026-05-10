package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppHarmonyUpdate 更新Harmony应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/update/harmony
func (a *KuaishouAdapter) AppHarmonyUpdate(ctx context.Context, req *kuaishouModel.AppHarmonyUpdateReq) (resp *kuaishouModel.AppHarmonyUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppHarmonyUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/update/harmony", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
