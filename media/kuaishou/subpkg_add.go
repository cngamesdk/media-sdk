package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SubpkgAdd 新建应用分包
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/add
func (a *KuaishouAdapter) SubpkgAdd(ctx context.Context, req *kuaishouModel.SubpkgAddReq) (resp *kuaishouModel.SubpkgAddResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SubpkgAddResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/subpkg/add", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
