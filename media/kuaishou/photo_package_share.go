package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageShare 素材包推送
func (a *KuaishouAdapter) PhotoPackageShare(ctx context.Context, req *kuaishouModel.PhotoPackageShareReq) (resp *kuaishouModel.PhotoPackageShareResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageShareResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/share", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
