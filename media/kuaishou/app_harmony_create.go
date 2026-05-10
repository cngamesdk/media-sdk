package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppHarmonyCreate 创建Harmony应用
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/harmony
func (a *KuaishouAdapter) AppHarmonyCreate(ctx context.Context, req *kuaishouModel.AppHarmonyCreateReq) (resp *kuaishouModel.AppHarmonyCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppHarmonyCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/app/create/harmony", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
