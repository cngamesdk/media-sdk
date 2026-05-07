package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageDel 删除素材包
func (a *KuaishouAdapter) PhotoPackageDel(ctx context.Context, req *kuaishouModel.PhotoPackageDelReq) (resp *kuaishouModel.PhotoPackageDelResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageDelResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/del", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
