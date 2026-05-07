package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoRelateCreatives 视频关联创意数查询
func (a *KuaishouAdapter) VideoRelateCreatives(ctx context.Context, req *kuaishouModel.VideoRelateCreativesReq) (resp *kuaishouModel.VideoRelateCreativesResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoRelateCreativesResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/file/ad/video/relate/creatives", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
