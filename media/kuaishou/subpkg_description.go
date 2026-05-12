package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SubpkgDescription 修改应用分包备注
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/description
func (a *KuaishouAdapter) SubpkgDescription(ctx context.Context, req *kuaishouModel.SubpkgDescriptionReq) (resp *kuaishouModel.SubpkgDescriptionResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SubpkgDescriptionResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/subpkg/description", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
