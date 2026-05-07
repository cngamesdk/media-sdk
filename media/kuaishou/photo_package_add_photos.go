package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// PhotoPackageAddPhotos 添加视频至素材包
func (a *KuaishouAdapter) PhotoPackageAddPhotos(ctx context.Context, req *kuaishouModel.PhotoPackageAddPhotosReq) (resp *kuaishouModel.PhotoPackageAddPhotosResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.PhotoPackageAddPhotosResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/photoPackage/addPhotosToPackage", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
