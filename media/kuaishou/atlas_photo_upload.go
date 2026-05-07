package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AtlasPhotoUpload 上传图文视频
func (a *KuaishouAdapter) AtlasPhotoUpload(ctx context.Context, req *kuaishouModel.AtlasPhotoUploadReq) (resp *kuaishouModel.AtlasPhotoUploadResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AtlasPhotoUploadResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/upload/atlasPhoto", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
