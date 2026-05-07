package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageList 查询图片list
func (a *KuaishouAdapter) ImageList(ctx context.Context, req *kuaishouModel.ImageListReq) (resp *kuaishouModel.ImageListResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImageListResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/image/list", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
