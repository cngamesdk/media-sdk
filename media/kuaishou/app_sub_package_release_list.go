package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AppSubPackageReleaseList 获取新版分包发布列表【单元创编】
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subPackage/release/list
func (a *KuaishouAdapter) AppSubPackageReleaseList(ctx context.Context, req *kuaishouModel.AppSubPackageReleaseListReq) (resp *kuaishouModel.AppSubPackageReleaseListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AppSubPackageReleaseListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/subPackage/release/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
