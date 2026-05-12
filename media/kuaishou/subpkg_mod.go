package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SubpkgMod 更新/恢复/删除应用分包
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/mod
func (a *KuaishouAdapter) SubpkgMod(ctx context.Context, req *kuaishouModel.SubpkgModReq) (resp *kuaishouModel.SubpkgModResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SubpkgModResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/subpkg/mod", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
