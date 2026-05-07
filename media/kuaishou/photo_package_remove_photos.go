package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageRemovePhotos 从素材包内删除视频
func (a *KuaishouAdapter) PhotoPackageRemovePhotos(ctx context.Context, req *kuaishouModel.PhotoPackageRemovePhotosReq) (resp *kuaishouModel.PhotoPackageRemovePhotosResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageRemovePhotosResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/removePhotosFromPackage", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
