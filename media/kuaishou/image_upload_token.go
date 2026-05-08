package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageUploadToken 获取图片上传token
// https://ad.e.kuaishou.com/rest/openapi/gw/ad/common/upload/token/generate
func (a *KuaishouAdapter) ImageUploadToken(ctx context.Context, req *kuaishouModel.ImageUploadTokenReq) (resp *kuaishouModel.ImageUploadTokenResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImageUploadTokenResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/ad/common/upload/token/generate", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
