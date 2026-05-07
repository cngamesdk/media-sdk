package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageGet 查询图片信息
func (a *KuaishouAdapter) ImageGet(ctx context.Context, req *kuaishouModel.ImageGetReq) (resp *kuaishouModel.ImageGetResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImageGetResp
	if errRequest := a.RequestGet(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/image/get", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
