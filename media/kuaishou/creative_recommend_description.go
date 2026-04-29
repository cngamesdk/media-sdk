package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CreativeRecommendDescription 获取推荐广告语
func (a *KuaishouAdapter) CreativeRecommendDescription(ctx context.Context, req *kuaishouModel.CreativeRecommendDescriptionReq) (resp *kuaishouModel.CreativeRecommendDescriptionResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CreativeRecommendDescriptionResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/creative/recmmend/description", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
