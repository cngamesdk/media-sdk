package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoAiRecommend 获取AI推荐视频
func (a *KuaishouAdapter) VideoAiRecommend(ctx context.Context, req *kuaishouModel.VideoAiRecommendReq) (resp *kuaishouModel.VideoAiRecommendResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoAiRecommendResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/video/aiRecommend", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
