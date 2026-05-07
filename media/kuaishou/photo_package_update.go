package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageUpdate 编辑素材包
func (a *KuaishouAdapter) PhotoPackageUpdate(ctx context.Context, req *kuaishouModel.PhotoPackageUpdateReq) (resp *kuaishouModel.PhotoPackageUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
