package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// ImageRelateCreatives 查询图片关联创意
func (a *KuaishouAdapter) ImageRelateCreatives(ctx context.Context, req *kuaishouModel.ImageRelateCreativesReq) (resp *kuaishouModel.ImageRelateCreativesResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.ImageRelateCreativesResp
	if errRequest := a.RequestGet(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/pic/relate/creatives", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
