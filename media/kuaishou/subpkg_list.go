package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// SubpkgList 获取分包管理/回收站列表
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subPackage/list
func (a *KuaishouAdapter) SubpkgList(ctx context.Context, req *kuaishouModel.SubpkgListReq) (resp *kuaishouModel.SubpkgListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.SubpkgListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/appcenter/subPackage/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
