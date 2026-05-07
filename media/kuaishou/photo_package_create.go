package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageCreate 新建素材包
func (a *KuaishouAdapter) PhotoPackageCreate(ctx context.Context, req *kuaishouModel.PhotoPackageCreateReq) (resp *kuaishouModel.PhotoPackageCreateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageCreateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/add", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
