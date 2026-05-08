package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageUploadTokenVerify 领用上传token
// https://ad.e.kuaishou.com/rest/openapi/gw/ad/common/upload/token/verify
func (a *KuaishouAdapter) ImageUploadTokenVerify(ctx context.Context, req *kuaishouModel.ImageUploadTokenVerifyReq) (resp *kuaishouModel.ImageUploadTokenVerifyResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImageUploadTokenVerifyResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/ad/common/upload/token/verify", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
