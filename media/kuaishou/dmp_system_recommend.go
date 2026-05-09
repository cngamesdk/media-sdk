package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// DmpSystemRecommend 系统推荐定向/排除人群包
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/v1/tool/population/system/recommend
func (a *KuaishouAdapter) DmpSystemRecommend(ctx context.Context, req *kuaishouModel.DmpSystemRecommendReq) (resp *kuaishouModel.DmpSystemRecommendResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.DmpSystemRecommendResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/tool/population/system/recommend", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
