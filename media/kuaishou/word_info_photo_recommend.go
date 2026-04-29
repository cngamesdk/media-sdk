package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// WordInfoPhotoRecommend 创意高相关词推荐
func (a *KuaishouAdapter) WordInfoPhotoRecommend(ctx context.Context, req *kuaishouModel.WordInfoPhotoRecommendReq) (resp *kuaishouModel.WordInfoPhotoRecommendResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.WordInfoPhotoRecommendResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/photo_recommend_word", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
