package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// VideoShareNew 视频库推送视频（新版）
func (a *KuaishouAdapter) VideoShareNew(ctx context.Context, req *kuaishouModel.VideoShareNewReq) (resp *kuaishouModel.VideoShareNewResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.VideoShareNewResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/file/ad/video/share/new", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
