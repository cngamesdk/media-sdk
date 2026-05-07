package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageQuery 查询素材包
func (a *KuaishouAdapter) PhotoPackageQuery(ctx context.Context, req *kuaishouModel.PhotoPackageQueryReq) (resp *kuaishouModel.PhotoPackageQueryResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageQueryResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/query", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
